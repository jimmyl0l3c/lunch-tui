package menu

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

func PrintTitle(version string, currentDate string, physicalWidth int) {
	title := lipgloss.JoinVertical(lipgloss.Center,
		styles.Title(fmt.Sprint(
			"Toolchain lunch menu scraper",
			styles.Divider,
			currentDate,
			styles.Divider,
			styles.Url(version),
		)),
		styles.Subtitle(fmt.Sprint(
			"Source available",
			styles.Divider,
			styles.Url("https://github.com/jimmyl0l3c/lunch-tui"),
		)),
	)

	centeredTitle := lipgloss.PlaceHorizontal(physicalWidth, lipgloss.Center, title)

	fmt.Print(centeredTitle, "\n\n")
}
