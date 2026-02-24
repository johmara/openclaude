package app

import (
	"context"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/johanmontorfano/openclaude/internal/claude"
	"github.com/johanmontorfano/openclaude/internal/pubsub"
	"github.com/johanmontorfano/openclaude/internal/session"
	"github.com/johanmontorfano/openclaude/internal/tui"
)

// App wires together the Claude client, pub/sub broker, session manager, and TUI.
type App struct {
	Client   *claude.Client
	Broker   *pubsub.Broker
	Sessions *session.Manager
	cancel   context.CancelFunc
}

// Verify App implements tui.App.
var _ tui.App = (*App)(nil)

// NewApp creates a new application instance.
func NewApp() *App {
	return &App{
		Client:   claude.NewClient(),
		Broker:   pubsub.NewBroker(),
		Sessions: session.NewManager(),
	}
}

// SendMessage spawns a Claude subprocess for the given prompt.
func (a *App) SendMessage(prompt string) error {
	sess := a.Sessions.Active()
	if sess == nil {
		sess = a.Sessions.Create("New Session")
	}

	ctx, cancel := context.WithCancel(context.Background())
	a.cancel = cancel

	opts := claude.RunOptions{
		Prompt:    prompt,
		SessionID: sess.ID,
	}

	ch, err := a.Client.Run(ctx, opts)
	if err != nil {
		cancel()
		return fmt.Errorf("run claude: %w", err)
	}

	sess.UpdatedAt = time.Now()
	sess.Turns++

	go a.Broker.PublishStream(ch)

	return nil
}

// CancelGeneration aborts the current Claude subprocess.
func (a *App) CancelGeneration() {
	if a.cancel != nil {
		a.cancel()
		a.cancel = nil
	}
}

// Subscribe creates a broker subscription.
func (a *App) Subscribe(id string) <-chan claude.Event {
	return a.Broker.Subscribe(id)
}

// Unsubscribe removes a broker subscription.
func (a *App) Unsubscribe(id string) {
	a.Broker.Unsubscribe(id)
}

// CreateSession creates a new session and makes it active.
func (a *App) CreateSession(name string) {
	a.Sessions.Create(name)
}

// SetActiveSession sets the active session by index.
func (a *App) SetActiveSession(idx int) {
	a.Sessions.SetActive(idx)
}

// SessionCount returns the number of sessions.
func (a *App) SessionCount() int {
	return a.Sessions.Count()
}

// ActiveSessionName returns the name of the active session.
func (a *App) ActiveSessionName() string {
	s := a.Sessions.Active()
	if s == nil {
		return ""
	}
	return s.Name
}

// Run starts the TUI application.
func Run() error {
	app := NewApp()
	model := tui.New(app)

	p := tea.NewProgram(model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	return nil
}
