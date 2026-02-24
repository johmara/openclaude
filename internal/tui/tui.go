package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johmara/openclaude/internal/claude"
	"github.com/johmara/openclaude/internal/tui/components/chat"
	"github.com/johmara/openclaude/internal/tui/components/dialog"
	"github.com/johmara/openclaude/internal/tui/components/status"
	"github.com/johmara/openclaude/internal/tui/layout"
	"github.com/johmara/openclaude/internal/tui/theme"
)

// App is the interface that the application layer must satisfy.
type App interface {
	SendMessage(prompt string) error
	CancelGeneration()
	Subscribe(id string) <-chan claude.Event
	Unsubscribe(id string)
	CreateSession(name string)
	SetActiveSession(idx int)
	SessionCount() int
	ActiveSessionName() string
}

// DialogType identifies the currently active dialog.
type DialogType int

const (
	DialogNone DialogType = iota
	DialogCommands
	DialogSessions
	DialogTheme
	DialogFilePicker
	DialogHelp
)

// ClaudeEventMsg wraps a claude event for the Bubble Tea message loop.
type ClaudeEventMsg struct {
	Event claude.Event
}

// sendDoneMsg signals that message sending finished (with possible error).
type sendDoneMsg struct {
	err error
}

// leaderTimeoutMsg fires when the leader key window expires.
type leaderTimeoutMsg struct{}

// ctrlCTimeoutMsg fires when the Ctrl+C double-press window expires.
type ctrlCTimeoutMsg struct{}

// Model is the root Bubble Tea model.
type Model struct {
	app App

	// Components
	chatPage  chat.Page
	statusBar status.Bar
	keys      KeyMap

	// Dialogs
	activeDialog  DialogType
	commandDialog dialog.Commands
	themeDialog   dialog.ThemePicker
	fileDialog    dialog.FilePicker
	helpDialog    dialog.Help

	// Leader key state
	leaderActive bool

	// Double Ctrl+C state
	ctrlCPressed bool

	// State
	width       int
	height      int
	subID       string
	subCh       <-chan claude.Event
	initialized bool
}

// New creates the root TUI model.
func New(app App) Model {
	return Model{
		app:       app,
		keys:      DefaultKeyMap(),
		statusBar: status.New(),
		subID:     "tui-main",
	}
}

// Init initializes the Bubble Tea model.
func (m Model) Init() tea.Cmd {
	return nil
}

// subscribeCmd returns a tea.Cmd that reads one event from the broker channel.
func (m Model) subscribeCmd() tea.Cmd {
	ch := m.subCh
	if ch == nil {
		return nil
	}
	return func() tea.Msg {
		evt, ok := <-ch
		if !ok {
			return nil
		}
		return ClaudeEventMsg{Event: evt}
	}
}

// Update is the main Bubble Tea update loop.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Initialize chat page on first resize (when we know terminal size)
		if !m.initialized {
			m.initialized = true
			m.chatPage = chat.NewPage()
			m.chatPage.SetSize(msg.Width, msg.Height-1)
			cmds = append(cmds, m.chatPage.Init())
		} else {
			m.chatPage.SetSize(msg.Width, msg.Height-1) // -1 for status bar
		}

		m.statusBar.SetSize(msg.Width)

		return m, tea.Batch(cmds...)

	case ClaudeEventMsg:
		m.chatPage.HandleEvent(msg.Event)

		switch msg.Event.Type {
		case claude.EventInit:
			m.statusBar.SetModel(msg.Event.Model)
			m.statusBar.SetStreaming(true)
		case claude.EventResult:
			m.statusBar.SetCost(msg.Event.TotalCostUSD)
			m.statusBar.SetTokens(msg.Event.InputTokens + msg.Event.OutputTokens)
			m.statusBar.SetStreaming(false)
			m.chatPage.FocusEditor()
		case claude.EventError:
			m.statusBar.SetMessage("Error: " + msg.Event.ResultError)
			m.statusBar.SetStreaming(false)
		}

		// Re-subscribe for next event
		cmds = append(cmds, m.subscribeCmd())
		return m, tea.Batch(cmds...)

	case sendDoneMsg:
		if msg.err != nil {
			m.statusBar.SetMessage("Error: " + msg.err.Error())
			m.statusBar.SetStreaming(false)
		}
		return m, nil

	case chat.SendMsg:
		return m.handleSend(msg.Text)

	case dialog.CloseMsg:
		m.activeDialog = DialogNone
		m.chatPage.FocusEditor()
		return m, nil

	case dialog.CommandMsg:
		return m.handleCommand(msg.Command)

	case dialog.ThemeChangedMsg:
		theme.SetByIndex(msg.Index)
		m.activeDialog = DialogNone
		m.chatPage.FocusEditor()
		return m, nil

	case dialog.FileSelectedMsg:
		m.activeDialog = DialogNone
		m.chatPage.FocusEditor()
		// Insert file path into editor context
		m.statusBar.SetMessage("Selected: " + msg.Path)
		return m, nil

	case leaderTimeoutMsg:
		m.leaderActive = false
		m.statusBar.SetLeader(false)
		return m, nil

	case ctrlCTimeoutMsg:
		m.ctrlCPressed = false
		m.statusBar.SetMessage("")
		return m, nil

	case tea.KeyMsg:
		// Handle dialog-level keys first
		if m.activeDialog != DialogNone {
			return m.updateDialog(msg)
		}

		// Leader key second-press dispatch
		if m.leaderActive {
			m.leaderActive = false
			m.statusBar.SetLeader(false)
			return m.handleLeaderAction(msg)
		}

		// Global keybindings
		switch {
		case key.Matches(msg, m.keys.Quit):
			if m.ctrlCPressed {
				// Second Ctrl+C → quit
				return m, tea.Quit
			}
			// First Ctrl+C → start window
			m.ctrlCPressed = true
			m.statusBar.SetMessage("Press Ctrl+C again to quit")
			cmds = append(cmds, tea.Tick(2*time.Second, func(time.Time) tea.Msg {
				return ctrlCTimeoutMsg{}
			}))
			return m, tea.Batch(cmds...)

		case key.Matches(msg, m.keys.Leader):
			m.leaderActive = true
			m.statusBar.SetLeader(true)
			cmds = append(cmds, tea.Tick(2*time.Second, func(time.Time) tea.Msg {
				return leaderTimeoutMsg{}
			}))
			return m, tea.Batch(cmds...)

		case key.Matches(msg, m.keys.CommandPalette):
			m.activeDialog = DialogCommands
			m.commandDialog = dialog.NewCommands()
			m.commandDialog.SetSize(m.width, m.height)
			m.chatPage.BlurEditor()
			return m, m.commandDialog.Init()

		case key.Matches(msg, m.keys.Cancel):
			if m.chatPage.IsStreaming() {
				m.app.CancelGeneration()
				m.statusBar.SetStreaming(false)
				m.statusBar.SetMessage("Generation cancelled")
				return m, nil
			}
		}
	}

	// Update chat page
	var cmd tea.Cmd
	m.chatPage, cmd = m.chatPage.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// handleLeaderAction processes the second key after the leader key.
func (m Model) handleLeaderAction(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "s":
		// Session switcher — currently maps to command palette
		m.activeDialog = DialogCommands
		m.commandDialog = dialog.NewCommands()
		m.commandDialog.SetSize(m.width, m.height)
		m.chatPage.BlurEditor()
		return m, m.commandDialog.Init()

	case "t":
		m.activeDialog = DialogTheme
		m.themeDialog = dialog.NewThemePicker()
		m.themeDialog.SetSize(m.width, m.height)
		m.chatPage.BlurEditor()
		return m, m.themeDialog.Init()

	case "f":
		m.activeDialog = DialogFilePicker
		m.fileDialog = dialog.NewFilePicker()
		m.fileDialog.SetSize(m.width, m.height)
		m.chatPage.BlurEditor()
		return m, m.fileDialog.Init()

	case "?":
		m.activeDialog = DialogHelp
		m.helpDialog = dialog.NewHelp()
		m.helpDialog.SetSize(m.width, m.height)
		m.chatPage.BlurEditor()
		return m, nil

	case "n":
		m.app.CreateSession("New Session")
		m.chatPage.Clear()
		sess := m.app.ActiveSessionName()
		m.statusBar.SetSession(sess)
		m.statusBar.SetMessage("New session created")
		return m, nil
	}

	// Unknown leader sequence — ignore
	return m, nil
}

// handleSend sends a message and starts streaming.
func (m Model) handleSend(text string) (tea.Model, tea.Cmd) {
	m.chatPage.AddUserMessage(text)
	m.statusBar.SetStreaming(true)
	m.statusBar.SetMessage("")

	// Subscribe to broker events
	m.subCh = m.app.Subscribe(m.subID)

	// Send the message in background
	sendCmd := func() tea.Msg {
		err := m.app.SendMessage(text)
		return sendDoneMsg{err: err}
	}

	return m, tea.Batch(sendCmd, m.subscribeCmd())
}

// handleCommand processes a command palette selection.
func (m Model) handleCommand(cmd string) (tea.Model, tea.Cmd) {
	m.activeDialog = DialogNone
	m.chatPage.FocusEditor()

	switch cmd {
	case "new_session":
		m.app.CreateSession("New Session")
		m.chatPage.Clear()
		m.statusBar.SetMessage("New session created")
	case "change_theme":
		m.activeDialog = DialogTheme
		m.themeDialog = dialog.NewThemePicker()
		m.themeDialog.SetSize(m.width, m.height)
		m.chatPage.BlurEditor()
		return m, m.themeDialog.Init()
	case "file_picker":
		m.activeDialog = DialogFilePicker
		m.fileDialog = dialog.NewFilePicker()
		m.fileDialog.SetSize(m.width, m.height)
		m.chatPage.BlurEditor()
		return m, m.fileDialog.Init()
	case "toggle_sidebar":
		m.chatPage.ToggleSidebar()
	case "clear_chat":
		m.chatPage.Clear()
		m.statusBar.SetMessage("Chat cleared")
	case "help":
		m.activeDialog = DialogHelp
		m.helpDialog = dialog.NewHelp()
		m.helpDialog.SetSize(m.width, m.height)
		m.chatPage.BlurEditor()
	case "quit":
		return m, tea.Quit
	}

	return m, nil
}

// updateDialog routes messages to the active dialog.
func (m Model) updateDialog(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.activeDialog {
	case DialogCommands:
		m.commandDialog, cmd = m.commandDialog.Update(msg)
	case DialogTheme:
		m.themeDialog, cmd = m.themeDialog.Update(msg)
	case DialogFilePicker:
		m.fileDialog, cmd = m.fileDialog.Update(msg)
	case DialogHelp:
		m.helpDialog, cmd = m.helpDialog.Update(msg)
	}

	return m, cmd
}

// View renders the full TUI.
func (m Model) View() string {
	if m.width == 0 || m.height == 0 || !m.initialized {
		return "Loading..."
	}

	// Render main content
	chatView := m.chatPage.View()
	statusView := m.statusBar.View()

	mainContent := lipgloss.JoinVertical(lipgloss.Left, chatView, statusView)

	// Overlay active dialog
	if m.activeDialog != DialogNone {
		var dialogView string
		switch m.activeDialog {
		case DialogCommands:
			dialogView = m.commandDialog.View()
		case DialogTheme:
			dialogView = m.themeDialog.View()
		case DialogFilePicker:
			dialogView = m.fileDialog.View()
		case DialogHelp:
			dialogView = m.helpDialog.View()
		}

		if dialogView != "" {
			mainContent = layout.PlaceOverlay(mainContent, dialogView, m.width, m.height)
		}
	}

	return mainContent
}
