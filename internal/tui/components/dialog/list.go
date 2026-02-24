package dialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johmara/openclaude/internal/tui/styles"
	"github.com/johmara/openclaude/internal/tui/theme"
)

// Item represents an item in a dialog list.
type Item struct {
	Title       string
	Description string
	Value       string
}

// SelectMsg is emitted when a list item is selected.
type SelectMsg struct {
	Value string
}

// CloseMsg is emitted when a dialog should close.
type CloseMsg struct{}

// List is a reusable filterable list dialog.
type List struct {
	title    string
	items    []Item
	filtered []Item
	selected int
	input    textinput.Model
	width    int
	height   int
	maxVisible int
}

// NewList creates a new list dialog.
func NewList(title string, items []Item) List {
	ti := textinput.New()
	ti.Placeholder = "Filter..."
	ti.Focus()
	ti.CharLimit = 100

	return List{
		title:      title,
		items:      items,
		filtered:   items,
		input:      ti,
		width:      50,
		height:     20,
		maxVisible: 10,
	}
}

// SetSize adjusts dialog dimensions.
func (l *List) SetSize(w, h int) {
	l.width = w
	l.height = h
	l.maxVisible = h - 8
	if l.maxVisible < 3 {
		l.maxVisible = 3
	}
}

// Update handles input.
func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return l, func() tea.Msg { return CloseMsg{} }
		case "enter":
			if len(l.filtered) > 0 && l.selected < len(l.filtered) {
				return l, func() tea.Msg {
					return SelectMsg{Value: l.filtered[l.selected].Value}
				}
			}
		case "up":
			if l.selected > 0 {
				l.selected--
			}
			return l, nil
		case "down":
			if l.selected < len(l.filtered)-1 {
				l.selected++
			}
			return l, nil
		}
	}

	var cmd tea.Cmd
	l.input, cmd = l.input.Update(msg)

	// Filter items
	query := strings.ToLower(l.input.Value())
	if query == "" {
		l.filtered = l.items
	} else {
		l.filtered = nil
		for _, item := range l.items {
			if strings.Contains(strings.ToLower(item.Title), query) ||
				strings.Contains(strings.ToLower(item.Description), query) {
				l.filtered = append(l.filtered, item)
			}
		}
	}

	if l.selected >= len(l.filtered) {
		l.selected = len(l.filtered) - 1
	}
	if l.selected < 0 {
		l.selected = 0
	}

	return l, cmd
}

// View renders the list dialog.
func (l List) View() string {
	t := theme.Current()

	var sb strings.Builder

	// Title
	title := lipgloss.NewStyle().
		Foreground(t.Primary()).
		Bold(true).
		Render(l.title)
	sb.WriteString(title + "\n\n")

	// Filter input
	sb.WriteString(l.input.View() + "\n\n")

	// Items
	visible := l.filtered
	if len(visible) > l.maxVisible {
		start := l.selected - l.maxVisible/2
		if start < 0 {
			start = 0
		}
		end := start + l.maxVisible
		if end > len(visible) {
			end = len(visible)
			start = end - l.maxVisible
			if start < 0 {
				start = 0
			}
		}
		visible = visible[start:end]
	}

	for i, item := range visible {
		actualIdx := i
		// Adjust index for scrolling
		if len(l.filtered) > l.maxVisible {
			start := l.selected - l.maxVisible/2
			if start < 0 {
				start = 0
			}
			actualIdx = start + i
		}

		prefix := "  "
		itemStyle := lipgloss.NewStyle().Foreground(t.Text())
		descStyle := lipgloss.NewStyle().Foreground(t.TextMuted())

		if actualIdx == l.selected {
			prefix = styles.IconArrowRight + " "
			itemStyle = lipgloss.NewStyle().
				Foreground(t.SelectionFg()).
				Background(t.SelectionBg()).
				Bold(true)
			descStyle = lipgloss.NewStyle().
				Foreground(t.TextMuted()).
				Background(t.SelectionBg())
		}

		titleStr := itemStyle.Render(prefix + item.Title)
		if item.Description != "" {
			titleStr += " " + descStyle.Render(item.Description)
		}
		sb.WriteString(titleStr + "\n")
	}

	if len(l.filtered) == 0 {
		sb.WriteString(lipgloss.NewStyle().
			Foreground(t.TextMuted()).
			Render("  No matches found") + "\n")
	}

	// Footer
	sb.WriteString("\n" + lipgloss.NewStyle().
		Foreground(t.TextMuted()).
		Render("↑↓ navigate • enter select • esc close"))

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.DialogBorder()).
		Background(t.DialogBg()).
		Padding(1, 2).
		Width(l.width).
		Render(sb.String())
}

// Init initializes the list.
func (l List) Init() tea.Cmd {
	return textinput.Blink
}
