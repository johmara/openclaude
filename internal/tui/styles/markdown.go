package styles

import (
	"github.com/charmbracelet/glamour"
)

// NewMarkdownRenderer creates a glamour renderer configured for the terminal.
func NewMarkdownRenderer(width int) (*glamour.TermRenderer, error) {
	return glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
}

// RenderMarkdown renders markdown text to styled terminal output.
func RenderMarkdown(text string, width int) string {
	if text == "" {
		return ""
	}
	r, err := NewMarkdownRenderer(width)
	if err != nil {
		return text
	}
	out, err := r.Render(text)
	if err != nil {
		return text
	}
	return out
}
