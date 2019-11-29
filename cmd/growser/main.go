package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/cuongcb/growser/pkg/proto"
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

	if os.Args[1] == "init" {
		fullPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		name := path.Base(fullPath)

		p := &proto.Project{
			Name: name,
			Path: fullPath,
		}

		if err := addProject(m, p); err != nil {
			log.Fatal(err)
		}
	}

	if os.Args[1] == "go" {
		name := os.Args[2]
		list, err := m.List()
		if err != nil {
			log.Fatal(err)
		}

		if len(list) == 0 {
			log.Fatalf("There is no project: %q", name)
		}

		dir := list[name]

		cmd := exec.Command("gnome-terminal", "--tab")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = dir
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
	}

	if os.Args[1] == "list" {
		list, err := m.List()
		if err != nil {
			log.Fatal(err)
		}

		p.Present(list)
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

func addProject(m storage.Mapper, p *proto.Project) error {
	return m.Add(p.Name, p.Path)
}

func listProject(m storage.Mapper, p view.Presenter) {
	list, _ := m.List()
	p.Present(list)
}

func removeProject(m storage.Mapper, name string) error {
	return m.Remove(name)
}

func updateProject(m storage.Mapper, p *proto.Project) error {
	return m.Update(p.Name, p.Path)
}

func cleanProject(m storage.Mapper) error {
	return m.Clean()
}
