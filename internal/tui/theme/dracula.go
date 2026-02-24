package theme

import "github.com/charmbracelet/lipgloss"

// DraculaTheme implements the Dracula color palette.
type DraculaTheme struct{}

func (t DraculaTheme) Name() string { return "Dracula" }

func (t DraculaTheme) Primary() lipgloss.Color         { return lipgloss.Color("#BD93F9") }
func (t DraculaTheme) PrimaryDark() lipgloss.Color      { return lipgloss.Color("#9B6FD7") }
func (t DraculaTheme) PrimaryLight() lipgloss.Color     { return lipgloss.Color("#D6ACFF") }
func (t DraculaTheme) Secondary() lipgloss.Color        { return lipgloss.Color("#8BE9FD") }
func (t DraculaTheme) Accent() lipgloss.Color           { return lipgloss.Color("#FF79C6") }

func (t DraculaTheme) Text() lipgloss.Color             { return lipgloss.Color("#F8F8F2") }
func (t DraculaTheme) TextMuted() lipgloss.Color        { return lipgloss.Color("#6272A4") }
func (t DraculaTheme) TextInverse() lipgloss.Color      { return lipgloss.Color("#282A36") }
func (t DraculaTheme) TextAccent() lipgloss.Color       { return lipgloss.Color("#BD93F9") }

func (t DraculaTheme) Background() lipgloss.Color          { return lipgloss.Color("#282A36") }
func (t DraculaTheme) BackgroundSecondary() lipgloss.Color  { return lipgloss.Color("#21222C") }
func (t DraculaTheme) BackgroundDark() lipgloss.Color       { return lipgloss.Color("#191A21") }
func (t DraculaTheme) Surface() lipgloss.Color              { return lipgloss.Color("#44475A") }
func (t DraculaTheme) SurfaceHighlight() lipgloss.Color     { return lipgloss.Color("#6272A4") }

func (t DraculaTheme) Border() lipgloss.Color           { return lipgloss.Color("#44475A") }
func (t DraculaTheme) BorderFocused() lipgloss.Color    { return lipgloss.Color("#BD93F9") }
func (t DraculaTheme) BorderMuted() lipgloss.Color      { return lipgloss.Color("#383A46") }

func (t DraculaTheme) Success() lipgloss.Color          { return lipgloss.Color("#50FA7B") }
func (t DraculaTheme) Warning() lipgloss.Color          { return lipgloss.Color("#F1FA8C") }
func (t DraculaTheme) Error() lipgloss.Color            { return lipgloss.Color("#FF5555") }
func (t DraculaTheme) Info() lipgloss.Color             { return lipgloss.Color("#8BE9FD") }

func (t DraculaTheme) SyntaxKeyword() lipgloss.Color    { return lipgloss.Color("#FF79C6") }
func (t DraculaTheme) SyntaxString() lipgloss.Color     { return lipgloss.Color("#F1FA8C") }
func (t DraculaTheme) SyntaxNumber() lipgloss.Color     { return lipgloss.Color("#BD93F9") }
func (t DraculaTheme) SyntaxComment() lipgloss.Color    { return lipgloss.Color("#6272A4") }
func (t DraculaTheme) SyntaxFunction() lipgloss.Color   { return lipgloss.Color("#50FA7B") }
func (t DraculaTheme) SyntaxOperator() lipgloss.Color   { return lipgloss.Color("#FF79C6") }
func (t DraculaTheme) SyntaxType() lipgloss.Color       { return lipgloss.Color("#8BE9FD") }

func (t DraculaTheme) StatusBarBg() lipgloss.Color      { return lipgloss.Color("#44475A") }
func (t DraculaTheme) StatusBarFg() lipgloss.Color      { return lipgloss.Color("#F8F8F2") }
func (t DraculaTheme) InputBg() lipgloss.Color          { return lipgloss.Color("#21222C") }
func (t DraculaTheme) InputBorder() lipgloss.Color      { return lipgloss.Color("#44475A") }
func (t DraculaTheme) SidebarBg() lipgloss.Color        { return lipgloss.Color("#21222C") }
func (t DraculaTheme) DialogBg() lipgloss.Color         { return lipgloss.Color("#44475A") }
func (t DraculaTheme) DialogBorder() lipgloss.Color     { return lipgloss.Color("#BD93F9") }
func (t DraculaTheme) SelectionBg() lipgloss.Color      { return lipgloss.Color("#44475A") }
func (t DraculaTheme) SelectionFg() lipgloss.Color      { return lipgloss.Color("#F8F8F2") }
func (t DraculaTheme) SpinnerColor() lipgloss.Color     { return lipgloss.Color("#BD93F9") }
func (t DraculaTheme) ToolCallBg() lipgloss.Color       { return lipgloss.Color("#21222C") }
func (t DraculaTheme) ToolCallBorder() lipgloss.Color   { return lipgloss.Color("#44475A") }
