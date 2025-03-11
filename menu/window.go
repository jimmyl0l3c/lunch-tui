package menu

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/lunch-tui/styles"
	"golang.org/x/term"
)

func RenderWindow(scraperVersion string, currentDate string, restaurants []Restaurant) {
	physicalWidth, physicalHeight, _ := term.GetSize(int(os.Stdout.Fd()))

	style := styles.WindowStyle.Height(physicalHeight)

	title := RenderTitle(scraperVersion, currentDate, physicalWidth)

	restaurantRow := RestaurantRow(restaurants, physicalWidth)

	fmt.Println(style.Render(
		lipgloss.JoinVertical(lipgloss.Bottom, fmt.Sprint(title, "\n\n"), restaurantRow),
	))
}
