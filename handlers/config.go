package handlers

import (
	"fmt"
	"strconv"
	"github.com/ryzmae/ioson/utils"
	"github.com/urfave/cli/v2"
)

func Config(cCtx *cli.Context) error {
	ttl := cCtx.Args().Get(0)
	
	sttl, err := strconv.Atoi(ttl)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	handlers.HandlerConfig(sttl)

	fmt.Println("TTL set to", ttl)
	
	return nil
}