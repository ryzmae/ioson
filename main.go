package main

import (
	"fmt"
	"os"

	"github.com/ryzmae/ioson/handlers"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "ioson",
		Usage: "Ioson, a simple cli tool for pinging hosts",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name: "version",
				Aliases: []string{"v"},
				Usage: "Prints the current version of Ioson",
				Action: func(ctx *cli.Context) error {
					handlers.HandlerVersion(ctx)
					return nil
				},
			},
			{
				Name: "config",
				Aliases: []string{"c"},
				Usage: "Set the TTL value for the ping command",
				Action: func(ctx *cli.Context) error {
					handlers.HandlerConfig(ctx)
					return nil
				},
			},
			{
				Name: "ping",
				Aliases: []string{"p"},
				Usage: "Ping a host",
				Action: func(ctx *cli.Context) error {
					handlers.HandlerPing(ctx)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}