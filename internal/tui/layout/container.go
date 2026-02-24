package layout

import "github.com/charmbracelet/lipgloss"

// Container wraps content with optional border and padding.
type Container struct {
	style   lipgloss.Style
	width   int
	height  int
	title   string
}

// NewContainer creates a new container.
func NewContainer() *Container {
	return &Container{
		style: lipgloss.NewStyle(),
	}
}

// WithBorder adds a rounded border.
func (c *Container) WithBorder(color lipgloss.Color) *Container {
	c.style = c.style.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(color)
	return c
}

// WithPadding adds padding.
func (c *Container) WithPadding(top, right, bottom, left int) *Container {
	c.style = c.style.Padding(top, right, bottom, left)
	return c
}

// WithSize sets dimensions.
func (c *Container) WithSize(w, h int) *Container {
	c.width = w
	c.height = h
	return c
}

// WithTitle sets the container title shown in the border.
func (c *Container) WithTitle(title string) *Container {
	c.title = title
	return c
}

// WithBackground sets the background color.
func (c *Container) WithBackground(color lipgloss.Color) *Container {
	c.style = c.style.Background(color)
	return c
}

// Render wraps content with the container style.
func (c *Container) Render(content string) string {
	s := c.style
	if c.width > 0 {
		s = s.Width(c.width)
	}
	if c.height > 0 {
		s = s.Height(c.height)
	}

	rendered := s.Render(content)

	if c.title != "" {
		// Replace top-left corner area with title
		titleStyle := lipgloss.NewStyle().
			Foreground(c.style.GetBorderBottomForeground()).
			Bold(true)
		titleStr := " " + titleStyle.Render(c.title) + " "
		lines := splitLines(rendered)
		if len(lines) > 0 {
			runes := []rune(lines[0])
			if len(runes) > 4 {
				titleRunes := []rune(titleStr)
				insertAt := 2
				for i, r := range titleRunes {
					pos := insertAt + i
					if pos < len(runes) {
						runes[pos] = r
					}
				}
				lines[0] = string(runes)
			}
			rendered = joinLines(lines)
		}
	}

	return rendered
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

func joinLines(lines []string) string {
	result := ""
	for i, l := range lines {
		if i > 0 {
			result += "\n"
		}
		result += l
	}
	return result
}
