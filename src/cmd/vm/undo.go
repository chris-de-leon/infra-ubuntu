package vm

import (
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var undoCmd = &cli.Command{
	Name:  "undo",
	Usage: "Reverts all the changes made by the vm init command",
	Action: func(ctx *cli.Context) error {
		if err := ansible.GitUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.BashrcUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.DotfilesUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.DockerUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.AptUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		return nil
	},
}
