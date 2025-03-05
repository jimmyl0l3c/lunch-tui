package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"lunch/menu"
	"regexp"
)

const (
	soupPrice     = "V ceně jídla"
	unknownDetail = "n/a"
)

func ScrapeOlomouc(url string, restaurantName string, dateFilter string) menu.Restaurant {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println(errorStyle.Render("Something went wrong:"), err)
	})

	restaurant := menu.Restaurant{Name: restaurantName, Meals: make([]menu.Meal, 0)}

	c.OnHTML("section[class=detail-restaurace]", func(e *colly.HTMLElement) {
		meals := make([]menu.Meal, 0, 5)

		menuIndex := -1

		dateRegex := regexp.MustCompile(`(?P<Day>\d+)\.\s+(?P<Month>\d+)\.`)

		e.ForEach("h3", func(i int, h *colly.HTMLElement) {
			matchedDate := dateRegex.FindStringSubmatch(h.Text)

			if len(matchedDate) == 0 {
				return
			}

			foundDate := fmt.Sprintf("%s.%s.", matchedDate[1], matchedDate[2])

			if foundDate == dateFilter {
				menuIndex = i
			}
		})

		if menuIndex < 0 {
			fmt.Println(errorStyle.Render("Could not match date"))
			return
		}

		e.ForEach("table", func(i int, h *colly.HTMLElement) {
			if i != menuIndex {
				return
			}

			h.ForEach("tr", func(i int, tr *colly.HTMLElement) {
				mealIndex := tr.ChildText("td:first-child")
				mealName := tr.ChildText("td:nth-child(2)")
				mealPrice := tr.ChildText("td:nth-child(3)")

				if len(mealIndex) > 0 {
					mealName = mealIndex + " " + mealName
				} else {
					mealPrice = soupPrice
				}

				meals = append(meals, menu.Meal{Name: mealName, Detail: unknownDetail, Price: mealPrice})
			})
		})

		restaurant = menu.Restaurant{Name: restaurantName, Meals: meals}
	})

	c.Visit(url)

	return restaurant
}

func ScrapeMd(dateFilter string) menu.Restaurant {
	return ScrapeOlomouc("https://www.olomouc.cz/poledni-menu/MD-Original-1869-id2208", "M.D. Original 1869", dateFilter)
}

func ScrapePaulus(dateFilter string) menu.Restaurant {
	return ScrapeOlomouc("https://www.olomouc.cz/poledni-menu/Bistro-Paulus-6806", "Bistro Paulus", dateFilter)
}
