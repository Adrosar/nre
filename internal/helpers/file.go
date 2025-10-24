package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func IsExist(p string) (bool, error) {
	if _, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func Link(target string) error {
	var err error = nil

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("link: getwd: %s", err.Error())
	}

	sep := string(os.PathSeparator)
	err = os.Symlink(target, wd+sep+`node_modules`+sep+filepath.Base(target))
	if err != nil {
		return fmt.Errorf("link: symlink: %s", err.Error())
	}

	return nil
}
