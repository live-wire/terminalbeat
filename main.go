package main

import (
	"os"

	"github.com/live-wire/terminalbeat/cmd"

	_ "github.com/live-wire/terminalbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
