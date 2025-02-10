package main

import (
	"log"
	"os"
	"ubctl/src/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "ubctl",
		Usage:    "A CLI tool for managing Ubuntu VMs",
		Commands: cmd.Commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
