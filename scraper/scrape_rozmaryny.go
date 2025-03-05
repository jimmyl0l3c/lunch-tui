package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"lunch/menu"
)

const rozmarynyTitle = "Rozmarýny"

func ScrapeRozmaryny(dateFilter string) menu.Restaurant {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println(errorStyle.Render("Something went wrong:"), err)
	})

	restaurant := menu.Restaurant{Name: rozmarynyTitle, Meals: make([]menu.Meal, 0)}

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

			meals = append(meals, menu.Meal{Name: h.ChildText(".dailyMenu"), Detail: "", Price: h.ChildText(".dailyMenuPrice")})
		})

		e.ForEach(".dailyMenuDescRow", func(i int, h *colly.HTMLElement) {
			if i >= len(meals) {
				fmt.Println(errorStyle.Render("Out of bounds:"), h.ChildText(".dailyMenuDesc"))
				return
			}

			meals[i].Detail = h.ChildText(".dailyMenuDesc")
		})

		restaurant = menu.Restaurant{Name: rozmarynyTitle, Meals: meals}
	})

	c.Visit("https://rozmaryny.cz/menu")

	return restaurant
}
