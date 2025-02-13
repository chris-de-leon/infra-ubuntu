package paths

import (
	"os"
	"path/filepath"
	"ubctl/src/lib/version"
)

const (
	UBCTL = "ubctl"
)

var (
	UserConfigDir string
	UserCacheDir  string
	AppConfigDir  string
	AppCacheDir   string
)

func init() {
	var err error

	UserConfigDir, err = os.UserConfigDir()
	if err != nil {
		panic(err)
	} else {
		AppConfigDir = filepath.Join(UserConfigDir, UBCTL, version.WithPrefix())
	}

	UserCacheDir, err = os.UserCacheDir()
	if err != nil {
		panic(err)
	} else {
		AppCacheDir = filepath.Join(UserCacheDir, UBCTL, version.WithPrefix())
	}
}

func RemoveAppConfigDir() error {
	return os.RemoveAll(filepath.Dir(AppConfigDir))
}

func RemoveAppCacheDir() error {
	return os.RemoveAll(filepath.Dir(AppCacheDir))
}
