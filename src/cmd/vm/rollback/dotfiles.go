package rollback

import (
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var dotfiles = &cli.Command{
	Name:  "dotfiles",
	Usage: "Removes dotfiles",
	Action: func(ctx *cli.Context) error {
		if err := ansible.DotfilesUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
