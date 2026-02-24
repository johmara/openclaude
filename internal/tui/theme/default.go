package theme

import "github.com/charmbracelet/lipgloss"

// DefaultTheme is a Nord-inspired dark theme.
type DefaultTheme struct{}

func (t DefaultTheme) Name() string { return "Nord" }

func (t DefaultTheme) Primary() lipgloss.Color         { return lipgloss.Color("#88C0D0") }
func (t DefaultTheme) PrimaryDark() lipgloss.Color      { return lipgloss.Color("#6EA8B8") }
func (t DefaultTheme) PrimaryLight() lipgloss.Color     { return lipgloss.Color("#8FBCBB") }
func (t DefaultTheme) Secondary() lipgloss.Color        { return lipgloss.Color("#81A1C1") }
func (t DefaultTheme) Accent() lipgloss.Color           { return lipgloss.Color("#B48EAD") }

func (t DefaultTheme) Text() lipgloss.Color             { return lipgloss.Color("#D8DEE9") }
func (t DefaultTheme) TextMuted() lipgloss.Color        { return lipgloss.Color("#4C566A") }
func (t DefaultTheme) TextInverse() lipgloss.Color      { return lipgloss.Color("#2E3440") }
func (t DefaultTheme) TextAccent() lipgloss.Color       { return lipgloss.Color("#88C0D0") }

func (t DefaultTheme) Background() lipgloss.Color          { return lipgloss.Color("#2E3440") }
func (t DefaultTheme) BackgroundSecondary() lipgloss.Color  { return lipgloss.Color("#3B4252") }
func (t DefaultTheme) BackgroundDark() lipgloss.Color       { return lipgloss.Color("#272C36") }
func (t DefaultTheme) Surface() lipgloss.Color              { return lipgloss.Color("#434C5E") }
func (t DefaultTheme) SurfaceHighlight() lipgloss.Color     { return lipgloss.Color("#4C566A") }

func (t DefaultTheme) Border() lipgloss.Color           { return lipgloss.Color("#4C566A") }
func (t DefaultTheme) BorderFocused() lipgloss.Color    { return lipgloss.Color("#88C0D0") }
func (t DefaultTheme) BorderMuted() lipgloss.Color      { return lipgloss.Color("#3B4252") }

func (t DefaultTheme) Success() lipgloss.Color          { return lipgloss.Color("#A3BE8C") }
func (t DefaultTheme) Warning() lipgloss.Color          { return lipgloss.Color("#EBCB8B") }
func (t DefaultTheme) Error() lipgloss.Color            { return lipgloss.Color("#BF616A") }
func (t DefaultTheme) Info() lipgloss.Color             { return lipgloss.Color("#5E81AC") }

func (t DefaultTheme) SyntaxKeyword() lipgloss.Color    { return lipgloss.Color("#81A1C1") }
func (t DefaultTheme) SyntaxString() lipgloss.Color     { return lipgloss.Color("#A3BE8C") }
func (t DefaultTheme) SyntaxNumber() lipgloss.Color     { return lipgloss.Color("#B48EAD") }
func (t DefaultTheme) SyntaxComment() lipgloss.Color    { return lipgloss.Color("#616E88") }
func (t DefaultTheme) SyntaxFunction() lipgloss.Color   { return lipgloss.Color("#88C0D0") }
func (t DefaultTheme) SyntaxOperator() lipgloss.Color   { return lipgloss.Color("#81A1C1") }
func (t DefaultTheme) SyntaxType() lipgloss.Color       { return lipgloss.Color("#EBCB8B") }

func (t DefaultTheme) StatusBarBg() lipgloss.Color      { return lipgloss.Color("#3B4252") }
func (t DefaultTheme) StatusBarFg() lipgloss.Color      { return lipgloss.Color("#D8DEE9") }
func (t DefaultTheme) InputBg() lipgloss.Color          { return lipgloss.Color("#3B4252") }
func (t DefaultTheme) InputBorder() lipgloss.Color      { return lipgloss.Color("#4C566A") }
func (t DefaultTheme) SidebarBg() lipgloss.Color        { return lipgloss.Color("#3B4252") }
func (t DefaultTheme) DialogBg() lipgloss.Color         { return lipgloss.Color("#434C5E") }
func (t DefaultTheme) DialogBorder() lipgloss.Color     { return lipgloss.Color("#88C0D0") }
func (t DefaultTheme) SelectionBg() lipgloss.Color      { return lipgloss.Color("#434C5E") }
func (t DefaultTheme) SelectionFg() lipgloss.Color      { return lipgloss.Color("#ECEFF4") }
func (t DefaultTheme) SpinnerColor() lipgloss.Color     { return lipgloss.Color("#88C0D0") }
func (t DefaultTheme) ToolCallBg() lipgloss.Color       { return lipgloss.Color("#3B4252") }
func (t DefaultTheme) ToolCallBorder() lipgloss.Color   { return lipgloss.Color("#4C566A") }
