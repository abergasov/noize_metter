package utils

import "sync"

type RWMap[K comparable, V any] struct {
	mu       sync.RWMutex
	Internal map[K]V
}

func NewRWMap[K comparable, V any]() *RWMap[K, V] {
	return &RWMap[K, V]{
		Internal: make(map[K]V),
	}
}

func NewRWMapFromStdMap[K comparable, V any](internal map[K]V) *RWMap[K, V] {
	return &RWMap[K, V]{
		Internal: internal,
	}
}

// Replace replaces the internal map with the provided one and returns the old one.
func (rm *RWMap[K, V]) Replace(internal map[K]V) map[K]V {
	rm.mu.Lock()
	current := rm.Internal
	rm.Internal = internal
	rm.mu.Unlock()
	return current
}

func (rm *RWMap[K, V]) Load(key K) (value V, ok bool) {
	rm.mu.RLock()
	result, ok := rm.Internal[key]
	rm.mu.RUnlock()
	return result, ok
}

func (rm *RWMap[K, V]) LoadAll() map[K]V {
	rm.mu.RLock()
	actual := make(map[K]V, len(rm.Internal))
	for key, value := range rm.Internal {
		actual[key] = value
	}
	rm.mu.RUnlock()
	return actual
}

func (rm *RWMap[K, V]) Delete(key K) {
	rm.mu.Lock()
	delete(rm.Internal, key)
	rm.mu.Unlock()
}

func (rm *RWMap[K, V]) DeleteAll() {
	rm.mu.Lock()
	for key := range rm.Internal {
		delete(rm.Internal, key)
	}
	rm.mu.Unlock()
}

func (rm *RWMap[K, V]) Store(key K, value V) {
	rm.mu.Lock()
	rm.Internal[key] = value
	rm.mu.Unlock()
}

func (rm *RWMap[K, V]) Keys() []K {
	rm.mu.RLock()
	keys := make([]K, 0, len(rm.Internal))
	for key := range rm.Internal {
		keys = append(keys, key)
	}
	rm.mu.RUnlock()
	return keys
}
