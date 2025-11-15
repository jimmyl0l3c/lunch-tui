package main

import (
	"github.com/alecthomas/kong"

	"github.com/jimmyl0l3c/lunch-tui/cli"
)

func main() {
	ctx := kong.Parse(&cli.CLI)
	err := ctx.Run(&cli.CLI.PrinterCfg)
	ctx.FatalIfErrorf(err)
}
