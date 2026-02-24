package dialog

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johmara/openclaude/internal/tui/theme"
)

// QuitConfirmedMsg is emitted when quit is confirmed.
type QuitConfirmedMsg struct{}

// Quit is the quit confirmation dialog (Ctrl+C).
type Quit struct {
	selected int // 0 = cancel, 1 = quit
	width    int
}

// NewQuit creates a quit confirmation dialog.
func NewQuit() Quit {
	return Quit{width: 40}
}

// SetSize adjusts dialog size.
func (q *Quit) SetSize(w, _ int) {
	q.width = w / 3
	if q.width < 35 {
		q.width = 35
	}
	if q.width > 45 {
		q.width = 45
	}
}

// Update handles input.
func (q Quit) Update(msg tea.Msg) (Quit, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "esc", "n":
			return q, func() tea.Msg { return CloseMsg{} }
		case "enter":
			if q.selected == 1 {
				return q, func() tea.Msg { return QuitConfirmedMsg{} }
			}
			return q, func() tea.Msg { return CloseMsg{} }
		case "y":
			return q, func() tea.Msg { return QuitConfirmedMsg{} }
		case "left", "h":
			q.selected = 0
		case "right", "l":
			q.selected = 1
		case "tab":
			q.selected = (q.selected + 1) % 2
		}
	}
	return q, nil
}

// View renders the quit dialog.
func (q Quit) View() string {
	t := theme.Current()

	var sb strings.Builder

	title := lipgloss.NewStyle().
		Foreground(t.Warning()).
		Bold(true).
		Render("Quit OpenClaude?")
	sb.WriteString(title + "\n\n")

	sb.WriteString(lipgloss.NewStyle().
		Foreground(t.Text()).
		Render("Are you sure you want to quit?") + "\n\n")

	cancelStyle := lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(t.Text()).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.Border())
	quitStyle := lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(t.Text()).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.Border())

	if q.selected == 0 {
		cancelStyle = cancelStyle.
			Bold(true).
			BorderForeground(t.Primary()).
			Foreground(t.Primary())
	} else {
		quitStyle = quitStyle.
			Bold(true).
			BorderForeground(t.Error()).
			Foreground(t.Error())
	}

	buttons := lipgloss.JoinHorizontal(lipgloss.Center,
		cancelStyle.Render("Cancel"),
		"  ",
		quitStyle.Render("Quit"),
	)
	sb.WriteString(buttons)

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.DialogBorder()).
		Background(t.DialogBg()).
		Padding(1, 2).
		Width(q.width).
		Render(sb.String())
}

// Init initializes the dialog.
func (q Quit) Init() tea.Cmd {
	return nil
}
