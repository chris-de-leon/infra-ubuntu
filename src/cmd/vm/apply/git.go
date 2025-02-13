package apply

import (
	"ubctl/src/lib/ansible"

	"github.com/urfave/cli/v2"
)

var git = &cli.Command{
	Name:  "git",
	Usage: "Confgiures .gitconfig and .git-credentials",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "gh-username", Required: true},
		&cli.StringFlag{Name: "gh-token", Required: true},
		&cli.StringFlag{Name: "gh-email", Required: true},
		&cli.StringFlag{Name: "gh-name", Required: true},
	},
	Action: func(ctx *cli.Context) error {
		if err := ansible.GitInit(ctx.Context,
			ctx.String("gh-username"),
			ctx.String("gh-token"),
			ctx.String("gh-name"),
			ctx.String("gh-email"),
		); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
