package collections

import (
	"github.com/avivatedgi/go-rust-std/iterator"
	"github.com/avivatedgi/go-rust-std/option"
)

type Map[K comparable, V any] map[K]V

// Clears the map, removing all key-value pairs.
func (m *Map[K, V]) Clear() {
	for k := range *m {
		delete(*m, k)
	}
}

// Returns true if the map contains a value for the specified key.
func (m Map[K, V]) ContainsKey(key K) bool {
	_, ok := m[key]
	return ok
}

// Clears the map, returning all key-value pairs as an iterator.
func (m *Map[K, V]) Drain() iterator.Iterator[Pair[K, V]] {
	defer m.Clear()
	return m.Iter()
}

// Gets the given key’s corresponding entry in the map for in-place manipulation.
// WARNING: In difference from rust, this method does not return a reference to the value (!).
// But, it does match the signatures of MapEntry in rust.
func (m *Map[K, V]) Entry(key K) MapEntry[K, V] {
	return MapEntry[K, V]{parent: m, key: key}
}

// Returns the value corresponding to the key.
func (m Map[K, V]) Get(key K) option.Option[V] {
	if m.ContainsKey(key) {
		return option.Some(m[key])
	}

	return option.None[V]()
}

// Returns the key-value pair corresponding to the supplied key in a Pair (first is key, second is value).
func (m Map[K, V]) GetKeyValue(key K) option.Option[Pair[K, V]] {
	if m.ContainsKey(key) {
		return option.Some(Pair[K, V]{First: key, Second: m[key]})
	}

	return option.None[Pair[K, V]]()
}

// Inserts a key-value pair into the map.
// If the map did not have this key present, option.None is returned.
// If the map did have this key present, the value is updated, and the old value is returned.
// The key is not updated, though;
func (m *Map[K, V]) Insert(key K, value V) option.Option[V] {
	old := option.None[V]()

	if m.ContainsKey(key) {
		old = option.Some((*m)[key])
	}

	(*m)[key] = value
	return old
}

// Retreive all the keys of the map.
func (m Map[K, _]) Keys() []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Retreive all the values of the map
func (m Map[_, V]) Values() []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func (m Map[K, V]) KeyValuePairs() []Pair[K, V] {
	pairs := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair[K, V]{First: k, Second: v})
	}

	return pairs
}

// Return the map iterator.
func (m Map[K, V]) Iter() iterator.Iterator[Pair[K, V]] {
	return &sliceIterator[Pair[K, V]]{data: m.KeyValuePairs(), index: 0}
}

// Executes the f function once for each map entry.
// There is an option to stop the iteration in the middle, if the handler function returns false.
func (m *Map[K, V]) ForEach(f func(*K, *V) bool) {
	for k, v := range *m {
		if !f(&k, &v) {
			break
		}
	}
}

// This struct is constructed from the Entry method on Map.
type MapEntry[K comparable, V any] struct {
	parent *Map[K, V]
	key    K
}

// Returns this entry’s key.
func (m MapEntry[K, V]) Key() K {
	return m.key
}

// Ensures a value is in the entry by inserting the default if empty.
// Returns the entry’s value.
func (m MapEntry[K, V]) OrInsert(value V) V {
	if !m.parent.ContainsKey(m.key) {
		(*m.parent)[m.key] = value
	}

	return (*m.parent)[m.key]
}

// Ensures a value is in the entry by inserting the result of the default function if empty.
// Returns the entry’s value.
func (m MapEntry[K, V]) OrInsertWith(f func() V) V {
	if !m.parent.ContainsKey(m.key) {
		(*m.parent)[m.key] = f()
	}

	return (*m.parent)[m.key]
}

// Ensures a value is in the entry by inserting, if empty, the result of the default function.
func (m MapEntry[K, V]) OrInsertWithKey(f func(K) V) V {
	if !m.parent.ContainsKey(m.key) {
		(*m.parent)[m.key] = f(m.key)
	}

	return (*m.parent)[m.key]
}

// Ensures a value is in the entry by inserting the default value if empty.
// Returns the entry's value.
func (m MapEntry[K, V]) OrDefault() V {
	if !m.parent.ContainsKey(m.key) {
		var value V
		(*m.parent)[m.key] = value
	}

	return (*m.parent)[m.key]
}

// Provides access to an occupied entry before any potential inserts into the map.
func (m MapEntry[K, V]) AndModify(f func(*V)) MapEntry[K, V] {
	if m.parent.ContainsKey(m.key) {
		value := (*m.parent)[m.key]
		f(&value)
		(*m.parent)[m.key] = value
	}

	return m
}
