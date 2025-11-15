package printer

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/jimmyl0l3c/lunch-tui/menu"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

func RenderWindow(scraperVersion string, ip string, menuDate string, restaurants []menu.PrintableColumn) {
	physicalWidth, physicalHeight, _ := term.GetSize(int(os.Stdout.Fd()))

	style := styles.WindowStyle.Height(physicalHeight)

	title := RenderTitle(scraperVersion, menuDate, physicalWidth)

	restaurantRow := RestaurantRow(restaurants, physicalWidth)

	mainContent := lipgloss.JoinVertical(lipgloss.Center, fmt.Sprint(title, "\n\n"), restaurantRow)

	if len(ip) > 0 {
		ipLine := styles.DetailStyle.Faint(true).Render(fmt.Sprintf("IP: %s", ip))

		fmt.Println(style.Render(
			lipgloss.JoinVertical(lipgloss.Right, mainContent, ipLine),
		))

		return
	}

	fmt.Println(style.Render(mainContent))
}

func RestaurantRow(restaurants []menu.PrintableColumn, physicalWidth int) (row string) {
	maxMenuWidth := (physicalWidth - 4 - (6 * len(restaurants))) / len(restaurants)

	lastIndex := len(restaurants) - 1

	for i, restaurant := range restaurants {
		column := restaurant.Render(maxMenuWidth)

		if i == 0 && len(restaurants) > 1 {
			row = styles.List(column)
		} else if i == lastIndex {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, column)
		} else {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, styles.List(column))
		}
	}

	return
}

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
