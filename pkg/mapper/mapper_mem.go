package mapper

type memMapper struct {
	projectMap map[string]string
}

func newMemMapper() *memMapper {
	return &memMapper{
		projectMap: make(map[string]string, 0),
	}
}

func (m *memMapper) List() (map[string]string, error) {
	if err := must(m); err != nil {
		return nil, err
	}

	return m.projectMap, nil
}

func (m *memMapper) Add(k, v string) error {
	if err := must(m); err != nil {
		return err
	}

	if _, ok := m.projectMap[k]; ok {
		// existed key
		return errDuplicatedRecord
	}

	m.projectMap[k] = v
	return nil
}

func (m *memMapper) Update(k, v string) error {
	if err := must(m); err != nil {
		return err
	}

	if _, ok := m.projectMap[k]; !ok {
		// not existed
		return errNotFoundRecord
	}

	m.projectMap[k] = v
	return nil
}

func (m *memMapper) Remove(k string) error {
	if err := must(m); err != nil {
		return err
	}

	if _, ok := m.projectMap[k]; !ok {
		// not existed
		return errNotFoundRecord
	}

	delete(m.projectMap, k)
	return nil
}

func (m *memMapper) Clean() error {
	if err := must(m); err != nil {
		return err
	}

	m.projectMap = make(map[string]string, 0)
	return nil
}

func (m *memMapper) Info(k string) (string, error) {
	if err := must(m); err != nil {
		return "", err
	}

	return "", nil
}
