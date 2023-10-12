package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var app *cli.App

	app = cli.NewApp()
	app.Name = "Eigenlayer batch key manager"
	app.Commands = []*cli.Command{
		commandGenerate,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
