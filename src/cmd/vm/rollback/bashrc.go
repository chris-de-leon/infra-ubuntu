package rollback

import (
	"ubctl/src/lib/ansible"

	"github.com/urfave/cli/v2"
)

var bashrc = &cli.Command{
	Name:  "bashrc",
	Usage: "Reverts the changes made to bashrc",
	Action: func(ctx *cli.Context) error {
		if err := ansible.BashrcUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
