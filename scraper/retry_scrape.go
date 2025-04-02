package scraper

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

const (
	retryCount = 3
	retryDelay = 5
)

func retryScrape(collector *colly.Collector, url string) error {
	var err error

	for i := range retryCount {
		err = collector.Visit(url)

		if err == nil {
			return nil
		}

		if i < retryCount-1 {
			fmt.Println(styles.Error("Something went wrong, retrying:"), err)
			time.Sleep(retryDelay * time.Second)
		} else {
			fmt.Println(styles.Error("Something went wrong:"), err)
		}
	}

	return err
}
