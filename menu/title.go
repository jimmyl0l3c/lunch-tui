package menu

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

func RenderTitle(version string, menuDate string, physicalWidth int) string {
	title := lipgloss.JoinVertical(lipgloss.Center,
		styles.Title(fmt.Sprint(
			"Toolchain lunch menu scraper",
			styles.Divider,
			menuDate,
			styles.Divider,
			styles.Url(version),
		)),
		styles.Subtitle(fmt.Sprint(
			"Source available",
			styles.Divider,
			styles.Url("https://github.com/jimmyl0l3c/lunch-tui"),
		)),
	)

	return lipgloss.PlaceHorizontal(physicalWidth-12, lipgloss.Center, title)
}
