package scraper

import (
	"time"

	"github.com/jimmyl0l3c/lunch-tui/menu"
)

type RestaurantScraper interface {
	Title() string
	Scrape(date time.Time) ([]menu.Meal, error)
}
