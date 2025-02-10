package vm

import "github.com/urfave/cli/v2"

var Commands = &cli.Command{
	Name:  "vm",
	Usage: "Commands for managing Ubuntu VMs",
	Subcommands: []*cli.Command{
		initCmd,
		undoCmd,
	},
}
