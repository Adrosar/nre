// https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
package helpers

import (
	"os/exec"
	"runtime"
)

func OpenURL(url string) error {
	var app string
	var args []string

	switch runtime.GOOS {
	case "windows":
		app = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		app = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		app = "xdg-open"
	}

	args = append(args, url)
	cmd := exec.Command(app, args...)
	_, err := cmd.CombinedOutput()
	return err
}
