package theme

import "github.com/charmbracelet/lipgloss"

// CatppuccinTheme implements the Catppuccin Mocha color palette.
type CatppuccinTheme struct{}

func (t CatppuccinTheme) Name() string { return "Catppuccin Mocha" }

func (t CatppuccinTheme) Primary() lipgloss.Color         { return lipgloss.Color("#CBA6F7") }
func (t CatppuccinTheme) PrimaryDark() lipgloss.Color      { return lipgloss.Color("#B4A0D1") }
func (t CatppuccinTheme) PrimaryLight() lipgloss.Color     { return lipgloss.Color("#DCC6FF") }
func (t CatppuccinTheme) Secondary() lipgloss.Color        { return lipgloss.Color("#89B4FA") }
func (t CatppuccinTheme) Accent() lipgloss.Color           { return lipgloss.Color("#F5C2E7") }

func (t CatppuccinTheme) Text() lipgloss.Color             { return lipgloss.Color("#CDD6F4") }
func (t CatppuccinTheme) TextMuted() lipgloss.Color        { return lipgloss.Color("#6C7086") }
func (t CatppuccinTheme) TextInverse() lipgloss.Color      { return lipgloss.Color("#1E1E2E") }
func (t CatppuccinTheme) TextAccent() lipgloss.Color       { return lipgloss.Color("#CBA6F7") }

func (t CatppuccinTheme) Background() lipgloss.Color          { return lipgloss.Color("#1E1E2E") }
func (t CatppuccinTheme) BackgroundSecondary() lipgloss.Color  { return lipgloss.Color("#181825") }
func (t CatppuccinTheme) BackgroundDark() lipgloss.Color       { return lipgloss.Color("#11111B") }
func (t CatppuccinTheme) Surface() lipgloss.Color              { return lipgloss.Color("#313244") }
func (t CatppuccinTheme) SurfaceHighlight() lipgloss.Color     { return lipgloss.Color("#45475A") }

func (t CatppuccinTheme) Border() lipgloss.Color           { return lipgloss.Color("#45475A") }
func (t CatppuccinTheme) BorderFocused() lipgloss.Color    { return lipgloss.Color("#CBA6F7") }
func (t CatppuccinTheme) BorderMuted() lipgloss.Color      { return lipgloss.Color("#313244") }

func (t CatppuccinTheme) Success() lipgloss.Color          { return lipgloss.Color("#A6E3A1") }
func (t CatppuccinTheme) Warning() lipgloss.Color          { return lipgloss.Color("#F9E2AF") }
func (t CatppuccinTheme) Error() lipgloss.Color            { return lipgloss.Color("#F38BA8") }
func (t CatppuccinTheme) Info() lipgloss.Color             { return lipgloss.Color("#89B4FA") }

func (t CatppuccinTheme) SyntaxKeyword() lipgloss.Color    { return lipgloss.Color("#CBA6F7") }
func (t CatppuccinTheme) SyntaxString() lipgloss.Color     { return lipgloss.Color("#A6E3A1") }
func (t CatppuccinTheme) SyntaxNumber() lipgloss.Color     { return lipgloss.Color("#FAB387") }
func (t CatppuccinTheme) SyntaxComment() lipgloss.Color    { return lipgloss.Color("#6C7086") }
func (t CatppuccinTheme) SyntaxFunction() lipgloss.Color   { return lipgloss.Color("#89B4FA") }
func (t CatppuccinTheme) SyntaxOperator() lipgloss.Color   { return lipgloss.Color("#89DCEB") }
func (t CatppuccinTheme) SyntaxType() lipgloss.Color       { return lipgloss.Color("#F9E2AF") }

func (t CatppuccinTheme) StatusBarBg() lipgloss.Color      { return lipgloss.Color("#313244") }
func (t CatppuccinTheme) StatusBarFg() lipgloss.Color      { return lipgloss.Color("#BAC2DE") }
func (t CatppuccinTheme) InputBg() lipgloss.Color          { return lipgloss.Color("#181825") }
func (t CatppuccinTheme) InputBorder() lipgloss.Color      { return lipgloss.Color("#45475A") }
func (t CatppuccinTheme) SidebarBg() lipgloss.Color        { return lipgloss.Color("#181825") }
func (t CatppuccinTheme) DialogBg() lipgloss.Color         { return lipgloss.Color("#313244") }
func (t CatppuccinTheme) DialogBorder() lipgloss.Color     { return lipgloss.Color("#CBA6F7") }
func (t CatppuccinTheme) SelectionBg() lipgloss.Color      { return lipgloss.Color("#45475A") }
func (t CatppuccinTheme) SelectionFg() lipgloss.Color      { return lipgloss.Color("#CDD6F4") }
func (t CatppuccinTheme) SpinnerColor() lipgloss.Color     { return lipgloss.Color("#CBA6F7") }
func (t CatppuccinTheme) ToolCallBg() lipgloss.Color       { return lipgloss.Color("#181825") }
func (t CatppuccinTheme) ToolCallBorder() lipgloss.Color   { return lipgloss.Color("#45475A") }
