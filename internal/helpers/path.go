package helpers

import (
	"os"
	"strings"
)

func ReducePathList(pathList string, toRemoved []string) string {
	var sep = string(os.PathListSeparator)
	var list = strings.Split(pathList, sep)
	var output []string

	for _, a := range list {
		for _, b := range toRemoved {
			if a != b {
				output = append(output, a)
			}
		}
	}

	return strings.Join(output, sep)
}

func ExpandPathList(pathList string, toAdded ...string) string {
	s := string(os.PathListSeparator)
	return strings.Trim(strings.Join(toAdded, s)+s+pathList, s)
}

func CreatePath(elements ...string) string {
	var l = len(elements)
	var b = strings.Builder{}
	for i, p := range elements {
		b.WriteString(p)
		if (i + 1) < l {
			b.WriteRune(os.PathSeparator)
		}
	}

	return b.String()
}

func NodeModulesBinDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	p := CreatePath(wd, `node_modules`, `.bin`)
	ok, err := IsExist(p)
	if err != nil {
		return "", err
	}

	if ok {
		return p, nil
	}

	return "", nil
}
