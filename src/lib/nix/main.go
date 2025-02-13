package nix

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"ubctl/src/lib/files"
	"ubctl/src/lib/paths"
)

//go:embed assets/devsh/flake.nix
var DevShFlake []byte

var (
	DevShFlakePath string
	NixDir         string
)

func init() {
	NixDir = filepath.Join(paths.AppCacheDir, "nix")
	DevShFlakePath = filepath.Join(NixDir, "devsh", "flake.nix")
}

func NixDevelop(profile string) error {
	// HACK: if `nix develop` is used in a Git repo and we pass it the path to a directory that
	// contains a flake.nix file which lives outside the repo, then this will result in an error.
	// To get around this, we use two `nix develop` commands. The first one saves our custom nix
	// dev shell build env to a profile WITHOUT starting the shell yet (if we were to start the
	// shell and save the profile in one command then the user's CWD would be changed to the dir
	// where we saved the profile which is not ideal). Once we save the profile, we will then call
	// `nix develop` again this time passing in the path to the newly created profile. The second
	// call will ensure that the user's CWD does not change unexpectedly and it also allows us to
	// avoid errors with Nix not "seeing" any untracked Git files.
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
		if i == 0 {
			cmd.Dir = filepath.Dir(profile)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func NixDevelopDevShell() error {
	if _, err := files.WriteFile(DevShFlakePath, DevShFlake); err != nil {
		return err
	} else {
		return NixDevelop(
			filepath.Join(
				filepath.Dir(DevShFlakePath),
				fmt.Sprintf("%x", md5.Sum(DevShFlake)),
			),
		)
	}
}
