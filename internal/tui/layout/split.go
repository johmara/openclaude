package layout

import "github.com/charmbracelet/lipgloss"

// SplitDirection defines how a split divides space.
type SplitDirection int

const (
	Horizontal SplitDirection = iota // Left/Right
	Vertical                         // Top/Bottom
)

// RenderSplit renders two panels side by side or stacked.
func RenderSplit(dir SplitDirection, left, right string, leftSize, rightSize int) string {
	switch dir {
	case Horizontal:
		l := lipgloss.NewStyle().Width(leftSize).Render(left)
		r := lipgloss.NewStyle().Width(rightSize).Render(right)
		return lipgloss.JoinHorizontal(lipgloss.Top, l, r)
	case Vertical:
		t := lipgloss.NewStyle().Width(leftSize + rightSize).Render(left)
		b := lipgloss.NewStyle().Width(leftSize + rightSize).Render(right)
		return lipgloss.JoinVertical(lipgloss.Left, t, b)
	}
	return left
}

// RenderThreePane renders a main area with bottom and optional right sidebar.
func RenderThreePane(main, sidebar, bottom string, mainW, sidebarW, mainH, bottomH int) string {
	var topSection string
	if sidebarW > 0 && sidebar != "" {
		mainPanel := lipgloss.NewStyle().
			Width(mainW).
			Height(mainH).
			Render(main)
		sidePanel := lipgloss.NewStyle().
			Width(sidebarW).
			Height(mainH).
			Render(sidebar)
		topSection = lipgloss.JoinHorizontal(lipgloss.Top, mainPanel, sidePanel)
	} else {
		topSection = lipgloss.NewStyle().
			Width(mainW + sidebarW).
			Height(mainH).
			Render(main)
	}

	if bottomH > 0 && bottom != "" {
		bottomPanel := lipgloss.NewStyle().
			Width(mainW + sidebarW).
			Height(bottomH).
			Render(bottom)
		return lipgloss.JoinVertical(lipgloss.Left, topSection, bottomPanel)
	}

	return topSection
}
