package dialog

import tea "github.com/charmbracelet/bubbletea"

// CommandMsg is emitted when a command is selected from the palette.
type CommandMsg struct {
	Command string
}

// Commands is the command palette dialog (Ctrl+K).
type Commands struct {
	list List
}

// NewCommands creates a command palette.
func NewCommands() Commands {
	items := []Item{
		{Title: "New Session", Description: "Start a new conversation", Value: "new_session"},
		{Title: "Switch Session", Description: "Switch between sessions", Value: "switch_session"},
		{Title: "Change Theme", Description: "Switch color theme", Value: "change_theme"},
		{Title: "Open File", Description: "Open file picker", Value: "file_picker"},
		{Title: "Toggle Sidebar", Description: "Show/hide sidebar", Value: "toggle_sidebar"},
		{Title: "Clear Chat", Description: "Clear current conversation", Value: "clear_chat"},
		{Title: "Help", Description: "Show keybinding reference", Value: "help"},
		{Title: "Quit", Description: "Exit OpenClaude", Value: "quit"},
	}

	return Commands{
		list: NewList("Command Palette", items),
	}
}

// SetSize adjusts dialog size.
func (c *Commands) SetSize(w, h int) {
	dialogW := w * 2 / 3
	if dialogW > 60 {
		dialogW = 60
	}
	dialogH := h * 2 / 3
	if dialogH > 20 {
		dialogH = 20
	}
	c.list.SetSize(dialogW, dialogH)
}

// Update handles input.
func (c Commands) Update(msg tea.Msg) (Commands, tea.Cmd) {
	switch msg := msg.(type) {
	case SelectMsg:
		return c, func() tea.Msg { return CommandMsg{Command: msg.Value} }
	}

	var cmd tea.Cmd
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

// View renders the command palette.
func (c Commands) View() string {
	return c.list.View()
}

// Init initializes the dialog.
func (c Commands) Init() tea.Cmd {
	return c.list.Init()
}
