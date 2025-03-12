package main

import (
	"fmt"
	"net"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/jimmyl0l3c/lunch-tui/menu"
	"github.com/jimmyl0l3c/lunch-tui/scraper"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

const scraperVersion = "v1.1.0"

func getIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		fmt.Println(styles.Error("Error when determining local IP:"), err)
	}

	defer conn.Close()

	return conn.LocalAddr().String()
}

func printMenu(currentTime time.Time) {
	currentDate := currentTime.Format("2.1.")

	restaurants := []menu.Restaurant{
		scraper.ScrapeRozmaryny(currentDate),
		scraper.ScrapeMd(currentDate),
		scraper.ScrapePaulus(currentDate),
	}

	menu.RenderWindow(scraperVersion, getIp(), currentDate, restaurants)
}

func main() {
	for {
		currentTime := time.Now().Local()
		weekday := currentTime.Weekday()

		if weekday == time.Sunday || weekday == time.Saturday {
			fmt.Println(styles.Error("Cannot display menu during weekend"))
			return
		}

		noonTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 12, 0, 0, 0, currentTime.Location())

		if hour := currentTime.Hour(); hour >= 12 && weekday != time.Friday {
			printMenu(currentTime.Add(12 * time.Hour))

			time.Sleep(time.Until(noonTime.Add(24 * time.Hour)))
		} else if hour >= 12 && weekday == time.Friday {
			fmt.Println(styles.Url(
				lipgloss.JoinVertical(lipgloss.Center,
					"Tomorow is weekend, nothing to show.",
					"... and what are you doing here anyway?",
				),
			))
			return
		} else if hour < 12 {
			printMenu(currentTime)

			time.Sleep(time.Until(noonTime))
		}
	}
}
