package action

import (
	"os"
	"path"

	"github.com/cuongcb/growser/pkg/service"
	"github.com/urfave/cli"
)

// InitAction provides the behavior for adding current project
var InitAction = cli.Command{
	Name:   "init",
	Usage:  "save current project info to db",
	Action: initAction,
}

func initAction(ctx *cli.Context) error {
	fullPath, err := os.Getwd()
	if err != nil {
		return err
	}

	name := path.Base(fullPath)

	return service.AddProject(name, fullPath)
}
