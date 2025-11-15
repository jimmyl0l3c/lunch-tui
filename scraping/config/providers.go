package config

import (
	"github.com/jimmyl0l3c/lunch-tui/scraping/providers"
	"github.com/jimmyl0l3c/lunch-tui/scraping/scraper"
	t "github.com/jimmyl0l3c/lunch-tui/scraping/types"
)

var _ ScraperCfg = (*OlomoucCfg)(nil)

type OlomoucCfg struct {
	BaseScraperCfg

	MenuID string `json:"menuId"`
	Title  string `json:"title"`
}

func NewOlomoucCfg(title string, menuID string) *OlomoucCfg {
	return &OlomoucCfg{
		BaseScraperCfg: BaseScraperCfg{Type: t.Olomouc},
		Title:          title,
		MenuID:         menuID,
	}
}

func (cfg OlomoucCfg) NewScraper() scraper.RestaurantScraper {
	return providers.NewOlomoucScraper(cfg.Title, cfg.MenuID)
}

var _ ScraperCfg = (*RozmarynyCfg)(nil)

type RozmarynyCfg struct {
	BaseScraperCfg
}

func NewRozamrynyCfg() *RozmarynyCfg {
	return &RozmarynyCfg{BaseScraperCfg: BaseScraperCfg{Type: t.Rozmaryny}}
}

func (cfg RozmarynyCfg) NewScraper() scraper.RestaurantScraper {
	return providers.NewRozmarynyScraper()
}
