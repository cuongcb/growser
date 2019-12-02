package service

import (
	"github.com/cuongcb/growser/pkg/browser"
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
func Init() error {
	return nil
}

// AddProject adds new project to db
func AddProject(name, path string) error {
	return client.Add(name, path)
}

// RemoveProject removes an existing project
func RemoveProject(name string) error {
	return client.Remove(name)
}

// GotoProject opens a new terminal with registerd project's path
func GotoProject(name string) error {
	path, err := client.Get(name)
	if err != nil {
		return err
	}

	return client.Go(path)
}

// ListProject shows all saved projects
func ListProject() error {
	list, err := client.List()
	if err != nil {
		return err
	}

	client.Present(list)
	return nil
}
