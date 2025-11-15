package providers

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/gocolly/colly"

	"github.com/jimmyl0l3c/lunch-tui/menu"
	"github.com/jimmyl0l3c/lunch-tui/scraping/scraper"
)

const olomoucBaseURL = "https://www.olomouc.cz/poledni-menu"

var _ scraper.RestaurantScraper = (*OlomoucScraper)(nil)

type OlomoucScraper struct {
	title string
	url   string
}

func NewOlomoucScraper(title string, menuID string) *OlomoucScraper {
	return &OlomoucScraper{
		title: title,
		url:   fmt.Sprintf("%s/%s", olomoucBaseURL, menuID),
	}
}

func (s *OlomoucScraper) Title() string {
	return s.title
}

func (s *OlomoucScraper) Scrape(date time.Time) ([]menu.Meal, error) {
	dateFilter := date.Format("2.1.")

	c := colly.NewCollector()
	c.CheckHead = true

	var meals []menu.Meal
	var parsingErr error

	c.OnHTML("section[class=detail-restaurace]", func(e *colly.HTMLElement) {
		meals = make([]menu.Meal, 0, 5)

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
			parsingErr = fmt.Errorf("could not match date")
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
					mealPrice = ""
				}

				meals = append(meals, menu.Meal{
					Name:  mealName,
					Price: mealPrice,
				})
			})
		})
	})

	requestErr := scraper.RetryScrape(c, s.url)

	if len(meals) == 0 && parsingErr == nil {
		parsingErr = fmt.Errorf("menu could not be read")
	}

	return meals, errors.Join(parsingErr, requestErr)
}
