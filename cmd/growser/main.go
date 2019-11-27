package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cuongcb/growser/pkg/mapper"
	"github.com/cuongcb/growser/pkg/presenter"
	"github.com/cuongcb/growser/pkg/proto"
)

func initLoader() (mapper.Mapper, error) {
	cfg := &mapper.Config{Type: mapper.File}
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
		fmt.Print("Action > ")
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
			fmt.Println("added new project")
		case "u":
			p, err := inputProject(r)
			if err != nil {
				panic(err)
			}
			err = updateProject(m, p)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("updated project")
		case "r":
			fmt.Printf("> (name) ")
			name, err := r.ReadString('\n')
			if err != nil {
				panic(err)
			}

			name = strings.TrimSuffix(name, "\n")

			err = removeProject(m, name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("removed project")
		case "c":
			err := cleanProject(m)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("removed all projects")
		case "h":
			showHelp()
		case "q":
			fmt.Println("growser stopped...")
			return
		default:
			fmt.Printf("Unsupported action: '%s'\n", c)
			fmt.Println("'h' for help")
		}
	}
}

func inputProject(r *bufio.Reader) (*proto.Project, error) {
	fmt.Printf("> (name) ")
	name, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	name = strings.TrimSuffix(name, "\n")

	fmt.Printf("> (path) ")
	path, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	path = strings.TrimSuffix(path, "\n")

	return &proto.Project{
		Name: name,
		Path: path,
	}, nil
}

func addProject(m mapper.Mapper, p *proto.Project) error {
	return m.Add(p.Name, p.Path)
}

func listProject(m mapper.Mapper, p presenter.Presenter) {
	list, _ := m.List()
	p.Present(list)
}

func removeProject(m mapper.Mapper, name string) error {
	return m.Remove(name)
}

func updateProject(m mapper.Mapper, p *proto.Project) error {
	return m.Update(p.Name, p.Path)
}

func cleanProject(m mapper.Mapper) error {
	return m.Clean()
}

func showHelp() {
	fmt.Println("growser: l, r, a, u, c, h, q")
	fmt.Println("- l: list all projects")
	fmt.Println("- a: add a project")
	fmt.Println("- u: update an existing project")
	fmt.Println("- r: remove a project")
	fmt.Println("- c: remove all projects")
	fmt.Println("- h: show help")
	fmt.Println("- q: quit")
}
