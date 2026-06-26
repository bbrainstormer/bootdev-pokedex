package lib

import (
	"sync"
)

type MutexMap[K comparable, V any] struct {
	internalMap map[K]V
	mut         *sync.Mutex
}

func NewMutexMap[K comparable, V any]() MutexMap[K, V] {
	return MutexMap[K, V]{
		internalMap: map[K]V{},
		mut:         &sync.Mutex{},
	}
}

func (m *MutexMap[K, V]) Has(key K) bool {
	m.mut.Lock()
	defer m.mut.Unlock()
	_, exists := m.internalMap[key]
	return exists
}

func (m *MutexMap[K, V]) Get(Key K) (V, bool) {
	m.mut.Lock()
	defer m.mut.Unlock()

	val, exists := m.internalMap[Key]
	return val, exists
}

func (m *MutexMap[K, V]) Set(key K, val V) {
	m.mut.Lock()
	defer m.mut.Unlock()

	if m.internalMap == nil {
		m.internalMap = map[K]V{}
	}

	m.internalMap[key] = val
}

func (m *MutexMap[K, V]) Iter(yield func(K, V) bool) {
	m.mut.Lock()
	defer m.mut.Unlock()
	for k, v := range m.internalMap {
		if !yield(k, v) {
			return
		}
	}
}
