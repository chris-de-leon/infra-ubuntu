package clean

import (
	"os"
	"path/filepath"
	"ubctl/src/dirs"

	"github.com/urfave/cli/v2"
)

var Commands = &cli.Command{
	Name:  "clean",
	Usage: "Removes the ubctl config and ubctl cache directories",
	Action: func(ctx *cli.Context) error {
		if err := os.RemoveAll(filepath.Dir(dirs.AppConfig)); err != nil {
			return cli.Exit(err, 1)
		}
		if err := os.RemoveAll(filepath.Dir(dirs.AppCache)); err != nil {
			return cli.Exit(err, 1)
		}
		return nil
	},
}
