package vm

import (
	"os"
	"ubctl/src/lib/ansible"

	"github.com/urfave/cli/v2"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Run a custom ansible playbook",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{Name: "arg", Aliases: []string{"a"}, Required: false},
		&cli.StringFlag{Name: "filepath", Aliases: []string{"f"}, Required: true},
	},
	Action: func(ctx *cli.Context) error {
		fpath := ctx.String("filepath")
		args := ctx.StringSlice("arg")

		data, err := os.ReadFile(fpath)
		if err != nil {
			return cli.Exit(err, 1)
		}

		if err := ansible.Run(ctx.Context, data, args); err != nil {
			return cli.Exit(err, 1)
		}

		return nil
	},
}
