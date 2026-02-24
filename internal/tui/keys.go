package tui

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the global keybindings.
type KeyMap struct {
	Send          key.Binding
	Cancel        key.Binding
	Quit          key.Binding
	CommandPalette key.Binding
	SessionSwitch key.Binding
	ThemePicker   key.Binding
	FilePicker    key.Binding
	Help          key.Binding
	NewSession    key.Binding
	ScrollUp      key.Binding
	ScrollDown    key.Binding
	PageUp        key.Binding
	PageDown      key.Binding
	Escape        key.Binding
}

// DefaultKeyMap returns the default keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Send: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "send message"),
		),
		Cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel generation"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit"),
		),
		CommandPalette: key.NewBinding(
			key.WithKeys("ctrl+k"),
			key.WithHelp("ctrl+k", "command palette"),
		),
		SessionSwitch: key.NewBinding(
			key.WithKeys("ctrl+s"),
			key.WithHelp("ctrl+s", "sessions"),
		),
		ThemePicker: key.NewBinding(
			key.WithKeys("ctrl+t"),
			key.WithHelp("ctrl+t", "themes"),
		),
		FilePicker: key.NewBinding(
			key.WithKeys("ctrl+f"),
			key.WithHelp("ctrl+f", "file picker"),
		),
		Help: key.NewBinding(
			key.WithKeys("ctrl+/"),
			key.WithHelp("ctrl+?", "help"),
		),
		NewSession: key.NewBinding(
			key.WithKeys("ctrl+n"),
			key.WithHelp("ctrl+n", "new session"),
		),
		ScrollUp: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "scroll up"),
		),
		ScrollDown: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "scroll down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("pgup"),
			key.WithHelp("pgup", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("pgdown"),
			key.WithHelp("pgdn", "page down"),
		),
		Escape: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "close dialog"),
		),
	}
}
