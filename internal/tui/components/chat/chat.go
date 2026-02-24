package chat

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johmara/openclaude/internal/claude"
)

// Page assembles the full chat layout with messages, editor, and sidebar.
type Page struct {
	messages Messages
	editor   Editor
	sidebar  Sidebar

	width       int
	height      int
	sidebarW    int
	editorH     int
	showSidebar bool
	initialized bool
}

// NewPage creates a new chat page.
func NewPage() Page {
	return Page{
		messages:    NewMessages(),
		editor:      NewEditor(),
		sidebar:     NewSidebar(),
		sidebarW:    30,
		editorH:     5,
		showSidebar: true,
		initialized: true,
	}
}

// SetSize updates all component sizes.
func (p *Page) SetSize(w, h int) {
	if !p.initialized {
		return
	}
	p.width = w
	p.height = h

	statusH := 1
	availH := h - statusH
	if availH < 5 {
		availH = 5
	}
	messagesH := availH - p.editorH

	if p.showSidebar && w > 80 {
		mainW := w - p.sidebarW
		p.messages.SetSize(mainW, messagesH)
		p.editor.SetSize(mainW, p.editorH)
		p.sidebar.SetSize(p.sidebarW, availH)
	} else {
		p.messages.SetSize(w, messagesH)
		p.editor.SetSize(w, p.editorH)
	}
}

// HandleEvent forwards a Claude event to the messages view.
func (p *Page) HandleEvent(evt claude.Event) {
	p.messages.HandleEvent(evt)

	// Update sidebar metadata
	switch evt.Type {
	case claude.EventInit:
		p.sidebar.SetModel(evt.Model)
	case claude.EventResult:
		p.sidebar.SetCost(evt.TotalCostUSD)
		p.sidebar.SetTokens(evt.InputTokens + evt.OutputTokens)
	}
}

// AddUserMessage adds a user message to the display.
func (p *Page) AddUserMessage(text string) {
	p.messages.AddUserMessage(text)
}

// IsStreaming returns whether content is currently streaming.
func (p Page) IsStreaming() bool {
	return p.messages.IsStreaming()
}

// Update handles messages.
func (p Page) Update(msg tea.Msg) (Page, tea.Cmd) {
	var cmds []tea.Cmd

	var cmd tea.Cmd
	p.messages, cmd = p.messages.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	p.editor, cmd = p.editor.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return p, tea.Batch(cmds...)
}

// View renders the full chat page.
func (p Page) View() string {
	messagesView := p.messages.View()
	editorView := p.editor.View()

	mainContent := lipgloss.JoinVertical(lipgloss.Left, messagesView, editorView)

	if p.showSidebar && p.width > 80 {
		sidebarView := p.sidebar.View()
		return lipgloss.JoinHorizontal(lipgloss.Top, mainContent, sidebarView)
	}

	return mainContent
}

// Init initializes the chat page.
func (p Page) Init() tea.Cmd {
	return tea.Batch(p.messages.Init(), p.editor.Init())
}

// FocusEditor gives focus to the editor.
func (p *Page) FocusEditor() {
	p.editor.Focus()
}

// BlurEditor removes focus from the editor.
func (p *Page) BlurEditor() {
	p.editor.Blur()
}

// Clear resets the chat page for a new session.
func (p *Page) Clear() {
	p.messages.Clear()
	p.editor.Reset()
}

// ToggleSidebar shows/hides the sidebar.
func (p *Page) ToggleSidebar() {
	p.showSidebar = !p.showSidebar
	p.SetSize(p.width, p.height)
}

// SetSessionInfo updates sidebar session info.
func (p *Page) SetSessionInfo(name, id string, turns int) {
	p.sidebar.SetSession(name, id, turns)
}
