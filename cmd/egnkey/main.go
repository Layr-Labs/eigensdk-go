package main

import (
	"fmt"
	"os"

	"github.com/Layr-Labs/eigensdk-go/cmd/egnkey/generate"
	"github.com/Layr-Labs/eigensdk-go/cmd/egnkey/operatorid"
	"github.com/Layr-Labs/eigensdk-go/cmd/egnkey/store"
	"github.com/urfave/cli/v2"
)

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

// We structure run in this way to make it testable.
// see https://github.com/urfave/cli/issues/731
func run(args []string) error {
	app := cli.NewApp()
	app.Name = "egnkey"
	app.Description = "Eigenlayer batch keys manager"
	app.Commands = []*cli.Command{
		generate.Command,
		store.Command,
		operatorid.Command,
	}
	app.Usage = "Used to manage batch keys for testing"

	err := app.Run(args)
	return err
}
