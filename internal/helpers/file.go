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
	var sep = string(os.PathSeparator)
	var err error = nil
	var name string = ""
	var pth string = ""
	var wd string = ""
	var ok bool = false

	pth = target + sep + `package.json`
	ok, err = IsExist(pth)

	if !ok {
		return fmt.Errorf("folder \"%s\" is not an NPM package", target)
	}

	if err != nil {
		return fmt.Errorf("packet check error: %s", err.Error())
	}

	wd, err = os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get working directory: %s", err.Error())
	}

	name = filepath.Base(target)
	pth = wd + sep + `node_modules` + sep + name
	ok, err = IsExist(pth)

	if ok {
		return fmt.Errorf("package \"%s\" already exists", name)
	}

	if err != nil {
		return fmt.Errorf("failed to check location \"%s\" due to error: %s", pth, err.Error())
	}

	err = os.Symlink(target, pth)
	if err != nil {
		return fmt.Errorf("unable to create link: %s", err.Error())
	}

	return nil
}
