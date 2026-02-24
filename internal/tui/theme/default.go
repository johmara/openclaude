package theme

import "github.com/charmbracelet/lipgloss"

// DefaultTheme is the Claude-branded dark theme.
type DefaultTheme struct{}

func (t DefaultTheme) Name() string { return "Claude Dark" }

func (t DefaultTheme) Primary() lipgloss.Color         { return lipgloss.Color("#D97757") }
func (t DefaultTheme) PrimaryDark() lipgloss.Color      { return lipgloss.Color("#B85C3A") }
func (t DefaultTheme) PrimaryLight() lipgloss.Color     { return lipgloss.Color("#E89B7D") }
func (t DefaultTheme) Secondary() lipgloss.Color        { return lipgloss.Color("#6E8898") }
func (t DefaultTheme) Accent() lipgloss.Color           { return lipgloss.Color("#C9A96E") }

func (t DefaultTheme) Text() lipgloss.Color             { return lipgloss.Color("#E0E0E0") }
func (t DefaultTheme) TextMuted() lipgloss.Color        { return lipgloss.Color("#808080") }
func (t DefaultTheme) TextInverse() lipgloss.Color      { return lipgloss.Color("#1A1A2E") }
func (t DefaultTheme) TextAccent() lipgloss.Color       { return lipgloss.Color("#D97757") }

func (t DefaultTheme) Background() lipgloss.Color          { return lipgloss.Color("#1A1A2E") }
func (t DefaultTheme) BackgroundSecondary() lipgloss.Color  { return lipgloss.Color("#1E1E36") }
func (t DefaultTheme) BackgroundDark() lipgloss.Color       { return lipgloss.Color("#141424") }
func (t DefaultTheme) Surface() lipgloss.Color              { return lipgloss.Color("#252540") }
func (t DefaultTheme) SurfaceHighlight() lipgloss.Color     { return lipgloss.Color("#2D2D4A") }

func (t DefaultTheme) Border() lipgloss.Color           { return lipgloss.Color("#3A3A5C") }
func (t DefaultTheme) BorderFocused() lipgloss.Color    { return lipgloss.Color("#D97757") }
func (t DefaultTheme) BorderMuted() lipgloss.Color      { return lipgloss.Color("#2A2A44") }

func (t DefaultTheme) Success() lipgloss.Color          { return lipgloss.Color("#73C991") }
func (t DefaultTheme) Warning() lipgloss.Color          { return lipgloss.Color("#E5C07B") }
func (t DefaultTheme) Error() lipgloss.Color            { return lipgloss.Color("#E06C75") }
func (t DefaultTheme) Info() lipgloss.Color             { return lipgloss.Color("#61AFEF") }

func (t DefaultTheme) SyntaxKeyword() lipgloss.Color    { return lipgloss.Color("#C678DD") }
func (t DefaultTheme) SyntaxString() lipgloss.Color     { return lipgloss.Color("#98C379") }
func (t DefaultTheme) SyntaxNumber() lipgloss.Color     { return lipgloss.Color("#D19A66") }
func (t DefaultTheme) SyntaxComment() lipgloss.Color    { return lipgloss.Color("#5C6370") }
func (t DefaultTheme) SyntaxFunction() lipgloss.Color   { return lipgloss.Color("#61AFEF") }
func (t DefaultTheme) SyntaxOperator() lipgloss.Color   { return lipgloss.Color("#56B6C2") }
func (t DefaultTheme) SyntaxType() lipgloss.Color       { return lipgloss.Color("#E5C07B") }

func (t DefaultTheme) StatusBarBg() lipgloss.Color      { return lipgloss.Color("#252540") }
func (t DefaultTheme) StatusBarFg() lipgloss.Color      { return lipgloss.Color("#B0B0B0") }
func (t DefaultTheme) InputBg() lipgloss.Color          { return lipgloss.Color("#1E1E36") }
func (t DefaultTheme) InputBorder() lipgloss.Color      { return lipgloss.Color("#3A3A5C") }
func (t DefaultTheme) SidebarBg() lipgloss.Color        { return lipgloss.Color("#1E1E36") }
func (t DefaultTheme) DialogBg() lipgloss.Color         { return lipgloss.Color("#252540") }
func (t DefaultTheme) DialogBorder() lipgloss.Color     { return lipgloss.Color("#D97757") }
func (t DefaultTheme) SelectionBg() lipgloss.Color      { return lipgloss.Color("#3A3A5C") }
func (t DefaultTheme) SelectionFg() lipgloss.Color      { return lipgloss.Color("#E0E0E0") }
func (t DefaultTheme) SpinnerColor() lipgloss.Color     { return lipgloss.Color("#D97757") }
func (t DefaultTheme) ToolCallBg() lipgloss.Color       { return lipgloss.Color("#1E1E36") }
func (t DefaultTheme) ToolCallBorder() lipgloss.Color   { return lipgloss.Color("#3A3A5C") }
