package session

import (
	"time"

	"github.com/google/uuid"
)

// Session holds metadata about a Claude conversation session.
type Session struct {
	ID        string
	Name      string
	Model     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Cost      float64
	Turns     int
}

// New creates a new session with a generated ID.
func New(name string) *Session {
	now := time.Now()
	return &Session{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Manager tracks active sessions.
type Manager struct {
	sessions []*Session
	active   int
}

// NewManager creates a session manager.
func NewManager() *Manager {
	return &Manager{}
}

// Create adds a new session and makes it active.
func (m *Manager) Create(name string) *Session {
	s := New(name)
	m.sessions = append(m.sessions, s)
	m.active = len(m.sessions) - 1
	return s
}

// Active returns the current session, or nil if none.
func (m *Manager) Active() *Session {
	if len(m.sessions) == 0 {
		return nil
	}
	return m.sessions[m.active]
}

// SetActive sets the active session by index.
func (m *Manager) SetActive(idx int) {
	if idx >= 0 && idx < len(m.sessions) {
		m.active = idx
	}
}

// All returns all sessions.
func (m *Manager) All() []*Session {
	return m.sessions
}

// Count returns the number of sessions.
func (m *Manager) Count() int {
	return len(m.sessions)
}
