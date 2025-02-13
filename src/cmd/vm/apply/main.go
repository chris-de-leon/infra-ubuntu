package apply

import "github.com/urfave/cli/v2"

var Commands = &cli.Command{
	Name:  "apply",
	Usage: "Apply a migration to a VM",
	Subcommands: []*cli.Command{
		dotfiles,
		docker,
		bashrc,
		apt,
		git,
	},
}
