package action

import (
	"github.com/cuongcb/growser/pkg/service"
	"github.com/urfave/cli"
)

// ListAction lists all saved projects
var ListAction = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "list all saved projects",
	Flags:   listFlags,
	Action:  listAction,
}

var listFlags = []cli.Flag{}

func listAction(ctx *cli.Context) error {
	return service.ListProject()
}
