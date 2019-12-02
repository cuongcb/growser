package main

import (
	"fmt"
	"os"

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
	fmt.Println("growser starting...")

	app := cli.NewApp()
	app.Run(os.Args)

	m, err := initLoader()
	if err != nil {
		panic(err)
	}

	p, err := initPresenter()
	if err != nil {
		panic(err)
	}

}
