package apply

import (
	"ubctl/src/ansible"

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
		uname := ctx.String("gh-username")
		token := ctx.String("gh-token")
		email := ctx.String("gh-email")
		name := ctx.String("gh-name")

		if err := ansible.GitInit(ctx.Context, uname, token, name, email); err != nil {
			return cli.Exit(err, 1)
		} else {
			return nil
		}
	},
}
