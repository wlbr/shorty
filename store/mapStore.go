package store

import (
	"fmt"
)

type MapStore struct {
	urls   map[int]string
	latest int
}

func (m *MapStore) Add(todo string) error {
	if m.latest == 0 {
		m.urls = make(map[int]string)
	}
	m.latest++
	m.urls[m.latest] = todo
	return nil
}

func (m *MapStore) List() (result string) {
	result = "MapStore\n"
	for index, url := range m.urls {
		result = fmt.Sprintf("%s%d: %s\n", result, index, url)
	}
	return result
}

func (m *MapStore) Delete(id int) (err error) {
	delete(m.urls, id)
	return nil
}

func (m *MapStore) Length() int {
	return len(m.urls)
}

func (m *MapStore) Update(id int, newurl string) (err error) {
	_, found := m.urls[id]
	if !found {
		err = fmt.Errorf("no shorturl found with localpart %d", id)
	} else {
		m.urls[id] = newurl
	}

	return err
}
