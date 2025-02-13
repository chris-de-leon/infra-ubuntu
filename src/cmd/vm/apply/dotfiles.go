package apply

import (
	"ubctl/src/lib/ansible"

	"github.com/urfave/cli/v2"
)

var dotfiles = &cli.Command{
	Name:  "dotfiles",
	Usage: "Creates dotfiles",
	Action: func(ctx *cli.Context) error {
		if err := ansible.DotfilesInit(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
