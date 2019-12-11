package action

import (
	"github.com/cuongcb/growser/pkg/service"
	"github.com/urfave/cli"
)

// RemoveAction moves/open a new terminal to specific project's path
var RemoveAction = cli.Command{
	Name:    "remove",
	Aliases: []string{"rm"},
	Usage:   "rm project",
	Flags:   removeFlags,
	Action:  removeAction,
}

var removeFlags = []cli.Flag{
	cli.StringFlag{
		Name:     "name, n",
		Usage:    "project name",
		Required: true,
	},
}

func removeAction(ctx *cli.Context) error {
	name := ctx.String("name")

	return service.RemoveProject(name)
}
