package vm

import (
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Initializes a VM",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "gh-username", Required: true},
		&cli.StringFlag{Name: "gh-token", Required: true},
		&cli.StringFlag{Name: "gh-email", Required: true},
		&cli.StringFlag{Name: "gh-name", Required: true},
	},
	Action: func(ctx *cli.Context) error {
		if err := ansible.AptInit(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.DockerInit(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.DotfilesInit(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.BashrcInit(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.GitInit(ctx.Context, ctx.String("gh-username"), ctx.String("gh-token"), ctx.String("gh-name"), ctx.String("gh-email")); err != nil {
			return cli.Exit(err, 1)
		}

		return nil
	},
}
