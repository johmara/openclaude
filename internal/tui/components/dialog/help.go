package dialog

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johmara/openclaude/internal/tui/theme"
)

// Help shows the keybinding reference dialog (Ctrl+?).
type Help struct {
	width  int
	height int
}

// NewHelp creates a help dialog.
func NewHelp() Help {
	return Help{width: 50, height: 22}
}

// SetSize adjusts dialog size.
func (h *Help) SetSize(w, ht int) {
	h.width = w * 2 / 3
	if h.width > 55 {
		h.width = 55
	}
	h.height = ht * 2 / 3
	if h.height > 24 {
		h.height = 24
	}
}

type helpBinding struct {
	key  string
	desc string
}

var bindings = []helpBinding{
	{"Enter", "Send message"},
	{"Esc", "Cancel generation / close dialog"},
	{"Ctrl+C", "Quit"},
	{"Ctrl+K", "Command palette"},
	{"Ctrl+S", "Session switcher"},
	{"Ctrl+T", "Theme picker"},
	{"Ctrl+F", "File picker"},
	{"Ctrl+?", "This help dialog"},
	{"Ctrl+N", "New session"},
	{"PgUp/PgDn", "Scroll messages"},
}

// Update handles input.
func (h Help) Update(msg tea.Msg) (Help, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.String() == "esc" || msg.String() == "enter" || msg.String() == "q" {
			return h, func() tea.Msg { return CloseMsg{} }
		}
	}
	return h, nil
}

// View renders the help dialog.
func (h Help) View() string {
	t := theme.Current()

	var sb strings.Builder

	title := lipgloss.NewStyle().
		Foreground(t.Primary()).
		Bold(true).
		Render("Keybindings")
	sb.WriteString(title + "\n\n")

	keyStyle := lipgloss.NewStyle().
		Foreground(t.Primary()).
		Bold(true).
		Width(14)
	descStyle := lipgloss.NewStyle().
		Foreground(t.Text())

	for _, b := range bindings {
		sb.WriteString(keyStyle.Render(b.key) + descStyle.Render(b.desc) + "\n")
	}

	sb.WriteString("\n" + lipgloss.NewStyle().
		Foreground(t.TextMuted()).
		Render("Press Esc to close"))

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.DialogBorder()).
		Background(t.DialogBg()).
		Padding(1, 2).
		Width(h.width).
		Render(sb.String())
}

// Init initializes the help dialog.
func (h Help) Init() tea.Cmd {
	return nil
}
