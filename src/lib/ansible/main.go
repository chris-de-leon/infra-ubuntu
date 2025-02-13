package ansible

import (
	"context"
	"crypto/md5"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"ubctl/src/lib/files"
	"ubctl/src/lib/paths"
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
	AnsibleFlakePath string
	PlaybooksDir     string
	DotfilesDir      string
	AnsibleDir       string
	StarshipSrc      string
	StarshipDst      string
	TmuxSrc          string
	TmuxDst          string
	NvimSrc          string
	NvimDst          string
	NvimDir          string
)

func init() {
	AnsibleDir = filepath.Join(paths.AppCacheDir, "ansible")
	AnsibleFlakePath = filepath.Join(AnsibleDir, "flake.nix")
	PlaybooksDir = filepath.Join(AnsibleDir, "playbooks")
	DotfilesDir = filepath.Join(AnsibleDir, "dotfiles")
	StarshipSrc = filepath.Join(DotfilesDir, "starship.toml")
	StarshipDst = filepath.Join(paths.UserConfigDir, "starship.toml")
	TmuxSrc = filepath.Join(DotfilesDir, "tmux", "tmux.conf")
	TmuxDst = filepath.Join(paths.UserConfigDir, "tmux", "tmux.conf")
	NvimSrc = filepath.Join(DotfilesDir, "nvim")
	NvimDst = filepath.Join(paths.UserConfigDir, "nvim")
	NvimDir = filepath.Join("assets", "dotfiles", "nvim")
}

func Run(ctx context.Context, data []byte, vars []string) error {
	name := fmt.Sprintf("%x.yml", md5.Sum(data))

	playbookPath, err := files.WriteFile(filepath.Join(PlaybooksDir, name), data)
	if err != nil {
		return err
	}

	flakePath, err := files.WriteFile(AnsibleFlakePath, flake)
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
	return Run(ctx, playbookAptInit, []string{})
}

func AptUndo(ctx context.Context) error {
	return Run(ctx, playbookAptUndo, []string{})
}

func DotfilesInit(ctx context.Context) error {
	if err := files.WriteDir(NvimDir, NvimSrc, nvim); err != nil {
		return err
	}

	if _, err := files.WriteFile(StarshipSrc, starship); err != nil {
		return err
	}

	if _, err := files.WriteFile(TmuxSrc, tmux); err != nil {
		return err
	}

	return Run(ctx, playbookDotfilesInit, []string{
		fmt.Sprintf("STARSHIP_SRC=%s", StarshipSrc),
		fmt.Sprintf("STARSHIP_DST=%s", StarshipDst),
		fmt.Sprintf("TMUX_SRC=%s", TmuxSrc),
		fmt.Sprintf("TMUX_DST=%s", TmuxDst),
		fmt.Sprintf("NVIM_SRC=%s", NvimSrc),
		fmt.Sprintf("NVIM_DST=%s", NvimDst),
	})
}

func DotfilesUndo(ctx context.Context) error {
	return Run(ctx, playbookDotfilesUndo, []string{
		fmt.Sprintf("STARSHIP_DST=%s", StarshipDst),
		fmt.Sprintf("TMUX_DST=%s", TmuxDst),
		fmt.Sprintf("NVIM_DST=%s", NvimDst),
	})
}

func DockerInit(ctx context.Context) error {
	return Run(ctx, playbookDockerInit, []string{})
}

func DockerUndo(ctx context.Context) error {
	return Run(ctx, playbookDockerUndo, []string{})
}

func GitInit(ctx context.Context, username string, token string, name string, email string) error {
	return Run(ctx, playbookGitInit, []string{
		fmt.Sprintf("GH_UNAME=%s", username),
		fmt.Sprintf("GH_TOKEN=%s", token),
		fmt.Sprintf("GH_EMAIL=%s", email),
		fmt.Sprintf("GH_NAME=%s", name),
	})
}

func GitUndo(ctx context.Context) error {
	return Run(ctx, playbookGitUndo, []string{})
}

func BashrcInit(ctx context.Context) error {
	return Run(ctx, playbookBashrcInit, []string{})
}

func BashrcUndo(ctx context.Context) error {
	return Run(ctx, playbookBashrcUndo, []string{})
}
