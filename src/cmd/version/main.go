package version

import (
	"fmt"
	"ubctl/src/lib/version"

	"github.com/urfave/cli/v2"
)

var Commands = &cli.Command{
	Name:  "version",
	Usage: "Prints the current CLI version",
	Action: func(ctx *cli.Context) error {
		fmt.Fprintln(ctx.App.Writer, version.WithoutPrefix())
		return nil
	},
}
