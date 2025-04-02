package scraper

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/jimmyl0l3c/lunch-tui/menu"
)

const rozmarynyTitle = "Rozmarýny"
const rozmarynyUrl = "https://rozmaryny.cz/menu"

func ScrapeRozmaryny(dateFilter string) menu.Restaurant {
	c := colly.NewCollector()
	c.CheckHead = true

	restaurant := menu.Restaurant{Name: rozmarynyTitle, Meals: make([]menu.Meal, 0)}

	var parsingErr error

	c.OnHTML("div[class=dailyMenuMainGroup]", func(e *colly.HTMLElement) {
		if e.ChildText(".dailyMenuDate") != dateFilter {
			return
		}

		meals := make([]menu.Meal, 0, 5)

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
				parsingErr = errors.New("Out of bounds: " + h.ChildText(".dailyMenuDesc"))
				return
			}

			meals[i].Detail = h.ChildText(".dailyMenuDesc")
		})

		restaurant = menu.Restaurant{Name: rozmarynyTitle, Meals: meals}
	})

	requestErr := RetryScrape(c, rozmarynyUrl)

	if requestErr != nil {
		restaurant.Err = requestErr
	} else {
		restaurant.Err = parsingErr
	}

	return restaurant
}
