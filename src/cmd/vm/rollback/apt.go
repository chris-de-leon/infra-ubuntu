package rollback

import (
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var apt = &cli.Command{
	Name:  "apt",
	Usage: "Removes APT packages",
	Action: func(ctx *cli.Context) error {
		if err := ansible.AptUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
