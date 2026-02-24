package status

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/johanmontorfano/openclaude/internal/tui/theme"
)

// Bar is the bottom status bar.
type Bar struct {
	width     int
	model     string
	cost      float64
	tokens    int
	streaming bool
	message   string
	session   string
}

// New creates a new status bar.
func New() Bar {
	return Bar{}
}

// SetSize updates the status bar width.
func (b *Bar) SetSize(w int) {
	b.width = w
}

// SetModel sets the model name display.
func (b *Bar) SetModel(model string) {
	b.model = model
}

// SetCost sets the cost display.
func (b *Bar) SetCost(cost float64) {
	b.cost = cost
}

// SetTokens sets the token count display.
func (b *Bar) SetTokens(tokens int) {
	b.tokens = tokens
}

// SetStreaming sets the streaming state.
func (b *Bar) SetStreaming(streaming bool) {
	b.streaming = streaming
}

// SetMessage sets a temporary status message.
func (b *Bar) SetMessage(msg string) {
	b.message = msg
}

// SetSession sets the session name display.
func (b *Bar) SetSession(name string) {
	b.session = name
}

// View renders the status bar.
func (b Bar) View() string {
	t := theme.Current()
	style := lipgloss.NewStyle().
		Background(t.StatusBarBg()).
		Foreground(t.StatusBarFg()).
		Width(b.width)

	// Left side: session + message
	var left []string
	if b.session != "" {
		left = append(left, lipgloss.NewStyle().
			Foreground(t.Primary()).
			Bold(true).
			Background(t.StatusBarBg()).
			Render(" "+b.session))
	}
	if b.streaming {
		left = append(left, lipgloss.NewStyle().
			Foreground(t.Warning()).
			Background(t.StatusBarBg()).
			Render("● streaming"))
	}
	if b.message != "" {
		left = append(left, lipgloss.NewStyle().
			Foreground(t.TextMuted()).
			Background(t.StatusBarBg()).
			Render(b.message))
	}

	// Right side: model + cost + tokens
	var right []string
	if b.tokens > 0 {
		right = append(right, lipgloss.NewStyle().
			Foreground(t.TextMuted()).
			Background(t.StatusBarBg()).
			Render(fmt.Sprintf("%dt", b.tokens)))
	}
	if b.cost > 0 {
		right = append(right, lipgloss.NewStyle().
			Foreground(t.Success()).
			Background(t.StatusBarBg()).
			Render(fmt.Sprintf("$%.4f", b.cost)))
	}
	if b.model != "" {
		right = append(right, lipgloss.NewStyle().
			Foreground(t.Info()).
			Background(t.StatusBarBg()).
			Render(b.model+" "))
	}

	leftStr := strings.Join(left, " │ ")
	rightStr := strings.Join(right, " │ ")

	// Calculate padding between left and right
	leftW := lipgloss.Width(leftStr)
	rightW := lipgloss.Width(rightStr)
	padding := b.width - leftW - rightW
	if padding < 1 {
		padding = 1
	}

	content := leftStr + strings.Repeat(" ", padding) + rightStr

	return style.Render(content)
}
