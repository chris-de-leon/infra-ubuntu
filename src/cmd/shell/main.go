package shell

import (
	"ubctl/src/lib/nix"

	"github.com/urfave/cli/v2"
)

var Commands = &cli.Command{
	Name:  "shell",
	Usage: "Starts a pre-configured Nix dev shell",
	Action: func(ctx *cli.Context) error {
		if err := nix.NixDevelopDevShell(); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
