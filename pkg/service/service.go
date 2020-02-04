package service

import (
	"github.com/cuongcb/growser/pkg/browser"
	"github.com/cuongcb/growser/pkg/config"
	"github.com/cuongcb/growser/pkg/log"
	"github.com/cuongcb/growser/pkg/storage"
	"github.com/cuongcb/growser/pkg/view"
)

type cliServices struct {
	storage.Mapper
	browser.Browser
	view.Presenter
}

var client cliServices

// Init initializes service
func Init(cfg *config.Config) error {
	mapper, err := storage.NewMapper(cfg.StorageType, cfg.DBPath)
	if err != nil {
		log.Error("failed in creating mapper, detailed %q", err)
		return err
	}

	client = cliServices{
		mapper,
		browser.New(),
		view.NewPresenter(),
	}

	return nil
}

// AddProject adds new project to db
func AddProject(name, path string) error {
	if err := client.Add(name, path); err != nil {
		log.Error("add project [%s] at [%s] failed, detailed %q", name, path, err)
		return err
	}

	return nil
}

// RemoveProject removes an existing project
func RemoveProject(name string) error {
	if err := client.Remove(name); err != nil {
		log.Error("remove project [%s] failed, detailed %q", name, err)
		return err
	}

	return nil
}

// GotoProject opens a new terminal with registerd project's path
func GotoProject(name string) error {
	path, err := client.Get(name)
	if err != nil {
		log.Error("project [%s] not found, detailed %q", name, err)
		return err
	}

	if err := client.Go(path); err != nil {
		log.Error("goto project [%s] at [%s] failed, detailed %q", name, path, err)
	}

	return nil
}

// ListProject shows all saved projects
func ListProject() error {
	list, err := client.List()
	if err != nil {
		log.Error("read db failed %v", err)
		return err
	}

	client.Present(list)
	return nil
}
