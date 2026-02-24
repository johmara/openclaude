package dialog

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/johanmontorfano/openclaude/internal/tui/theme"
)

// ThemeChangedMsg is emitted when a theme is selected.
type ThemeChangedMsg struct {
	Index int
}

// ThemePicker is the theme picker dialog (Ctrl+T).
type ThemePicker struct {
	list List
}

// NewThemePicker creates a theme picker.
func NewThemePicker() ThemePicker {
	themes := theme.All()
	items := make([]Item, len(themes))
	for i, t := range themes {
		desc := ""
		if i == theme.CurrentIndex() {
			desc = "(current)"
		}
		items[i] = Item{
			Title:       t.Name(),
			Description: desc,
			Value:       fmt.Sprintf("%d", i),
		}
	}

	return ThemePicker{
		list: NewList("Theme", items),
	}
}

// SetSize adjusts dialog size.
func (t *ThemePicker) SetSize(w, h int) {
	dialogW := w * 2 / 3
	if dialogW > 50 {
		dialogW = 50
	}
	dialogH := h * 2 / 3
	if dialogH > 15 {
		dialogH = 15
	}
	t.list.SetSize(dialogW, dialogH)
}

// Update handles input.
func (t ThemePicker) Update(msg tea.Msg) (ThemePicker, tea.Cmd) {
	switch msg := msg.(type) {
	case SelectMsg:
		var idx int
		if _, err := fmt.Sscanf(msg.Value, "%d", &idx); err == nil {
			return t, func() tea.Msg { return ThemeChangedMsg{Index: idx} }
		}
	}

	var cmd tea.Cmd
	t.list, cmd = t.list.Update(msg)
	return t, cmd
}

// View renders the theme picker.
func (t ThemePicker) View() string {
	return t.list.View()
}

// Init initializes the dialog.
func (t ThemePicker) Init() tea.Cmd {
	return t.list.Init()
}
