package main

import (
	"fmt"
	"github.com/cuongcb/growser/pkg/mapper"
	"github.com/cuongcb/growser/pkg/presenter"
	"os"
)

type project struct {
	name, path string
}

func initLoader() (mapper.Mapper, error) {
	cfg := &mapper.Config{Type: mapper.InMem}
	return mapper.New(cfg)
}

func initPresenter() (presenter.Presenter, error) {
	return presenter.New(), nil
}

func main() {
	fmt.Println("growser starting...")

	m, err := initLoader()
	if err != nil {
		panic(err)
	}

	p, err := initPresenter()
	if err != nil {
		panic(err)
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	m.Add(os.Args[1], path)

	list, err := m.List()
	if err != nil {
		panic(err)
	}

	p.Present(list)
}

func addProject(m mapper.Mapper, p project) error {
	return m.Add(p.name, p.path)
}

func listProject(m mapper.Mapper, p presenter.Presenter) {
	list, _ := m.List()
	p.Present(list)
}

func removeProject(m mapper.Mapper, name string) error {
	return m.Remove(name)
}
