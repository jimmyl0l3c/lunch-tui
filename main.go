package main

import (
	"fmt"
	"time"

	"github.com/jimmyl0l3c/lunch-tui/menu"
	"github.com/jimmyl0l3c/lunch-tui/scraper"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

func main() {
	scraperVersion := "v1.0.1"

	currentTime := time.Now().Local()
	currentDate := currentTime.Format("2.1.")

	if weekday := currentTime.Weekday(); weekday == time.Sunday || weekday == time.Saturday {
		fmt.Println(styles.Error("Cannot display menu during weekend"))
		return
	}

	restaurants := []menu.Restaurant{
		scraper.ScrapeRozmaryny(currentDate),
		scraper.ScrapeMd(currentDate),
		scraper.ScrapePaulus(currentDate),
	}

	menu.RenderWindow(scraperVersion, currentDate, restaurants)
}
