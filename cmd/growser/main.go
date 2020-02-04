package main

import (
	"os"

	"github.com/cuongcb/growser/pkg/action"
	"github.com/cuongcb/growser/pkg/config"
	"github.com/cuongcb/growser/pkg/log"
	"github.com/cuongcb/growser/pkg/service"
	"github.com/cuongcb/growser/pkg/storage"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "An application browses workspace in golang"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		action.InitAction,
		action.AddAction,
		action.RemoveAction,
		action.BrowseAction,
		action.ListAction,
	}

	log.SetLogLevel(log.ERROR)
	cfg := config.New().WithStorage(storage.File)
	service.Init(cfg)

	app.Run(os.Args)
}
