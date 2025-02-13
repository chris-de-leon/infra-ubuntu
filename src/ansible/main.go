package ansible

import (
	"context"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"ubctl/src/dirs"
)

//go:embed assets/playbooks/apt/init.yml
var playbookAptInit []byte

//go:embed assets/playbooks/apt/undo.yml
var playbookAptUndo []byte

//go:embed assets/playbooks/docker/init.yml
var playbookDockerInit []byte

//go:embed assets/playbooks/docker/undo.yml
var playbookDockerUndo []byte

//go:embed assets/playbooks/git/init.yml
var playbookGitInit []byte

//go:embed assets/playbooks/git/undo.yml
var playbookGitUndo []byte

//go:embed assets/playbooks/bashrc/init.yml
var playbookBashrcInit []byte

//go:embed assets/playbooks/bashrc/undo.yml
var playbookBashrcUndo []byte

//go:embed assets/playbooks/dotfiles/init.yml
var playbookDotfilesInit []byte

//go:embed assets/playbooks/dotfiles/undo.yml
var playbookDotfilesUndo []byte

//go:embed assets/dotfiles/starship/starship.toml
var starship []byte

//go:embed assets/dotfiles/tmux/tmux.conf
var tmux []byte

//go:embed assets/dotfiles/nvim/*
var nvim embed.FS

//go:embed assets/flake.nix
var flake []byte

var (
	PlaybooksDir string
	DotfilesDir  string
	AnsibleDir   string
)

func init() {
	AnsibleDir = filepath.Join(dirs.AppCache, "ansible")
	PlaybooksDir = filepath.Join(AnsibleDir, "playbooks")
	DotfilesDir = filepath.Join(AnsibleDir, "dotfiles")
}

func Run(ctx context.Context, name string, data []byte, vars []string) error {
	playbookPath, err := dirs.WriteFile(filepath.Join(PlaybooksDir, name), data)
	if err != nil {
		return err
	}

	flakePath, err := dirs.WriteFile(filepath.Join(AnsibleDir, "flake.nix"), flake)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(
		ctx,
		"nix",
		"develop",
		"--show-trace",
		filepath.Dir(flakePath),
		"--command",
		"ansible-playbook",
		playbookPath,
	)

	cmd.Env = append(
		os.Environ(),
		vars...,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func AptInit(ctx context.Context) error {
	return Run(ctx, "apt.init.yml", playbookAptInit, []string{})
}

func AptUndo(ctx context.Context) error {
	return Run(ctx, "apt.undo.yml", playbookAptUndo, []string{})
}

func DotfilesInit(ctx context.Context) error {
	nvimDst := filepath.Join(dirs.Config, "nvim")
	nvimSrc := filepath.Join(DotfilesDir, "nvim")
	if err := dirs.WriteDir(filepath.Join("assets", "dotfiles", "nvim"), nvimSrc, nvim); err != nil {
		return err
	}

	starshipDst := filepath.Join(dirs.Config, "starship.toml")
	starshipSrc := filepath.Join(DotfilesDir, "starship.toml")
	if _, err := dirs.WriteFile(starshipSrc, starship); err != nil {
		return err
	}

	tmuxDst := filepath.Join(dirs.Config, "tmux", "tmux.conf")
	tmuxSrc := filepath.Join(DotfilesDir, "tmux", "tmux.conf")
	if _, err := dirs.WriteFile(tmuxSrc, tmux); err != nil {
		return err
	}

	return Run(ctx, "dotfiles.init.yml", playbookDotfilesInit, []string{
		fmt.Sprintf("STARSHIP_SRC=%s", starshipSrc),
		fmt.Sprintf("STARSHIP_DST=%s", starshipDst),
		fmt.Sprintf("TMUX_SRC=%s", tmuxSrc),
		fmt.Sprintf("TMUX_DST=%s", tmuxDst),
		fmt.Sprintf("NVIM_SRC=%s", nvimSrc),
		fmt.Sprintf("NVIM_DST=%s", nvimDst),
	})
}

func DotfilesUndo(ctx context.Context) error {
	starshipDst := filepath.Join(dirs.Config, "starship.toml")
	tmuxDst := filepath.Join(dirs.Config, "tmux", "tmx.conf")
	nvimDst := filepath.Join(dirs.Config, "nvim")

	return Run(ctx, "dotfiles.undo.yml", playbookDotfilesUndo, []string{
		fmt.Sprintf("STARSHIP_DST=%s", starshipDst),
		fmt.Sprintf("TMUX_DST=%s", tmuxDst),
		fmt.Sprintf("NVIM_DST=%s", nvimDst),
	})
}

func DockerInit(ctx context.Context) error {
	return Run(ctx, "docker.init.yml", playbookDockerInit, []string{})
}

func DockerUndo(ctx context.Context) error {
	return Run(ctx, "docker.undo.yml", playbookDockerUndo, []string{})
}

func GitInit(ctx context.Context, username string, token string, name string, email string) error {
	return Run(ctx, "git.init.yml", playbookGitInit, []string{
		fmt.Sprintf("GH_UNAME=%s", username),
		fmt.Sprintf("GH_TOKEN=%s", token),
		fmt.Sprintf("GH_EMAIL=%s", email),
		fmt.Sprintf("GH_NAME=%s", name),
	})
}

func GitUndo(ctx context.Context) error {
	return Run(ctx, "git.undo.yml", playbookGitUndo, []string{})
}

func BashrcInit(ctx context.Context) error {
	return Run(ctx, "bashrc.init.yml", playbookBashrcInit, []string{})
}

func BashrcUndo(ctx context.Context) error {
	return Run(ctx, "bashrc.undo.yml", playbookBashrcUndo, []string{})
}
