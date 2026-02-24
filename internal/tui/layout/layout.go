package layout

import tea "github.com/charmbracelet/bubbletea"

// Focusable is a component that can receive/lose focus.
type Focusable interface {
	Focus()
	Blur()
	Focused() bool
}

// Sizeable is a component that can be resized.
type Sizeable interface {
	SetSize(width, height int)
}

// Component combines all common component interfaces.
type Component interface {
	tea.Model
	Sizeable
}
