package claude

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

// StreamParser reads NDJSON lines from a reader and emits typed Events.
type StreamParser struct {
	reader io.Reader
}

// NewStreamParser creates a new stream parser.
func NewStreamParser(r io.Reader) *StreamParser {
	return &StreamParser{reader: r}
}

// Parse reads all NDJSON lines and sends typed events to the channel.
// It closes the channel when the reader is exhausted.
func (p *StreamParser) Parse(ch chan<- Event) {
	defer close(ch)

	scanner := bufio.NewScanner(p.reader)
	scanner.Buffer(make([]byte, 0, 1024*1024), 10*1024*1024)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var raw RawEvent
		if err := json.Unmarshal(line, &raw); err != nil {
			ch <- Event{Type: EventError, ResultError: fmt.Sprintf("parse error: %v", err)}
			continue
		}

		events := parseRawEvent(&raw)
		for _, evt := range events {
			ch <- evt
		}
	}

	if err := scanner.Err(); err != nil {
		ch <- Event{Type: EventError, ResultError: fmt.Sprintf("scanner error: %v", err)}
	}
}

func parseRawEvent(raw *RawEvent) []Event {
	switch raw.Type {
	case "system":
		if raw.Subtype == "init" {
			return []Event{{
				Type:      EventInit,
				SessionID: raw.SessionID,
				Model:     raw.Model,
				Tools:     raw.Tools,
			}}
		}

	case "stream_event":
		return parseStreamEvent(raw.StreamEvent)

	case "assistant":
		if raw.Message != nil {
			return []Event{{
				Type:    EventAssistantMessage,
				Message: raw.Message,
			}}
		}

	case "user":
		if raw.Message != nil {
			return parseUserMessage(raw.Message)
		}

	case "result":
		evt := Event{
			Type:         EventResult,
			TotalCostUSD: raw.TotalCostUSD,
			InputTokens:  raw.InputTokens,
			OutputTokens: raw.OutputTokens,
			Duration:     raw.Duration,
			Result:       raw.Result,
		}
		if raw.IsError {
			evt.ResultError = raw.ErrorMessage
		}
		return []Event{evt}
	}

	return nil
}

func parseStreamEvent(se *StreamEvent) []Event {
	if se == nil {
		return nil
	}

	switch se.Type {
	case "content_block_start":
		if se.ContentBlock != nil && se.ContentBlock.Type == "tool_use" {
			return []Event{{
				Type:     EventToolStart,
				ToolName: se.ContentBlock.Name,
				ToolID:   se.ContentBlock.ID,
				BlockIdx: se.Index,
			}}
		}

	case "content_block_delta":
		if se.Delta != nil {
			switch se.Delta.Type {
			case "text_delta":
				return []Event{{
					Type: EventTextDelta,
					Text: se.Delta.Text,
				}}
			case "input_json_delta":
				return []Event{{
					Type:        EventToolInputDelta,
					PartialJSON: se.Delta.PartialJSON,
					BlockIdx:    se.Index,
				}}
			}
		}

	case "content_block_stop":
		return []Event{{
			Type:     EventToolEnd,
			BlockIdx: se.Index,
		}}
	}

	return nil
}

func parseUserMessage(msg *Message) []Event {
	var events []Event
	for _, part := range msg.Content {
		if part.Type == "tool_result" {
			content := ""
			switch v := part.Content.(type) {
			case string:
				content = v
			case []any:
				for _, item := range v {
					if m, ok := item.(map[string]any); ok {
						if t, ok := m["text"].(string); ok {
							content += t
						}
					}
				}
			}
			events = append(events, Event{
				Type:      EventToolResult,
				ToolUseID: part.ToolUseID,
				Content:   content,
				IsError:   part.IsError,
			})
		}
	}
	return events
}
