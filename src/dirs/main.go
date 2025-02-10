package dirs

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"ubctl/src/version"
)

const (
	APPLICATION_DIR = "ubctl"
)

var (
	AppConfig string
	AppCache  string
	Config    string
	Cache     string
)

func init() {
	var err error

	Config, err = os.UserConfigDir()
	if err != nil {
		panic(err)
	} else {
		AppConfig = filepath.Join(Config, APPLICATION_DIR, version.WithPrefix())
	}

	Cache, err = os.UserCacheDir()
	if err != nil {
		panic(err)
	} else {
		AppCache = filepath.Join(Cache, APPLICATION_DIR, version.WithPrefix())
	}
}

func WriteDir(src string, dst string, data embed.FS) error {
	dir, err := fs.Sub(data, src)
	if err != nil {
		return err
	}

	return fs.WalkDir(dir, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		target := filepath.Join(dst, path)
		if entry.IsDir() {
			if err := os.MkdirAll(target, os.ModePerm); err != nil {
				return err
			}
		} else {
			item, err := data.ReadFile(filepath.Join(src, path))
			if err != nil {
				return err
			}
			if err := os.WriteFile(target, item, os.ModePerm); err != nil {
				return err
			}
		}

		return nil
	})
}

func WriteFile(dst string, data []byte) (string, error) {
	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return "", err
	}

	if err := os.WriteFile(dst, data, os.ModePerm); err != nil {
		return "", err
	}

	return dst, nil
}
