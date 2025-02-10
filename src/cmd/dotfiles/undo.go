package dotfiles

import (
	"os"
	"path/filepath"
	"ubctl/src/dirs"

	"github.com/urfave/cli/v2"
)

var undoCmd = &cli.Command{
	Name:  "undo",
	Usage: "Removes all dotfiles created by the dotfiles init command",
	Action: func(ctx *cli.Context) error {
		if err := os.RemoveAll(filepath.Join(dirs.Config, "starship")); err != nil {
			return cli.Exit(err, 1)
		}

		if err := os.RemoveAll(filepath.Join(dirs.Config, "tmux")); err != nil {
			return cli.Exit(err, 1)
		}

		if err := os.RemoveAll(filepath.Join(dirs.Config, "nvim")); err != nil {
			return cli.Exit(err, 1)
		}

		return nil
	},
}
