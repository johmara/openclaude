package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johanmontorfano/openclaude/internal/tui/theme"
)

// T returns the current theme for convenience.
func T() theme.Theme {
	return theme.Current()
}

// BaseStyle returns a clean base style.
func BaseStyle() lipgloss.Style {
	return lipgloss.NewStyle()
}

// Bold returns a bold text style.
func Bold() lipgloss.Style {
	return lipgloss.NewStyle().Bold(true)
}

// Muted returns muted text style.
func Muted() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(T().TextMuted())
}

// Accent returns accent-colored text.
func Accent() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(T().Primary())
}

// SuccessText returns success-colored text.
func SuccessText() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(T().Success())
}

// ErrorText returns error-colored text.
func ErrorText() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(T().Error())
}

// WarningText returns warning-colored text.
func WarningText() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(T().Warning())
}

// InfoText returns info-colored text.
func InfoText() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(T().Info())
}

// BorderedBox returns a bordered box style.
func BorderedBox(focused bool) lipgloss.Style {
	borderColor := T().Border()
	if focused {
		borderColor = T().BorderFocused()
	}
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor)
}

// StatusBar returns the status bar style.
func StatusBar() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(T().StatusBarBg()).
		Foreground(T().StatusBarFg()).
		Padding(0, 1)
}

// DialogBox returns the dialog box style.
func DialogBox() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(T().DialogBg()).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(T().DialogBorder()).
		Padding(1, 2)
}

// SelectedItem returns the selected list item style.
func SelectedItem() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(T().SelectionBg()).
		Foreground(T().SelectionFg()).
		Bold(true)
}

// ToolCallBox returns the tool call container style.
func ToolCallBox() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(T().ToolCallBg()).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(T().ToolCallBorder()).
		Padding(0, 1)
}

// Truncate truncates a string to maxLen, adding ellipsis.
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 1 {
		return IconEllipsis
	}
	return s[:maxLen-1] + IconEllipsis
}
