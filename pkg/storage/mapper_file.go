package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cuongcb/growser/pkg/proto"
	gproto "github.com/golang/protobuf/proto"
)

type fileMapper struct {
	file string
	hub  *proto.Hub
}

func newFileMapper(filePath string) *fileMapper {
	dir := filepath.Dir(filePath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}

		// create a file
		if _, err := os.Create(filePath); err != nil {
			panic(err)
		}
	}

	return &fileMapper{
		file: filePath,
		hub: &proto.Hub{
			Projects: make(map[string]*proto.Project),
		},
	}
}

func (m *fileMapper) List() (map[string]string, error) {
	if err := must(m); err != nil {
		return nil, err
	}

	if err := m.deserialize(); err != nil {
		return nil, err
	}

	list := make(map[string]string)
	projects := m.hub.GetProjects()
	for _, p := range projects {
		list[p.GetName()] = p.GetPath()
	}

	return list, nil
}

func (m *fileMapper) Get(name string) (string, error) {
	if err := must(m); err != nil {
		return "", err
	}

	if err := m.deserialize(); err != nil {
		return "", err
	}

	projects := m.hub.GetProjects()
	if p, ok := projects[name]; ok {
		return p.GetPath(), nil
	}

	return "", errNotFoundRecord
}

func (m *fileMapper) Add(name, path string) error {
	if err := must(m); err != nil {
		return err
	}

	if err := m.deserialize(); err != nil {
		return err
	}

	projects := m.hub.GetProjects()
	if _, ok := projects[name]; ok {
		return errDuplicatedRecord
	}

	projects[name] = &proto.Project{
		Name: name,
		Path: path,
	}

	return m.serialize()
}

func (m *fileMapper) Update(name, path string) error {
	if err := must(m); err != nil {
		return err
	}

	if err := m.deserialize(); err != nil {
		return err
	}

	projects := m.hub.GetProjects()
	if _, ok := projects[name]; !ok {
		return errNotFoundRecord
	}

	projects[name].Path = path

	return m.serialize()
}

func (m *fileMapper) Remove(name string) error {
	if err := must(m); err != nil {
		return err
	}

	if err := m.deserialize(); err != nil {
		return err
	}

	projects := m.hub.GetProjects()
	if _, ok := projects[name]; !ok {
		return errNotFoundRecord
	}

	delete(projects, name)

	return m.serialize()
}

func (m *fileMapper) Clean() error {
	if err := must(m); err != nil {
		return err
	}

	m.hub.Reset()
	return m.serialize()
}

func (m *fileMapper) Info(name string) (string, error) {
	if err := must(m); err != nil {
		return "", err
	}

	return "", nil
}

func (m *fileMapper) serialize() error {
	buf, err := gproto.Marshal(m.hub)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(m.file, buf, 0644); err != nil {
		return err
	}

	return nil
}

func (m *fileMapper) deserialize() error {
	file, err := os.Open(m.file)
	if err != nil {
		return err
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	h := &proto.Hub{}

	if err := gproto.Unmarshal(buf, h); err != nil {
		return err
	}

	if len(h.Projects) != 0 {
		m.hub = h
	}

	return nil
}
