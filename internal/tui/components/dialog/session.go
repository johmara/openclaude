package dialog

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/johanmontorfano/openclaude/internal/session"
)

// SessionSelectedMsg is emitted when a session is selected.
type SessionSelectedMsg struct {
	Index int
}

// SessionSwitcher is the session switcher dialog (Ctrl+S).
type SessionSwitcher struct {
	list     List
	sessions *session.Manager
}

// NewSessionSwitcher creates a session switcher.
func NewSessionSwitcher(sessions *session.Manager) SessionSwitcher {
	items := buildSessionItems(sessions)
	return SessionSwitcher{
		list:     NewList("Sessions", items),
		sessions: sessions,
	}
}

func buildSessionItems(sessions *session.Manager) []Item {
	all := sessions.All()
	items := make([]Item, len(all))
	for i, s := range all {
		desc := s.CreatedAt.Format("Jan 2 15:04")
		items[i] = Item{
			Title:       s.Name,
			Description: desc,
			Value:       s.ID,
		}
	}
	return items
}

// Refresh rebuilds the session list.
func (s *SessionSwitcher) Refresh() {
	items := buildSessionItems(s.sessions)
	s.list = NewList("Sessions", items)
}

// SetSize adjusts dialog size.
func (s *SessionSwitcher) SetSize(w, h int) {
	dialogW := w * 2 / 3
	if dialogW > 55 {
		dialogW = 55
	}
	dialogH := h * 2 / 3
	if dialogH > 20 {
		dialogH = 20
	}
	s.list.SetSize(dialogW, dialogH)
}

// Update handles input.
func (s SessionSwitcher) Update(msg tea.Msg) (SessionSwitcher, tea.Cmd) {
	switch msg := msg.(type) {
	case SelectMsg:
		// Find session index by ID
		for i, sess := range s.sessions.All() {
			if sess.ID == msg.Value {
				return s, func() tea.Msg { return SessionSelectedMsg{Index: i} }
			}
		}
	}

	var cmd tea.Cmd
	s.list, cmd = s.list.Update(msg)
	return s, cmd
}

// View renders the session switcher.
func (s SessionSwitcher) View() string {
	return s.list.View()
}

// Init initializes the dialog.
func (s SessionSwitcher) Init() tea.Cmd {
	return s.list.Init()
}
