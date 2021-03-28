package StorageEngines

import (
	"errors"
	"sync"
)

type InMemoryStorage struct {
	content map[string][]byte
	mux     sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		content: map[string][]byte{},
		mux:     sync.Mutex{},
	}
}

func (i *InMemoryStorage) Store(id string, content []byte) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.content[id] = content
	return nil
}

func (i *InMemoryStorage) FetchOne(id string) ([]byte, error) {
	i.mux.Lock()
	defer i.mux.Unlock()
	result, ok := i.content[id]
	if !ok {
		return nil, errors.New("index not found")
	}
	return result, nil
}

func (i *InMemoryStorage) FetchAll() (map[string][]byte, error) {
	i.mux.Lock()
	defer i.mux.Unlock()
	return i.content, nil
}

func (i *InMemoryStorage) Delete(id string) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	if _, ok := i.content[id]; !ok {
		return errors.New("key not found")
	}
	delete(i.content, id)
	return nil
}

func (i *InMemoryStorage) KeyExists(id string) (bool, error) {
	_, ok := i.content[id]
	return ok, nil
}

var _ Storage = (*InMemoryStorage)(nil)
