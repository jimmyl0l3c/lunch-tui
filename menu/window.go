package menu

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/lunch-tui/styles"
	"golang.org/x/term"
)

func RenderWindow(scraperVersion string, ip string, menuDate string, restaurants []PrintableColumn) {
	physicalWidth, physicalHeight, _ := term.GetSize(int(os.Stdout.Fd()))

	style := styles.WindowStyle.Height(physicalHeight)

	title := RenderTitle(scraperVersion, menuDate, physicalWidth)

	restaurantRow := RestaurantRow(restaurants, physicalWidth)

	mainContent := lipgloss.JoinVertical(lipgloss.Center, fmt.Sprint(title, "\n\n"), restaurantRow)

	ipLine := styles.DetailStyle.Faint(true).Render(fmt.Sprintf("IP: %s", ip))

	fmt.Println(style.Render(
		lipgloss.JoinVertical(lipgloss.Right, mainContent, ipLine),
	))
}
