package cli

import (
	"time"

	"github.com/jimmyl0l3c/lunch-tui/printer"
	"github.com/jimmyl0l3c/lunch-tui/scraping/config"
)

type PrintCmd struct {
	Date time.Time `arg:"" optional:"" format:"2006-01-02" placeholder:"2006-01-21" help:"Date to use instead of current."`
}

func (cmd *PrintCmd) Run(cfg *printer.PrinterConfig) error {
	scraperCfg, err := config.GetConfig(cfg.Scrapers)
	if err != nil {
		return err
	}

	if cmd.Date.IsZero() {
		currentTime := time.Now().Local()
		printer.PrintMenu(scraperCfg, cfg, currentTime)
		return nil
	}

	printer.PrintMenu(scraperCfg, cfg, cmd.Date)
	return nil
}
