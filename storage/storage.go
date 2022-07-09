package storage

import (
	"sync"
)

type Storage interface {
	Create(a *Author)
	Get(id int) (Author, error)
	Update(id int, a *Author) error
	Delete(id int) error
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type RunTimeMemoryStorage struct {
	data    map[int]Author
	counter int
	sync.Mutex
}

func NewRunTimeMemoryStorage() *RunTimeMemoryStorage {
	return &RunTimeMemoryStorage{
		data:    make(map[int]Author),
		counter: 1,
	}
}

func (s *RunTimeMemoryStorage) Create(a *Author) {
	s.Lock()
	a.ID = s.counter
	s.data[a.ID] = *a

	s.counter++

	s.Unlock()
}

func (s *RunTimeMemoryStorage) Get(id int) (Author, error) {
	return s.data[id], nil
}

func (s *RunTimeMemoryStorage) Update(id int, a *Author) error {
	s.data[id] = *a
	return nil
}

func (s *RunTimeMemoryStorage) Delete(id int) error {
	delete(s.data, id)
	return nil
}
