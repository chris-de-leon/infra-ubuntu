package vm

import (
	"crypto/md5"
	"fmt"
	"os"
	"ubctl/src/ansible"

	"github.com/urfave/cli/v2"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Run a custom ansible playbook",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{Name: "arg", Aliases: []string{"-a"}, Required: true},
		&cli.StringFlag{Name: "filepath", Aliases: []string{"-f"}, Required: true},
	},
	Action: func(ctx *cli.Context) error {
		fpath := ctx.String("filepath")
		args := ctx.StringSlice("arg")

		data, err := os.ReadFile(fpath)
		if err != nil {
			return cli.Exit(err, 1)
		}

		name := fmt.Sprintf("%x", md5.Sum(data))
		if err := ansible.Run(ctx.Context, name, data, args); err != nil {
			return cli.Exit(err, 1)
		}

		return nil
	},
}
