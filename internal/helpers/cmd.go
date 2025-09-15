package helpers

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

func RunCommand(sigs <-chan os.Signal, args []string) error {
	var cmd *exec.Cmd
	var isTerm bool = false
	var err error
	var app string
	var prs []string

	if len(args) > 0 {
		app = args[0]
	}

	if len(args) > 1 {
		prs = args[1:]
	}

	cmd = exec.Command(app, prs...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	go func(sigs <-chan os.Signal, cmd *exec.Cmd, isTerm *bool) {
		<-sigs

		if cmd.Process != nil {
			*isTerm = true
			cmd.Process.Signal(syscall.SIGTERM)
		}
	}(sigs, cmd, &isTerm)

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("launch error: %s", err.Error())
	}

	err = cmd.Wait()
	if err != nil && !isTerm {
		return fmt.Errorf("running error: %s", err.Error())
	}

	return nil
}

func FindInPath(appName string) ([]string, error) {
	var lines []string
	var stdout, stderr bytes.Buffer
	var cmd *exec.Cmd = nil

	if runtime.GOOS == "windows" {
		cmd = exec.Command("where", appName)
	} else {
		cmd = exec.Command("which", "-a", appName)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = `unknown error: ` + err.Error()
		}

		return nil, errors.New(msg)
	}

	sc := bufio.NewScanner(&stdout)
	sc.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)

	for sc.Scan() {
		line := filepath.Dir(strings.TrimSpace(sc.Text()))
		if line != "" && !Contains(lines, line) {
			lines = append(lines, line)
		}
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
