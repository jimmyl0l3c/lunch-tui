package providers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gocolly/colly"

	"github.com/jimmyl0l3c/lunch-tui/menu"
	"github.com/jimmyl0l3c/lunch-tui/scraping/scraper"
)

const (
	rozmarynyTitle = "Rozmarýny"
	rozmarynyURL   = "https://rozmaryny.cz/menu"
)

var _ scraper.RestaurantScraper = (*RozmarynyScraper)(nil)

type RozmarynyScraper struct{}

func NewRozmarynyScraper() *RozmarynyScraper {
	return &RozmarynyScraper{}
}

func (s *RozmarynyScraper) Title() string {
	return rozmarynyTitle
}

func (s *RozmarynyScraper) Scrape(date time.Time) ([]menu.Meal, error) {
	dateFilter := date.Format("2.1.")

	c := colly.NewCollector()
	c.CheckHead = true

	var meals []menu.Meal
	var parsingErr error

	c.OnHTML("div[class=dailyMenuMainGroup]", func(e *colly.HTMLElement) {
		if e.ChildText(".dailyMenuDate") != dateFilter {
			return
		}

		meals = make([]menu.Meal, 0, 5)

		e.ForEach(".dailyMenuRow", func(i int, h *colly.HTMLElement) {
			if i == 0 {
				// Skipping first line with date
				return
			}

			meals = append(meals, menu.Meal{
				Name:   h.ChildText(".dailyMenu"),
				Detail: "",
				Price:  h.ChildText(".dailyMenuPrice"),
			})
		})

		e.ForEach(".dailyMenuDescRow", func(i int, h *colly.HTMLElement) {
			if i >= len(meals) {
				parsingErr = fmt.Errorf("out of bounds: %s", h.ChildText(".dailyMenuDesc"))
				return
			}

			meals[i].Detail = h.ChildText(".dailyMenuDesc")
		})
	})

	requestErr := scraper.RetryScrape(c, rozmarynyURL)

	if len(meals) == 0 && parsingErr == nil {
		parsingErr = fmt.Errorf("menu could not be read")
	}

	return meals, errors.Join(parsingErr, requestErr)
}
