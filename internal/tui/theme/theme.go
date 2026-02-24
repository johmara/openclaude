package theme

import "github.com/charmbracelet/lipgloss"

// Theme defines the color palette for the entire application.
type Theme interface {
	Name() string

	// Primary colors
	Primary() lipgloss.Color
	PrimaryDark() lipgloss.Color
	PrimaryLight() lipgloss.Color
	Secondary() lipgloss.Color
	Accent() lipgloss.Color

	// Text colors
	Text() lipgloss.Color
	TextMuted() lipgloss.Color
	TextInverse() lipgloss.Color
	TextAccent() lipgloss.Color

	// Background colors
	Background() lipgloss.Color
	BackgroundSecondary() lipgloss.Color
	BackgroundDark() lipgloss.Color
	Surface() lipgloss.Color
	SurfaceHighlight() lipgloss.Color

	// Border colors
	Border() lipgloss.Color
	BorderFocused() lipgloss.Color
	BorderMuted() lipgloss.Color

	// Status colors
	Success() lipgloss.Color
	Warning() lipgloss.Color
	Error() lipgloss.Color
	Info() lipgloss.Color

	// Syntax highlighting
	SyntaxKeyword() lipgloss.Color
	SyntaxString() lipgloss.Color
	SyntaxNumber() lipgloss.Color
	SyntaxComment() lipgloss.Color
	SyntaxFunction() lipgloss.Color
	SyntaxOperator() lipgloss.Color
	SyntaxType() lipgloss.Color

	// Component-specific
	StatusBarBg() lipgloss.Color
	StatusBarFg() lipgloss.Color
	InputBg() lipgloss.Color
	InputBorder() lipgloss.Color
	SidebarBg() lipgloss.Color
	DialogBg() lipgloss.Color
	DialogBorder() lipgloss.Color
	SelectionBg() lipgloss.Color
	SelectionFg() lipgloss.Color
	SpinnerColor() lipgloss.Color
	ToolCallBg() lipgloss.Color
	ToolCallBorder() lipgloss.Color
}
