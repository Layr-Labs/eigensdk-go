package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Eigenlayer batch key manager"
	app.Commands = []*cli.Command{
		commandGenerate,
	}

	if err := app.Run(os.Args); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

}
