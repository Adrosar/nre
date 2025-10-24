package cli

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Adrosar/nre/internal/helpers"
)

func Run() {
	l := len(os.Args)

	if l <= 1 {
		fmt.Fprintln(os.Stdout, INFO)
		return
	}

	if l == 2 {
		arg1 := os.Args[1]
		switch arg1 {
		case `help`:
			fmt.Fprintln(os.Stdout, HELP)
		case `req`:
			fmt.Fprintln(os.Stdout, REQ)
		case `www`:
			helpers.OpenURL(`https://github.com/Adrosar/nre`)
		case `list`:
			fmt.Fprintf(os.Stdout, "NodeJS versions:\n")

			result, _ := helpers.NvmList()
			for key := range result {
				fmt.Fprintf(os.Stdout, " -> %s\n", key)
			}

			fmt.Fprintf(os.Stdout, "in \"%s\"\n", helpers.NvmHome())
		case `link`:
			fmt.Fprintf(os.Stdout, "Link command requires two parameters!")
		default:
			fmt.Fprintf(os.Stdout, `Unknown command "%s" !`, arg1)
		}

		return
	}

	if l == 3 && os.Args[1] == `link` {
		err := helpers.Link(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create link: %s", err.Error())
			os.Exit(1)
		}

		return
	}

	if l > 2 {
		njsVer := os.Args[1]
		cmdArgs := os.Args[2:]
		cmdText := strings.Join(cmdArgs, " ")

		ned, err := helpers.FindPathForNode(njsVer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "directory for NodeJS (%s) does not exist: %s\n", njsVer, err.Error())
			os.Exit(2)
		}

		wh, err := helpers.FindInPath(`node`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to find NodeJS location: %s\n", err.Error())
			os.Exit(3)
		}

		nmb, err := helpers.NodeModulesBinDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "it is impossible to determine whether the project contains additional commands: %s\n", err.Error())
			os.Exit(4)
		}

		envPath := helpers.ExpandPathList(helpers.ReducePathList(os.Getenv(`PATH`), wh), nmb, ned)
		os.Setenv(`PATH`, envPath)

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

		fmt.Fprintf(os.Stdout, "Run command \"%s\" with NodeJS (%s)\n", cmdText, njsVer)
		err = helpers.RunCommand(sigs, cmdArgs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "run \"%v\" failed: %s\n", strings.Join(cmdArgs, " "), err.Error())
			os.Exit(5)
		}
	}
}
