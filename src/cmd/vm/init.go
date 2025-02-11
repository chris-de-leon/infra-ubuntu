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

//go:embed assets/playbooks/apt/update-upgrade.yml
var playbookAptUpdateUpgrade []byte

//go:embed assets/playbooks/docker/init.yml
var playbookDockerInit []byte

//go:embed assets/playbooks/docker/init.group.yml
var playbookDockerInitGroup []byte

//go:embed assets/playbooks/git/init.config.yml
var playbookGitInitConfig []byte

//go:embed assets/playbooks/git/init.credentials.yml
var playbookGitInitCredentials []byte

//go:embed assets/playbooks/bashrc/init.yml
var playbookBashrcInit []byte

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Initializes a VM by installing Docker, configuring Git, and updating bashrc",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "gh-username", Required: true},
		&cli.StringFlag{Name: "gh-token", Required: true},
		&cli.StringFlag{Name: "gh-email", Required: true},
		&cli.StringFlag{Name: "gh-name", Required: true},
	},
	Action: func(ctx *cli.Context) error {
		vmCacheDir := filepath.Join(dirs.AppCache, "vm", "assets")

		flakePath, err := dirs.WriteFile(filepath.Join(vmCacheDir, "flake.nix"), flake)
		if err != nil {
			return cli.Exit(err, 1)
		}

		playbooks := [][]byte{
			playbookAptUpdateUpgrade,
			playbookDockerInit,
			playbookDockerInitGroup,
			playbookGitInitConfig,
			playbookGitInitCredentials,
			playbookBashrcInit,
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

			cmd.Env = append(
				os.Environ(),
				[]string{
					fmt.Sprintf("GH_UNAME=%s", ctx.String("gh-username")),
					fmt.Sprintf("GH_TOKEN=%s", ctx.String("gh-token")),
					fmt.Sprintf("GH_EMAIL=%s", ctx.String("gh-email")),
					fmt.Sprintf("GH_NAME=%s", ctx.String("gh-name")),
				}...,
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
