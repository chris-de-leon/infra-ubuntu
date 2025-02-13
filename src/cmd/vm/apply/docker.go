package apply

import (
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var docker = &cli.Command{
	Name:  "docker",
	Usage: "Installs docker, creates a docker group, and adds the currently logged in user to the docker group",
	Action: func(ctx *cli.Context) error {
		if err := ansible.DockerInit(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
