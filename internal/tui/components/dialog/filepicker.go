package dialog

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johanmontorfano/openclaude/internal/tui/styles"
	"github.com/johanmontorfano/openclaude/internal/tui/theme"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

// FileSelectedMsg is emitted when a file is selected.
type FileSelectedMsg struct {
	Path string
}

// FilePicker is a fuzzy file picker dialog (Ctrl+F).
type FilePicker struct {
	files    []string
	filtered []string
	selected int
	input    textinput.Model
	width    int
	height   int
	maxVisible int
	cwd      string
}

// NewFilePicker creates a file picker that scans the working directory.
func NewFilePicker() FilePicker {
	ti := textinput.New()
	ti.Placeholder = "Search files..."
	ti.Focus()
	ti.CharLimit = 200

	cwd, _ := os.Getwd()

	fp := FilePicker{
		input:      ti,
		width:      60,
		height:     20,
		maxVisible: 12,
		cwd:        cwd,
	}
	fp.scanFiles()

	return fp
}

func (f *FilePicker) scanFiles() {
	f.files = nil

	_ = filepath.Walk(f.cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Skip hidden dirs and common non-useful directories
		name := info.Name()
		if info.IsDir() {
			if name == ".git" || name == "node_modules" || name == ".cache" || name == "vendor" || name == "__pycache__" {
				return filepath.SkipDir
			}
			return nil
		}

		rel, _ := filepath.Rel(f.cwd, path)
		if rel != "" {
			f.files = append(f.files, rel)
		}

		// Limit for performance
		if len(f.files) > 5000 {
			return filepath.SkipAll
		}

		return nil
	})

	f.filtered = f.files
}

// SetSize adjusts dialog size.
func (f *FilePicker) SetSize(w, h int) {
	f.width = w * 2 / 3
	if f.width > 70 {
		f.width = 70
	}
	f.height = h * 2 / 3
	if f.height > 25 {
		f.height = 25
	}
	f.maxVisible = f.height - 8
	if f.maxVisible < 3 {
		f.maxVisible = 3
	}
}

// Update handles input.
func (f FilePicker) Update(msg tea.Msg) (FilePicker, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return f, func() tea.Msg { return CloseMsg{} }
		case "enter":
			if len(f.filtered) > 0 && f.selected < len(f.filtered) {
				path := f.filtered[f.selected]
				return f, func() tea.Msg { return FileSelectedMsg{Path: path} }
			}
		case "up":
			if f.selected > 0 {
				f.selected--
			}
			return f, nil
		case "down":
			if f.selected < len(f.filtered)-1 {
				f.selected++
			}
			return f, nil
		}
	}

	var cmd tea.Cmd
	f.input, cmd = f.input.Update(msg)

	// Fuzzy filter
	query := f.input.Value()
	if query == "" {
		f.filtered = f.files
	} else {
		f.filtered = fuzzy.Find(query, f.files)
	}

	if f.selected >= len(f.filtered) {
		f.selected = len(f.filtered) - 1
	}
	if f.selected < 0 {
		f.selected = 0
	}

	return f, cmd
}

// View renders the file picker.
func (f FilePicker) View() string {
	t := theme.Current()

	var sb strings.Builder

	title := lipgloss.NewStyle().
		Foreground(t.Primary()).
		Bold(true).
		Render("File Picker")
	sb.WriteString(title + "\n\n")
	sb.WriteString(f.input.View() + "\n\n")

	visible := f.filtered
	start := 0
	if len(visible) > f.maxVisible {
		start = f.selected - f.maxVisible/2
		if start < 0 {
			start = 0
		}
		end := start + f.maxVisible
		if end > len(visible) {
			end = len(visible)
			start = end - f.maxVisible
			if start < 0 {
				start = 0
			}
		}
		visible = visible[start:start+min(f.maxVisible, len(visible)-start)]
	}

	for i, file := range visible {
		idx := start + i
		prefix := "  "
		style := lipgloss.NewStyle().Foreground(t.Text())

		if idx == f.selected {
			prefix = styles.IconArrowRight + " "
			style = lipgloss.NewStyle().
				Foreground(t.SelectionFg()).
				Background(t.SelectionBg()).
				Bold(true)
		}

		line := style.Render(prefix + styles.Truncate(file, f.width-8))
		sb.WriteString(line + "\n")
	}

	if len(f.filtered) == 0 {
		sb.WriteString(lipgloss.NewStyle().
			Foreground(t.TextMuted()).
			Render("  No files found") + "\n")
	}

	count := lipgloss.NewStyle().
		Foreground(t.TextMuted()).
		Render(strings.Repeat(" ", 0))
	sb.WriteString("\n" + count)
	sb.WriteString(lipgloss.NewStyle().
		Foreground(t.TextMuted()).
		Render("↑↓ navigate • enter select • esc close"))

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.DialogBorder()).
		Background(t.DialogBg()).
		Padding(1, 2).
		Width(f.width).
		Render(sb.String())
}

// Init initializes the file picker.
func (f FilePicker) Init() tea.Cmd {
	return textinput.Blink
}
