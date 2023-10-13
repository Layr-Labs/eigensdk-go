package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Eigenlayer batch keys manager"
	app.Commands = []*cli.Command{
		commandGenerate,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}
