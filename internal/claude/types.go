package claude

// RawEvent represents a top-level NDJSON line from claude --output-format stream-json.
type RawEvent struct {
	Type    string `json:"type"`
	Subtype string `json:"subtype,omitempty"`

	// system init fields
	SessionID string   `json:"session_id,omitempty"`
	Model     string   `json:"model,omitempty"`
	Tools     []Tool   `json:"tools,omitempty"`
	CWD       string   `json:"cwd,omitempty"`
	MCP       []string `json:"mcp,omitempty"`

	// stream_event wrapper
	StreamEvent *StreamEvent `json:"event,omitempty"`

	// assistant / user message
	Message *Message `json:"message,omitempty"`

	// result fields
	TotalCostUSD   float64 `json:"total_cost_usd,omitempty"`
	InputTokens    int     `json:"input_tokens,omitempty"`
	OutputTokens   int     `json:"output_tokens,omitempty"`
	Duration       float64 `json:"duration_seconds,omitempty"`
	DurationAPI    float64 `json:"duration_api_seconds,omitempty"`
	NumTurns       int     `json:"num_turns,omitempty"`
	Result         string  `json:"result,omitempty"`
	IsError        bool    `json:"is_error,omitempty"`
	ErrorMessage   string  `json:"error_message,omitempty"`
}

// Tool represents a tool available to Claude.
type Tool struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
}

// StreamEvent is the inner event inside a stream_event wrapper.
type StreamEvent struct {
	Type  string `json:"type"`
	Index int    `json:"index,omitempty"`

	// content_block_start
	ContentBlock *ContentBlock `json:"content_block,omitempty"`

	// content_block_delta
	Delta *Delta `json:"delta,omitempty"`
}

// ContentBlock describes a content block (text or tool_use).
type ContentBlock struct {
	Type  string `json:"type"`
	Name  string `json:"name,omitempty"`
	ID    string `json:"id,omitempty"`
	Text  string `json:"text,omitempty"`
}

// Delta carries incremental updates.
type Delta struct {
	Type        string `json:"type"`
	Text        string `json:"text,omitempty"`
	PartialJSON string `json:"partial_json,omitempty"`
}

// Message is a complete assistant or user message.
type Message struct {
	Role    string          `json:"role"`
	Content []ContentPart   `json:"content"`
	Model   string          `json:"model,omitempty"`
	Usage   *Usage          `json:"usage,omitempty"`
}

// ContentPart is one element in a message's content array.
type ContentPart struct {
	Type      string `json:"type"`
	Text      string `json:"text,omitempty"`
	Name      string `json:"name,omitempty"`
	ID        string `json:"id,omitempty"`
	Input     any    `json:"input,omitempty"`
	ToolUseID string `json:"tool_use_id,omitempty"`
	Content   any    `json:"content,omitempty"`
	IsError   bool   `json:"is_error,omitempty"`
}

// Usage contains token usage information.
type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

// Parsed event types for the pub/sub system.

type EventType int

const (
	EventInit EventType = iota
	EventTextDelta
	EventToolStart
	EventToolInputDelta
	EventToolEnd
	EventAssistantMessage
	EventToolResult
	EventResult
	EventError
)

// Event is the parsed, typed event passed through the pub/sub broker.
type Event struct {
	Type EventType

	// EventInit
	SessionID string
	Model     string
	Tools     []Tool

	// EventTextDelta
	Text string

	// EventToolStart / EventToolEnd
	ToolName string
	ToolID   string
	BlockIdx int

	// EventToolInputDelta
	PartialJSON string

	// EventAssistantMessage
	Message *Message

	// EventToolResult
	ToolUseID string
	Content   string
	IsError   bool

	// EventResult
	TotalCostUSD float64
	InputTokens  int
	OutputTokens int
	Duration     float64
	Result       string
	ResultError  string
}
