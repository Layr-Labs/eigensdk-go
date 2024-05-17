package main

import (
	"fmt"
	"os"

	"github.com/Layr-Labs/eigensdk-go/cmd/egnkey/derive_operatorid"
	"github.com/Layr-Labs/eigensdk-go/cmd/egnkey/generate"
	"github.com/Layr-Labs/eigensdk-go/cmd/egnkey/store"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "egnkey"
	app.Description = "Eigenlayer batch keys manager"
	app.Commands = []*cli.Command{
		generate.Command,
		store.Command,
		derive_operatorid.Command,
	}

	app.Usage = "Used to manage batch keys for testing"

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}
