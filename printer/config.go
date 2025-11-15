package printer

type PrinterConfig struct {
	ShowIP     bool   `help:"If set, shows machine's IP at the bottom of the screen."`
	DateFormat string `default:"2.1." help:"Display format of the date."`
	Scrapers   []byte `type:"filecontent" help:"Path to scrapers JSON config."`
}
