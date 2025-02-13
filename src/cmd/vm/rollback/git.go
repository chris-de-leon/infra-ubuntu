package rollback

import (
	"ubctl/src/lib/ansible"

	"github.com/urfave/cli/v2"
)

var git = &cli.Command{
	Name:  "git",
	Usage: "Removes .gitconfig and .git-credentials",
	Action: func(ctx *cli.Context) error {
		if err := ansible.GitUndo(ctx.Context); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
