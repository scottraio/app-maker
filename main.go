package main

import (
	"flag"

	plum "github.com/scottraio/plum"
)

func main() {
	var cliMode bool

	flag.BoolVar(&cliMode, "cli", false, "run in CLI mode")
	flag.Parse()

	Boot()

	plum.Cli(plum.CliConfig{
		DefaultAgent: "code",
	})
}
