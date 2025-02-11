package shell

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"ubctl/src/dirs"

	"github.com/urfave/cli/v2"
)

//go:embed assets/flake.nix
var flake []byte

var Commands = &cli.Command{
	Name:  "shell",
	Usage: "Starts a Nix dev shell with the tools that chris-de-leon commonly uses",
	Flags: []cli.Flag{&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Required: false}},
	Action: func(ctx *cli.Context) error {
		shellCacheDir := filepath.Join(dirs.AppCache, "shell", "assets")
		if _, err := dirs.WriteFile(filepath.Join(shellCacheDir, "flake.nix"), flake); err != nil {
			return cli.Exit(err, 1)
		}

		// HACK: if `nix develop` is used in a Git repo and we pass it the path to a directory that
		// contains a flake.nix file which lives outside the repo, then this will result in an error.
		// To get around this, our custom nix dev shell build env is first saved to a profile which
		// lives in the user's config directory. Then, we pass the path to the custom profile to the
		// `nix develop` command. This approach allows us to avoid errors with Nix not "seeing" any
		// untracked Git files. Also, the second command is needed because the first one will start
		// a Nix shell inside the config directory rather than the user's CWD.
		profile := filepath.Join(shellCacheDir, "devsh")
		cmds := []*exec.Cmd{
			exec.Command(
				"nix",
				"develop",
				"--show-trace",
				"--profile", profile,
				"--command", "bash", "-c", "",
			),
			exec.Command(
				"nix",
				"develop",
				"--show-trace",
				profile,
			),
		}

		for i, cmd := range cmds {
			if ctx.Bool("verbose") {
				fmt.Fprintln(ctx.App.Writer, cmd.String())
			}
			if i == 0 {
				cmd.Dir = shellCacheDir
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			if err := cmd.Run(); err != nil {
				return cli.Exit(err, 1)
			}
		}

		return nil
	},
}
