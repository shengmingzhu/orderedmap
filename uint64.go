package orderedmap

import (
	"github.com/shengmingzhu/datastructures/pair"
)

type Uint64 struct {
	m OrderedMap
}

type Uint64KeyValue struct {
	Key   uint64
	Value interface{}
}

func cmpUint64(key1, key2 interface{}) int {
	if key1.(uint64) == key2.(uint64) {
		return 0
	} else if key1.(uint64) > key2.(uint64) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Uint64) Get(key uint64) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Uint64) Put(key uint64, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Uint64) Delete(key uint64) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Uint64) Min() (uint64, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(uint64), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Uint64) Max() (uint64, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(uint64), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Uint64) PopMin() (uint64, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(uint64), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Uint64) PopMax() (uint64, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(uint64), value
}

func (m *Uint64) Keys() []uint64 {
	r := m.m.Keys()
	res := make([]uint64, len(r))
	for i := range r {
		res[i] = r[i].(uint64)
	}
	return res
}

func (m *Uint64) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Uint64) RangeAll() []Uint64KeyValue {
	r := m.m.RangeAll()
	return transformUint64(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Uint64) RangeAllDesc() []Uint64KeyValue {
	r := m.m.RangeAllDesc()
	return transformUint64(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint64) Range(minKey, maxKey uint64) []Uint64KeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformUint64(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint64) RangeDesc(minKey, maxKey uint64) []Uint64KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformUint64(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uint64) RangeN(num int, key uint64) []Uint64KeyValue {
	r := m.m.RangeN(num, key)
	return transformUint64(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uint64) RangeDescN(num int, key uint64) []Uint64KeyValue {
	r := m.m.RangeDescN(num, key)
	return transformUint64(r)
}

func (m *Uint64) Len() int {
	return m.m.Len()
}

func (m *Uint64) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Uint64) String() string {
	return m.m.String()
}

func transformUint64(r []pair.Pair) []Uint64KeyValue {
	res := make([]Uint64KeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(uint64)
		res[i].Value = r[i].Second
	}
	return res
}
