package cli

import "github.com/jimmyl0l3c/lunch-tui/printer"

var CLI struct {
	PrinterCfg printer.PrinterConfig `embed:""`

	Print     PrintCmd     `cmd:"" help:"Print current menu and exit."`
	Dashboard DashboardCmd `cmd:"" help:"Dashboard mode. Prints current menu unless current time is over threshold, then prints the next day's menu."`
}
