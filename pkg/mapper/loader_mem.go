package mapper

type memLoader struct {
	projectMap map[string]string
}

func newMemLoader() *memLoader {
	return &memLoader{
		projectMap: make(map[string]string, 0),
	}
}

func (l *memLoader) List() (map[string]string, error) {
	if err := must(l); err != nil {
		return nil, err
	}

	return l.projectMap, nil
}

func (l *memLoader) Add(k, v string) error {
	if err := must(l); err != nil {
		return err
	}

	if _, ok := l.projectMap[k]; ok {
		// existed key
		return errDuplicatedRecord
	}

	l.projectMap[k] = v
	return nil
}

func (l *memLoader) Update(k, v string) error {
	if err := must(l); err != nil {
		return err
	}

	if _, ok := l.projectMap[k]; !ok {
		// not existed
		return errNotFoundRecord
	}

	l.projectMap[k] = v
	return nil
}

func (l *memLoader) Remove(k string) error {
	if err := must(l); err != nil {
		return err
	}

	if _, ok := l.projectMap[k]; !ok {
		// not existed
		return errNotFoundRecord
	}

	delete(l.projectMap, k)
	return nil
}

func (l *memLoader) Info(k string) (string, error) {
	if err := must(l); err != nil {
		return "", err
	}

	return "", nil
}
