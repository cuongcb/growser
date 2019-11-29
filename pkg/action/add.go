package action

import (
	"github.com/urfave/cli"
)

// AddAction provides the behavior for adding new project
var AddAction = cli.Command{
	Name:   "add",
	Usage:  "save project info to db",
	Flags:  addFlags,
	Action: addAction,
}

var addFlags = []cli.Flag{
	cli.StringFlag{
		Name:     "name, n",
		Usage:    "project name",
		Required: true,
	},
	cli.StringFlag{
		Name:     "path, p",
		Usage:    "project path",
		Required: true,
	},
}

func addAction(ctx *cli.Context) error {
	name := ctx.String("name")
	path := ctx.String("path")

	return nil
}
