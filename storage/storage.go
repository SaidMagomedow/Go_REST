package storage

import (
	"sync"
)

type Storage interface {
	Insert(a *Author)
	Get(id string) (Author, error)
	Update(id int, a *Author) (Author, error)
	Delete(id int) error
}

type Author struct {
	id   int    `json:"id"`
	name string `json:"name"`
	city string `json:"city"`
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

func (s *RunTimeMemoryStorage) Insert(a *Author) {
	s.Lock()
	a.id = s.counter
	s.data[a.id] = *a

	s.counter++

	s.Unlock()
}

func (s *RunTimeMemoryStorage) Get(id int) (Author, error) {
	return s.data[id], nil
}

func (s *RunTimeMemoryStorage) Update(id int, a *Author) (Author, error) {
	oldAuthor := s.data[a.id]
	oldAuthor.name = a.name
	oldAuthor.city = a.city
	return oldAuthor, nil
}

func (s *RunTimeMemoryStorage) Delete(id int) error {
	delete(s.data, id)
	return nil
}
