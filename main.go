package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:      "gz",
		Usage:     "The GZ Family",
		UsageText: "gz <command> [args...]",
		Version:   Version,
		Flags:     []cli.Flag{},
		// HideHelp:  true,
	})

	app.Command(func(ctx *cli.Context) error {
		if ctx.Args().Len() == 0 {
			return nil
		}

		subCommandName := ctx.Args().Get(0)
		subCommandArgs := ctx.Args().Slice()[1:]

		switch subCommandName {
		case "list":
			return nil
		default:
			cmd := exec.Command(fmt.Sprintf("gz%s", subCommandName), subCommandArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			return cmd.Run()
		}
	})

	app.Run()
}
