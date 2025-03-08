package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/lunch-tui/menu"
	"github.com/jimmyl0l3c/lunch-tui/scraper"
	"golang.org/x/term"
)

var docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)

func main() {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	scraperVersion := "v1.0.0"

	currentTime := time.Now().Local()
	currentDate := currentTime.Format("2.1.")

	menu.PrintTitle(scraperVersion, currentDate, physicalWidth)

	restaurants := []menu.Restaurant{
		scraper.ScrapeRozmaryny(currentDate),
		scraper.ScrapeMd(currentDate),
		scraper.ScrapePaulus(currentDate),
	}

	doc := strings.Builder{}

	doc.WriteString(menu.RestaurantRow(restaurants, physicalWidth))

	fmt.Println(docStyle.Render(doc.String()))

}
