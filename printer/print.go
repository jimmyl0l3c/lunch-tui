package printer

import (
	"fmt"
	"net"
	"time"

	"github.com/jimmyl0l3c/lunch-tui/menu"
	"github.com/jimmyl0l3c/lunch-tui/scraping/config"
	"github.com/jimmyl0l3c/lunch-tui/styles"
	"github.com/jimmyl0l3c/lunch-tui/version"
)

func getIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		fmt.Println(styles.Error("Error when determining local IP:"), err)
	}

	defer conn.Close()

	return conn.LocalAddr().String()
}

func PrintMenu(scraperCfg config.ScrapersConfig, printerCfg *PrinterConfig, date time.Time) {
	restaurants := make([]menu.PrintableColumn, 0, len(scraperCfg.Scrapers))

	for _, sc := range scraperCfg.Scrapers {
		s := sc.NewScraper()

		meals, err := s.Scrape(date)
		if err != nil {
			restaurants = append(restaurants, &menu.RestaurantError{
				RestaurantName: s.Title(),
				Msg:            err.Error(),
			})
			continue
		}

		restaurants = append(restaurants, menu.RestaurantData{
			Name:  s.Title(),
			Meals: meals,
		})
	}
	var ip string
	if printerCfg.ShowIP {
		ip = getIp()
	}

	menuDate := date.Format(printerCfg.DateFormat)

	RenderWindow(version.Version, ip, menuDate, restaurants)
}
