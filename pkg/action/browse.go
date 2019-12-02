package action

import (
	"github.com/cuongcb/growser/pkg/service"
	"github.com/urfave/cli"
)

// BrowseAction moves/open a new terminal to specific project's path
var BrowseAction = cli.Command{
	Name:   "browse",
	Usage:  "go to project",
	Flags:  browseFlags,
	Action: browseAction,
}

var browseFlags = []cli.Flag{
	cli.StringFlag{
		Name:     "name, n",
		Usage:    "project name",
		Required: true,
	},
}

func browseAction(ctx *cli.Context) error {
	name := ctx.String("name")

	return service.GotoProject(name)
}
