package safemap

import (
	"fmt"
	"sync"
)

type Safemap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func New[K comparable, V any]() *Safemap[K, V] {
	return &Safemap[K, V]{
		data: make(map[K]V),
	}
}

func (m *Safemap[K, V]) Insert(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}

func (m *Safemap[K, V]) Get(key K) (V, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.data[key]
	if !ok {
		return value, fmt.Errorf("Key %v not found", key)
	}

	return value, nil
}

func (m *Safemap[K, V]) Update(key K, value V) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; !ok {
		return fmt.Errorf("Key %v not found", key)
	}
	m.data[key] = value
	return nil
}

func (m *Safemap[K, V]) Delete(key K) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; !ok {
		return fmt.Errorf("Key %v not found", key)
	}

	delete(m.data, key)
	return nil

}

func (m *Safemap[K, V]) HasKey(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.data[key]

	return ok
}
