package rollback

import (
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var docker = &cli.Command{
	Name:  "docker",
	Usage: "Uninstalls docker and removes the docker group",
	Action: func(ctx *cli.Context) error {
		if err := ansible.DockerUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
