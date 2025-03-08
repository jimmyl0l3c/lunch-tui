package menu

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	url = lipgloss.NewStyle().Foreground(special).Render

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	titleStyle = base.MarginTop(1)

	sourceStyle = base.
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)
)

func PrintTitle(version string, currentDate string, physicalWidth int) {
	title := lipgloss.JoinVertical(lipgloss.Center,
		titleStyle.Render(fmt.Sprint(
			"Toolchain lunch menu scraper",
			divider,
			currentDate,
			divider,
			url(version),
		)),
		sourceStyle.Render(fmt.Sprint(
			"Source available",
			divider,
			url("https://github.com/jimmyl0l3c/lunch-tui"),
		)),
	)

	centeredTitle := lipgloss.PlaceHorizontal(physicalWidth, lipgloss.Center, title)

	fmt.Print(centeredTitle, "\n\n")
}
