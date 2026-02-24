package chat

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johmara/openclaude/internal/claude"
	"github.com/johmara/openclaude/internal/tui/styles"
	"github.com/johmara/openclaude/internal/tui/theme"
)

// ToolCall tracks a tool call's state.
type ToolCall struct {
	Name      string
	ID        string
	Input     strings.Builder
	Result    string
	IsError   bool
	Done      bool
	BlockIdx  int
}

// ChatMessage represents a message in the conversation.
type ChatMessage struct {
	Role    string // "user" or "assistant"
	Content string
}

// Messages manages the scrollable message display area.
type Messages struct {
	viewport  viewport.Model
	spinner   spinner.Model
	width     int
	height    int

	// Accumulated state
	messages   []ChatMessage
	currentText strings.Builder
	toolCalls  []*ToolCall
	streaming  bool
	lastRender string
}

// NewMessages creates a new messages view.
func NewMessages() Messages {
	vp := viewport.New(80, 20)
	vp.SetContent("")

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#D97757"))

	return Messages{
		viewport: vp,
		spinner:  sp,
	}
}

// SetSize updates viewport dimensions.
func (m *Messages) SetSize(w, h int) {
	m.width = w
	m.height = h
	m.viewport.Width = w - 2
	m.viewport.Height = h - 2
	m.rerender()
}

// AddUserMessage adds a user message.
func (m *Messages) AddUserMessage(text string) {
	m.messages = append(m.messages, ChatMessage{Role: "user", Content: text})
	m.currentText.Reset()
	m.toolCalls = nil
	m.streaming = true
	m.rerender()
}

// HandleEvent processes a claude event and updates the display.
func (m *Messages) HandleEvent(evt claude.Event) {
	switch evt.Type {
	case claude.EventTextDelta:
		m.currentText.WriteString(evt.Text)
		m.rerender()

	case claude.EventToolStart:
		tc := &ToolCall{
			Name:     evt.ToolName,
			ID:       evt.ToolID,
			BlockIdx: evt.BlockIdx,
		}
		m.toolCalls = append(m.toolCalls, tc)
		m.rerender()

	case claude.EventToolInputDelta:
		for _, tc := range m.toolCalls {
			if tc.BlockIdx == evt.BlockIdx {
				tc.Input.WriteString(evt.PartialJSON)
				break
			}
		}
		m.rerender()

	case claude.EventToolEnd:
		for _, tc := range m.toolCalls {
			if tc.BlockIdx == evt.BlockIdx {
				tc.Done = true
				break
			}
		}
		m.rerender()

	case claude.EventToolResult:
		for _, tc := range m.toolCalls {
			if tc.ID == evt.ToolUseID {
				tc.Result = evt.Content
				tc.IsError = evt.IsError
				tc.Done = true
				break
			}
		}
		m.rerender()

	case claude.EventAssistantMessage:
		// Finalize tool inputs from the complete message
		if evt.Message != nil {
			for _, part := range evt.Message.Content {
				if part.Type == "tool_use" {
					for _, tc := range m.toolCalls {
						if tc.ID == part.ID {
							if input, ok := part.Input.(map[string]any); ok {
								tc.Input.Reset()
								tc.Input.WriteString(formatToolInput(tc.Name, input))
							}
							break
						}
					}
				}
			}
		}
		m.rerender()

	case claude.EventResult:
		// Finalize the assistant message
		if m.currentText.Len() > 0 || len(m.toolCalls) > 0 {
			m.messages = append(m.messages, ChatMessage{
				Role:    "assistant",
				Content: m.currentText.String(),
			})
		}
		m.streaming = false
		m.rerender()
	}
}

// rerender rebuilds the viewport content from all messages.
func (m *Messages) rerender() {
	var sb strings.Builder
	t := theme.Current()
	contentWidth := m.width - 6

	if contentWidth < 20 {
		contentWidth = 20
	}

	for _, msg := range m.messages {
		switch msg.Role {
		case "user":
			header := lipgloss.NewStyle().
				Foreground(t.Primary()).
				Bold(true).
				Render(styles.IconUser + " You")
			sb.WriteString(header + "\n")
			sb.WriteString(msg.Content + "\n\n")

		case "assistant":
			header := lipgloss.NewStyle().
				Foreground(t.Secondary()).
				Bold(true).
				Render(styles.IconAssistant + " Claude")
			sb.WriteString(header + "\n")
			rendered := styles.RenderMarkdown(msg.Content, contentWidth)
			sb.WriteString(rendered + "\n")
		}
	}

	// Render streaming content
	if m.streaming {
		if m.currentText.Len() > 0 || len(m.toolCalls) > 0 {
			header := lipgloss.NewStyle().
				Foreground(t.Secondary()).
				Bold(true).
				Render(styles.IconAssistant + " Claude")
			sb.WriteString(header + "\n")
		}

		if m.currentText.Len() > 0 {
			rendered := styles.RenderMarkdown(m.currentText.String(), contentWidth)
			sb.WriteString(rendered)
		}

		// Render tool calls
		for _, tc := range m.toolCalls {
			sb.WriteString(m.renderToolCall(tc))
			sb.WriteString("\n")
		}
	}

	if sb.Len() == 0 {
		welcome := lipgloss.NewStyle().
			Foreground(t.TextMuted()).
			Align(lipgloss.Center).
			Width(m.width - 4).
			Render("\n\n\n  Welcome to OpenClaude\n\n  Type a message to start a conversation with Claude.\n")
		sb.WriteString(welcome)
	}

	content := sb.String()
	if content != m.lastRender {
		m.lastRender = content
		m.viewport.SetContent(content)
		m.viewport.GotoBottom()
	}
}

// renderToolCall renders a single tool call.
func (m *Messages) renderToolCall(tc *ToolCall) string {
	t := theme.Current()
	var sb strings.Builder

	toolLabel := formatToolLabel(tc.Name, tc.Input.String())

	if tc.Done && tc.Result != "" {
		// Completed tool call
		icon := styles.IconSuccess
		iconStyle := lipgloss.NewStyle().Foreground(t.Success())
		if tc.IsError {
			icon = styles.IconError
			iconStyle = lipgloss.NewStyle().Foreground(t.Error())
		}

		header := fmt.Sprintf("  %s %s: %s",
			iconStyle.Render(icon),
			lipgloss.NewStyle().Foreground(t.TextAccent()).Bold(true).Render(tc.Name),
			lipgloss.NewStyle().Foreground(t.TextMuted()).Render(toolLabel),
		)
		sb.WriteString(header + "\n")

		// Show truncated result
		result := tc.Result
		lines := strings.Split(result, "\n")
		maxLines := 8
		if len(lines) > maxLines {
			lines = lines[:maxLines]
			lines = append(lines, lipgloss.NewStyle().Foreground(t.TextMuted()).Render("  ... (truncated)"))
		}
		for _, line := range lines {
			sb.WriteString(fmt.Sprintf("  %s %s\n",
				lipgloss.NewStyle().Foreground(t.BorderMuted()).Render(styles.IconPipe),
				lipgloss.NewStyle().Foreground(t.TextMuted()).Render(styles.Truncate(line, m.width-10)),
			))
		}
	} else {
		// In-progress tool call
		header := fmt.Sprintf("  %s %s: %s",
			m.spinner.View(),
			lipgloss.NewStyle().Foreground(t.TextAccent()).Bold(true).Render(tc.Name),
			lipgloss.NewStyle().Foreground(t.TextMuted()).Render(toolLabel),
		)
		sb.WriteString(header + "\n")
	}

	return sb.String()
}

func formatToolLabel(name, input string) string {
	if input == "" {
		return "..."
	}

	// Return a short summary based on tool type
	summary := styles.Truncate(strings.TrimSpace(input), 60)
	if summary == "" {
		return "Building " + strings.ToLower(name) + "..."
	}
	return summary
}

func formatToolInput(name string, input map[string]any) string {
	switch name {
	case "Bash":
		if cmd, ok := input["command"].(string); ok {
			return cmd
		}
	case "Read":
		if path, ok := input["file_path"].(string); ok {
			return path
		}
	case "Write":
		if path, ok := input["file_path"].(string); ok {
			return path
		}
	case "Edit":
		if path, ok := input["file_path"].(string); ok {
			return path
		}
	case "Grep":
		if pattern, ok := input["pattern"].(string); ok {
			return pattern
		}
	case "Glob":
		if pattern, ok := input["pattern"].(string); ok {
			return pattern
		}
	case "WebSearch":
		if query, ok := input["query"].(string); ok {
			return query
		}
	case "WebFetch":
		if url, ok := input["url"].(string); ok {
			return url
		}
	}
	return ""
}

// Update handles messages.
func (m Messages) Update(msg tea.Msg) (Messages, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
		if m.streaming {
			m.rerender()
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the messages area.
func (m Messages) View() string {
	t := theme.Current()

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.Border()).
		Width(m.width - 2).
		Height(m.height - 2)

	return style.Render(m.viewport.View())
}

// Init initializes the messages component.
func (m Messages) Init() tea.Cmd {
	return m.spinner.Tick
}

// IsStreaming returns whether content is currently streaming.
func (m Messages) IsStreaming() bool {
	return m.streaming
}

// Clear resets all messages.
func (m *Messages) Clear() {
	m.messages = nil
	m.currentText.Reset()
	m.toolCalls = nil
	m.streaming = false
	m.lastRender = ""
	m.viewport.SetContent("")
}
