package chat

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/johanmontorfano/openclaude/internal/tui/styles"
	"github.com/johanmontorfano/openclaude/internal/tui/theme"
)

// Sidebar shows session info and metadata.
type Sidebar struct {
	width       int
	height      int
	model       string
	cost        float64
	tokens      int
	sessionName string
	sessionID   string
	turns       int
}

// NewSidebar creates a new sidebar.
func NewSidebar() Sidebar {
	return Sidebar{}
}

// SetSize updates sidebar dimensions.
func (s *Sidebar) SetSize(w, h int) {
	s.width = w
	s.height = h
}

// SetModel sets the current model name.
func (s *Sidebar) SetModel(model string) {
	s.model = model
}

// SetCost sets the total cost.
func (s *Sidebar) SetCost(cost float64) {
	s.cost = cost
}

// SetTokens sets total tokens used.
func (s *Sidebar) SetTokens(tokens int) {
	s.tokens = tokens
}

// SetSession sets session display info.
func (s *Sidebar) SetSession(name, id string, turns int) {
	s.sessionName = name
	s.sessionID = id
	s.turns = turns
}

// View renders the sidebar.
func (s Sidebar) View() string {
	t := theme.Current()
	innerWidth := s.width - 4

	if innerWidth < 10 {
		return ""
	}

	var sb strings.Builder

	// Title
	title := lipgloss.NewStyle().
		Foreground(t.Primary()).
		Bold(true).
		Width(innerWidth).
		Render("Session Info")
	sb.WriteString(title + "\n")
	sb.WriteString(lipgloss.NewStyle().
		Foreground(t.BorderMuted()).
		Width(innerWidth).
		Render(strings.Repeat("─", innerWidth)) + "\n\n")

	// Session info
	if s.sessionName != "" {
		sb.WriteString(s.renderField("Session", styles.Truncate(s.sessionName, innerWidth-10), t))
		if s.sessionID != "" && len(s.sessionID) > 8 {
			sb.WriteString(s.renderField("ID", s.sessionID[:8]+"...", t))
		}
		sb.WriteString(s.renderField("Turns", fmt.Sprintf("%d", s.turns), t))
		sb.WriteString("\n")
	}

	// Model
	if s.model != "" {
		sb.WriteString(s.renderField("Model", styles.Truncate(s.model, innerWidth-10), t))
	}

	// Cost
	if s.cost > 0 {
		sb.WriteString(s.renderField("Cost", fmt.Sprintf("$%.4f", s.cost), t))
	}

	// Tokens
	if s.tokens > 0 {
		sb.WriteString(s.renderField("Tokens", fmt.Sprintf("%d", s.tokens), t))
	}

	// Keybindings help at bottom
	sb.WriteString("\n")
	sb.WriteString(lipgloss.NewStyle().
		Foreground(t.BorderMuted()).
		Width(innerWidth).
		Render(strings.Repeat("─", innerWidth)) + "\n")
	sb.WriteString(s.renderKey("ctrl+k", "commands", t))
	sb.WriteString(s.renderKey("ctrl+s", "sessions", t))
	sb.WriteString(s.renderKey("ctrl+t", "theme", t))
	sb.WriteString(s.renderKey("ctrl+f", "files", t))
	sb.WriteString(s.renderKey("ctrl+?", "help", t))
	sb.WriteString(s.renderKey("ctrl+n", "new session", t))

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.Border()).
		Width(s.width - 2).
		Height(s.height - 2)

	return style.Render(sb.String())
}

func (s Sidebar) renderField(label, value string, t theme.Theme) string {
	l := lipgloss.NewStyle().Foreground(t.TextMuted()).Render(label + ": ")
	v := lipgloss.NewStyle().Foreground(t.Text()).Render(value)
	return l + v + "\n"
}

func (s Sidebar) renderKey(key, desc string, t theme.Theme) string {
	k := lipgloss.NewStyle().Foreground(t.Primary()).Bold(true).Render(key)
	d := lipgloss.NewStyle().Foreground(t.TextMuted()).Render(" " + desc)
	return k + d + "\n"
}
