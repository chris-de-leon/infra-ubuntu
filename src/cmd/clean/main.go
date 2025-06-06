package clean

import (
	"ubctl/src/lib/paths"

	"github.com/urfave/cli/v2"
)

var Commands = &cli.Command{
	Name:  "clean",
	Usage: "Removes the ubctl config and ubctl cache directories",
	Action: func(ctx *cli.Context) error {
		if err := paths.RemoveAppConfigDir(); err != nil {
			return cli.Exit(err, 1)
		}
		if err := paths.RemoveAppCacheDir(); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	},
}
