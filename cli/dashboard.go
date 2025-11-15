package cli

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"

	"github.com/jimmyl0l3c/lunch-tui/printer"
	"github.com/jimmyl0l3c/lunch-tui/scraping/config"
	"github.com/jimmyl0l3c/lunch-tui/styles"
)

type DashboardCmd struct {
	NextDayThreshold uint8 `default:"12" help:"Hour after which next day's menu should be shown instead of today's."`
}

func (cmd *DashboardCmd) Validate() error {
	if cmd.NextDayThreshold >= 24 {
		return fmt.Errorf("NextDayThreshold has to be on the interval 0-23, got %d instead", cmd.NextDayThreshold)
	}

	return nil
}

func (cmd *DashboardCmd) Run(cfg *printer.PrinterConfig) error {
	scraperCfg, err := config.GetConfig(cfg.Scrapers)
	if err != nil {
		return err
	}

	currentTime := time.Now().Local()
	weekday := currentTime.Weekday()

	if weekday == time.Sunday || weekday == time.Saturday {
		fmt.Println(styles.Error("Cannot display menu during weekend"))
		return nil
	}

	threshold := int(cmd.NextDayThreshold)

	if hour := currentTime.Hour(); hour >= threshold && weekday != time.Friday {
		currentTime = currentTime.Add(24 * time.Hour)
	} else if hour >= threshold && weekday == time.Friday {
		fmt.Println(styles.Url(
			lipgloss.JoinVertical(lipgloss.Center,
				"Tomorow is weekend, nothing to show.",
				"... and what are you doing here anyway?",
			),
		))
		return nil
	}

	printer.PrintMenu(scraperCfg, cfg, currentTime)
	return nil
}
