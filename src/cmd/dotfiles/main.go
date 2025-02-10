package dotfiles

import "github.com/urfave/cli/v2"

var Commands = &cli.Command{
	Name:  "dotfiles",
	Usage: "Commands for managing dotfile configs",
	Subcommands: []*cli.Command{
		initCmd,
		undoCmd,
	},
}
