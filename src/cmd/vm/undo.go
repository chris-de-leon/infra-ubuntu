package vm

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"ubctl/src/dirs"

	"github.com/urfave/cli/v2"
)

//go:embed assets/playbooks/apt/autoremove.yml
var playbookAptAutoremove []byte

//go:embed assets/playbooks/docker/undo.yml
var playbookDockerUndo []byte

//go:embed assets/playbooks/docker/undo.group.yml
var playbookDockerUndoGroup []byte

//go:embed assets/playbooks/git/undo.config.yml
var playbookGitUndoConfig []byte

//go:embed assets/playbooks/git/undo.credentials.yml
var playbookGitUndoCredentials []byte

//go:embed assets/playbooks/bashrc/undo.yml
var playbookBashrcUndo []byte

//go:embed assets/flake.nix
var flake []byte

var undoCmd = &cli.Command{
	Name:  "undo",
	Usage: "Reverts all the changes made by the vm init command",
	Action: func(ctx *cli.Context) error {
		vmCacheDir := filepath.Join(dirs.AppCache, "vm", "assets")

		flakePath, err := dirs.WriteFile(filepath.Join(vmCacheDir, "flake.nix"), flake)
		if err != nil {
			return cli.Exit(err, 1)
		}

		playbooks := [][]byte{
			playbookBashrcUndo,
			playbookGitUndoCredentials,
			playbookGitUndoConfig,
			playbookDockerUndoGroup,
			playbookDockerUndo,
			playbookAptAutoremove,
		}

		for i, pb := range playbooks {
			pbPath, err := dirs.WriteFile(filepath.Join(vmCacheDir, "playbooks", fmt.Sprintf("%d.yml", i)), pb)
			if err != nil {
				return cli.Exit(err, 1)
			}

			cmd := exec.Command(
				"nix",
				"develop",
				"--show-trace",
				filepath.Dir(flakePath),
				"--command",
				"ansible-playbook",
				pbPath,
			)

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
