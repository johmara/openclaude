package layout

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// PlaceOverlay centers an overlay on top of a background.
func PlaceOverlay(bg string, overlay string, termWidth, termHeight int) string {
	bgLines := strings.Split(bg, "\n")

	// Ensure background fills the terminal
	for len(bgLines) < termHeight {
		bgLines = append(bgLines, strings.Repeat(" ", termWidth))
	}

	overlayLines := strings.Split(overlay, "\n")
	overlayH := len(overlayLines)
	overlayW := lipgloss.Width(overlay)

	// Center the overlay
	startY := (termHeight - overlayH) / 2
	startX := (termWidth - overlayW) / 2

	if startY < 0 {
		startY = 0
	}
	if startX < 0 {
		startX = 0
	}

	for i, line := range overlayLines {
		row := startY + i
		if row >= len(bgLines) {
			break
		}
		bgLine := bgLines[row]
		bgRunes := []rune(bgLine)
		overlayRunes := []rune(line)

		// Pad bg line if needed
		for len(bgRunes) < startX+len(overlayRunes) {
			bgRunes = append(bgRunes, ' ')
		}

		// Overwrite the background runes with the overlay
		for j, r := range overlayRunes {
			pos := startX + j
			if pos < len(bgRunes) {
				bgRunes[pos] = r
			}
		}

		bgLines[row] = string(bgRunes)
	}

	return strings.Join(bgLines[:termHeight], "\n")
}
