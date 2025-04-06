package scraper

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/jimmyl0l3c/lunch-tui/menu"
)

const soupPrice = "V ceně jídla"

func ScrapeOlomouc(url string, restaurantName string, dateFilter string) menu.PrintableColumn {
	c := colly.NewCollector()
	c.CheckHead = true

	var restaurant menu.PrintableColumn
	var parsingErr *menu.RestaurantError

	c.OnHTML("section[class=detail-restaurace]", func(e *colly.HTMLElement) {
		meals := make([]menu.Meal, 0, 5)

		menuIndex := -1

		dateRegex := regexp.MustCompile(`(?P<Day>\d+)\.\s+(?P<Month>\d+)\.`)

		e.ForEach("h3", func(i int, h *colly.HTMLElement) {
			matchedDate := dateRegex.FindStringSubmatch(h.Text)

			if len(matchedDate) == 0 {
				return
			}

			if dateFilter == fmt.Sprintf("%s.%s.", matchedDate[1], matchedDate[2]) {
				menuIndex = i
			}
		})

		if menuIndex < 0 {
			parsingErr = &menu.RestaurantError{RestaurantName: restaurantName, Msg: "Could not match date"}
			return
		}

		e.ForEach("table", func(i int, h *colly.HTMLElement) {
			if i != menuIndex {
				return
			}

			h.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
				mealName := tr.ChildText("td:nth-child(2)")
				mealPrice := tr.ChildText("td:nth-child(3)")

				if mealIndex := tr.ChildText("td:first-child"); len(mealIndex) > 0 {
					mealName = fmt.Sprintf("%s %s", mealIndex, mealName)
				} else {
					mealPrice = soupPrice
				}

				meals = append(meals, menu.Meal{
					Name:  mealName,
					Price: mealPrice,
				})
			})
		})

		if len(meals) == 0 {
			return
		}

		restaurant = menu.RestaurantData{Name: restaurantName, Meals: meals}
	})

	requestErr := retryScrape(c, url)

	if requestErr != nil {
		return &menu.RestaurantError{RestaurantName: restaurantName, Msg: requestErr.Error()}
	} else if parsingErr != nil {
		return parsingErr
	} else if restaurant == nil {
		return &menu.RestaurantError{RestaurantName: restaurantName, Msg: "Could not read menu"}
	}

	return restaurant
}

func ScrapeMd(dateFilter string) menu.PrintableColumn {
	return ScrapeOlomouc(
		"https://www.olomouc.cz/poledni-menu/MD-Original-1869-id2208",
		"M.D. Original 1869",
		dateFilter,
	)
}

func ScrapePaulus(dateFilter string) menu.PrintableColumn {
	return ScrapeOlomouc(
		"https://www.olomouc.cz/poledni-menu/Bistro-Paulus-6806",
		"Bistro Paulus",
		dateFilter,
	)
}
