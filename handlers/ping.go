package handlers

import (
	"fmt"
	"github.com/ryzmae/ioson/utils"
	"github.com/urfave/cli/v2"
)

func Ping(cCtx *cli.Context) error {
	host := cCtx.Args().Get(0)
	
	if host == "" {
		fmt.Println("Error: No host provided")
		return nil
	}

	handlers.HandlerPing(host)
	
	return nil
}