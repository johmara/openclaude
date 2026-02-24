package pubsub

import "github.com/johmara/openclaude/internal/claude"

// ClaudeEvent wraps a claude.Event for the pub/sub system.
type ClaudeEvent struct {
	Event claude.Event
}
