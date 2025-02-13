package rollback

import "github.com/urfave/cli/v2"

var Commands = &cli.Command{
	Name:  "rollback",
	Usage: "Rollback a migration applied to a VM",
	Subcommands: []*cli.Command{
		dotfiles,
		docker,
		bashrc,
		apt,
		git,
	},
}
