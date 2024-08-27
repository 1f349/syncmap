package syncmap

import "sync"

// Map wraps sync.Map to provide generic type constraints.
type Map[K comparable, V any] struct {
	items sync.Map
}

func (m *Map[K, V]) Clear() {
	m.items.Clear()
}

func (m *Map[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.items.CompareAndDelete(key, old)
}

func (m *Map[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return m.items.CompareAndSwap(key, old, new)
}

func (m *Map[K, V]) Delete(key K) {
	m.items.Delete(key)
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.items.Load(key)
	return v.(V), ok
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.items.LoadAndDelete(key)
	return v.(V), loaded
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.items.LoadOrStore(key, value)
	return a.(V), loaded
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.items.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *Map[K, V]) Store(key K, value V) {
	m.items.Store(key, value)
}

func (m *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	p, loaded := m.items.Swap(key, value)
	return p.(V), loaded
}
