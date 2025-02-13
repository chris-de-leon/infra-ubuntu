package apply

import (
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var bashrc = &cli.Command{
	Name:  "bashrc",
	Usage: "Configures bashrc",
	Action: func(ctx *cli.Context) error {
		if err := ansible.BashrcInit(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
