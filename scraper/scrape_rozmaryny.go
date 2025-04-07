package scraper

import (
	"github.com/gocolly/colly"
	"github.com/jimmyl0l3c/lunch-tui/menu"
)

const (
	rozmarynyTitle = "Rozmarýny"
	rozmarynyUrl   = "https://rozmaryny.cz/menu"
)

func ScrapeRozmaryny(dateFilter string) menu.PrintableColumn {
	c := colly.NewCollector()
	c.CheckHead = true

	var restaurant menu.PrintableColumn
	var parsingErr *menu.RestaurantError

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
				parsingErr = &menu.RestaurantError{
					RestaurantName: rozmarynyTitle,
					Msg:            "Out of bounds: " + h.ChildText(".dailyMenuDesc"),
				}
				return
			}

			meals[i].Detail = h.ChildText(".dailyMenuDesc")
		})

		if len(meals) == 0 {
			return
		}

		restaurant = menu.RestaurantData{Name: rozmarynyTitle, Meals: meals}
	})

	requestErr := retryScrape(c, rozmarynyUrl)

	if requestErr != nil {
		return &menu.RestaurantError{RestaurantName: rozmarynyTitle, Msg: requestErr.Error()}
	} else if parsingErr != nil {
		return parsingErr
	} else if restaurant == nil {
		return &menu.RestaurantError{RestaurantName: rozmarynyTitle, Msg: "Could not read menu"}
	}

	return restaurant
}
