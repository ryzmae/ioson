package handlers

import (
	"github.com/urfave/cli/v2"
)

func HandlerVersion(cCtx *cli.Context) error {
	Version(cCtx)
	return nil
}

func HandlerConfig(cCtx *cli.Context) error {
	Config(cCtx)
	return nil
}

func HandlerPing(cCtx *cli.Context) error {
	Ping(cCtx)
	return nil
}