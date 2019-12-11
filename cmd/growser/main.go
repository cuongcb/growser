package main

import (
	"os"

	"github.com/cuongcb/growser/pkg/action"
	"github.com/cuongcb/growser/pkg/service"
	"github.com/cuongcb/growser/pkg/storage"
	"github.com/cuongcb/growser/pkg/view"

	"github.com/urfave/cli"
)

func initLoader() (storage.Mapper, error) {
	cfg := &storage.Config{Type: storage.File}
	return storage.NewMapper(cfg)
}

func initPresenter() (view.Presenter, error) {
	return view.NewPresenter(), nil
}

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		action.InitAction,
		action.AddAction,
		action.RemoveAction,
		action.BrowseAction,
		action.ListAction,
	}

	service.Init()

	app.Run(os.Args)
}
