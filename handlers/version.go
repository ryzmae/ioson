package handlers

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Version(cCtx *cli.Context) error {
	fmt.Println("Ioson", cCtx.App.Version)
	return nil
}