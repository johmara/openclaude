package claude

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Client manages Claude CLI subprocess interactions.
type Client struct {
	claudePath string
}

// NewClient creates a new Claude CLI client.
func NewClient() *Client {
	path := "claude"
	if p := os.Getenv("CLAUDE_PATH"); p != "" {
		path = p
	}
	return &Client{claudePath: path}
}

// RunOptions configures a Claude subprocess invocation.
type RunOptions struct {
	Prompt    string
	SessionID string
	Model     string
}

// Run spawns a Claude subprocess and streams events to the returned channel.
// The channel is closed when the process finishes.
func (c *Client) Run(ctx context.Context, opts RunOptions) (<-chan Event, error) {
	args := []string{
		"--print", "--output-format", "stream-json",
		"--verbose",
		"--no-input",
		"--allowedTools", "Bash", "Read", "Write", "Edit", "Glob", "Grep", "WebSearch", "WebFetch",
	}

	if opts.SessionID != "" {
		args = append(args, "--resume", opts.SessionID)
	}

	if opts.Model != "" {
		args = append(args, "--model", opts.Model)
	}

	args = append(args, "--prompt", opts.Prompt)

	cmd := exec.CommandContext(ctx, c.claudePath, args...)

	// Unset CLAUDE_CODE env vars to avoid nested session detection
	env := os.Environ()
	filteredEnv := make([]string, 0, len(env))
	for _, e := range env {
		if !strings.HasPrefix(e, "CLAUDE_CODE") && !strings.HasPrefix(e, "CLAUDECODE") {
			filteredEnv = append(filteredEnv, e)
		}
	}
	cmd.Env = filteredEnv

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("stdout pipe: %w", err)
	}

	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start claude: %w", err)
	}

	ch := make(chan Event, 64)
	parser := NewStreamParser(stdout)

	go func() {
		parser.Parse(ch)
		_ = cmd.Wait()
	}()

	return ch, nil
}

// Cancel sends an interrupt to abort the current generation.
func (c *Client) Cancel() {
	// Cancellation is handled via context cancellation in Run
}
