package apply

import (
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var apt = &cli.Command{
	Name:  "apt",
	Usage: "Installs APT packages",
	Action: func(ctx *cli.Context) error {
		if err := ansible.AptInit(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
