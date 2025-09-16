package helpers

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func NvmHome() string {
	if runtime.GOOS != `windows` {
		return os.Getenv(`NVM_DIR`)
	}

	return os.Getenv(`NVM_HOME`)
}

func NvmList() (map[string]string, error) {
	sep := string(os.PathSeparator)
	result := make(map[string]string, 10)

	home := NvmHome()
	if runtime.GOOS != `windows` {
		home = home + sep + `versions` + sep + `node`
	}

	entries, err := os.ReadDir(home)
	if err != nil {
		return result, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dir := entry.Name()
			ok := strings.HasPrefix(dir, `v`)
			if ok {
				end := ``
				if runtime.GOOS != `windows` {
					end = sep + `bin`
				}

				full := home + sep + dir + end
				result[dir] = full
			}
		}
	}

	return result, nil
}

func FindPathForNode(nodeVer string) (string, error) {
	result, err := NvmList()
	if err != nil {
		return "", err
	}

	for key, val := range result {
		if strings.Contains(key, `v`+nodeVer) {
			return val, nil
		}
	}

	return "", fmt.Errorf("not found NodeJS (ver %s)", nodeVer)
}
