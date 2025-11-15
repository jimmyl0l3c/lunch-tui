package config

import (
	"encoding/json"
	"fmt"

	"github.com/jimmyl0l3c/lunch-tui/scraping/scraper"
	t "github.com/jimmyl0l3c/lunch-tui/scraping/types"
)

type BaseScraperCfg struct {
	Type t.ScraperType `json:"type"`
}

type ScraperCfg interface {
	NewScraper() scraper.RestaurantScraper
}

type ScrapersConfig struct {
	Scrapers []ScraperCfg
}

func (cfg *ScrapersConfig) UnmarshalJSON(data []byte) error {
	var raws []json.RawMessage
	if err := json.Unmarshal(data, &raws); err != nil {
		return err
	}

	cfg.Scrapers = make([]ScraperCfg, 0, len(raws))
	for _, raw := range raws {
		var header BaseScraperCfg
		if err := json.Unmarshal(raw, &header); err != nil {
			return err
		}

		switch header.Type {
		case t.Rozmaryny:
			var c RozmarynyCfg
			if err := json.Unmarshal(raw, &c); err != nil {
				return err
			}
			cfg.Scrapers = append(cfg.Scrapers, c)
		case t.Olomouc:
			var c OlomoucCfg
			if err := json.Unmarshal(raw, &c); err != nil {
				return err
			}
			cfg.Scrapers = append(cfg.Scrapers, c)
		default:
			return fmt.Errorf("unknown type: %s", header.Type)
		}
	}
	return nil
}
