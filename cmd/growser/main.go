package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cuongcb/growser/pkg/mapper"
	"github.com/cuongcb/growser/pkg/presenter"
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

	showHelp()

	for {
		r := bufio.NewReader(os.Stdin)
		c, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		c = strings.TrimSuffix(c, "\n")

		switch c {
		case "l":
			listProject(m, p)
		case "a":
			p, err := inputProject(r)
			if err != nil {
				panic(err)
			}
			err = addProject(m, p)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("project added")
		case "r":
			name, err := r.ReadString('\n')
			if err != nil {
				panic(err)
			}

			err = removeProject(m, name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("project removed")
		case "h":
			showHelp()
		case "q":
			fmt.Println("growser stopped...")
			return
		}
	}
}

func inputProject(r *bufio.Reader) (project, error) {
	fmt.Printf("> ")
	name, err := r.ReadString('\n')
	if err != nil {
		return project{}, err
	}

	name = strings.TrimSuffix(name, "\n")

	fmt.Printf("> ")
	path, err := r.ReadString('\n')
	if err != nil {
		return project{}, err
	}

	path = strings.TrimSuffix(path, "\n")

	return project{
		name: name,
		path: path,
	}, nil
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

func showHelp() {
	fmt.Println("growser: l, r, a, h, q")
	fmt.Println("- l: list all projects")
	fmt.Println("- a: add a project")
	fmt.Println("- r: remove a project")
	fmt.Println("- h: show help")
	fmt.Println("- q: quit")
}
