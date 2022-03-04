package cmd

import "github.com/urfave/cli/v2"

func Default(ctx *cli.Context) error {
	return Server(ctx)
}
