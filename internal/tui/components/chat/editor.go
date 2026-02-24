package chat

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johanmontorfano/openclaude/internal/tui/theme"
)

// SendMsg is emitted when the user presses Enter to send a message.
type SendMsg struct {
	Text string
}

// Editor wraps a textarea for user input.
type Editor struct {
	textarea textarea.Model
	width    int
	height   int
	focused  bool
}

// NewEditor creates a new message input editor.
func NewEditor() Editor {
	ta := textarea.New()
	ta.Placeholder = "Send a message... (Enter to send)"
	ta.ShowLineNumbers = false
	ta.CharLimit = 0
	ta.SetHeight(3)
	ta.Focus()

	return Editor{
		textarea: ta,
		focused:  true,
	}
}

// Focus gives focus to the editor.
func (e *Editor) Focus() {
	e.focused = true
	e.textarea.Focus()
}

// Blur removes focus from the editor.
func (e *Editor) Blur() {
	e.focused = false
	e.textarea.Blur()
}

// Focused returns whether the editor has focus.
func (e Editor) Focused() bool {
	return e.focused
}

// SetSize updates the editor dimensions.
func (e *Editor) SetSize(w, h int) {
	e.width = w
	e.height = h
	e.textarea.SetWidth(w - 4) // Account for padding/borders
	e.textarea.SetHeight(h - 2)
}

// Value returns the current text.
func (e Editor) Value() string {
	return e.textarea.Value()
}

// Reset clears the editor.
func (e *Editor) Reset() {
	e.textarea.Reset()
}

// Update handles input events.
func (e Editor) Update(msg tea.Msg) (Editor, tea.Cmd) {
	if !e.focused {
		return e, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			text := e.textarea.Value()
			if text != "" {
				e.textarea.Reset()
				return e, func() tea.Msg { return SendMsg{Text: text} }
			}
			return e, nil
		}
	}

	var cmd tea.Cmd
	e.textarea, cmd = e.textarea.Update(msg)
	return e, cmd
}

// View renders the editor.
func (e Editor) View() string {
	t := theme.Current()

	borderColor := t.Border()
	if e.focused {
		borderColor = t.BorderFocused()
	}

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Width(e.width - 2).
		Padding(0, 1)

	return style.Render(e.textarea.View())
}

// Init is the Bubble Tea init function.
func (e Editor) Init() tea.Cmd {
	return textarea.Blink
}
