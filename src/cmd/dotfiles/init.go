package dotfiles

import (
	"embed"
	"path/filepath"
	"ubctl/src/dirs"

	"github.com/urfave/cli/v2"
)

//go:embed assets/starship/starship.toml
var starship []byte

//go:embed assets/tmux/tmux.conf
var tmux []byte

//go:embed assets/nvim/*
var nvim embed.FS

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Creates all dotfiles in the user's config directory - overwrites any existing files",
	Action: func(ctx *cli.Context) error {
		if err := dirs.WriteDir(filepath.Join("assets", "nvim"), filepath.Join(dirs.Config, "nvim"), nvim); err != nil {
			return cli.Exit(err, 1)
		}

		if _, err := dirs.WriteFile(filepath.Join(dirs.Config, "starship", "starship.toml"), starship); err != nil {
			return cli.Exit(err, 1)
		}

		if _, err := dirs.WriteFile(filepath.Join(dirs.Config, "tmux", "tmux.conf"), tmux); err != nil {
			return cli.Exit(err, 1)
		}

		return nil
	},
}
