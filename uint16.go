package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Uint16 struct {
	m OrderedMap
}

type Uint16KeyValue struct {
	Key   uint16
	Value interface{}
}

func cmpUint16(key1, key2 interface{}) int {
	if key1.(uint16) == key2.(uint16) {
		return 0
	} else if key1.(uint16) > key2.(uint16) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Uint16) Get(key uint16) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Uint16) Put(key uint16, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Uint16) Delete(key uint16) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Uint16) Min() (uint16, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(uint16), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Uint16) Max() (uint16, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(uint16), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Uint16) PopMin() (uint16, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(uint16), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Uint16) PopMax() (uint16, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(uint16), value
}

func (m *Uint16) Keys() []uint16 {
	r := m.m.Keys()
	res := make([]uint16, len(r))
	for i := range r {
		res[i] = r[i].(uint16)
	}
	return res
}

func (m *Uint16) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Uint16) RangeAll() []Uint16KeyValue {
	r := m.m.RangeAll()
	return transformUint16(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Uint16) RangeAllDesc() []Uint16KeyValue {
	r := m.m.RangeAllDesc()
	return transformUint16(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint16) Range(minKey, maxKey uint16) []Uint16KeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformUint16(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint16) RangeDesc(minKey, maxKey uint16) []Uint16KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformUint16(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uint16) RangeN(num int, key uint16) []Uint16KeyValue {
	r := m.m.RangeN(num, key)
	return transformUint16(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uint16) RangeDescN(num int, key uint16) []Uint16KeyValue {
	r := m.m.RangeDescN(num, key)
	return transformUint16(r)
}

func (m *Uint16) Len() int {
	return m.m.Len()
}

func (m *Uint16) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Uint16) String() string {
	return m.m.String()
}

func transformUint16(r []pair.Pair) []Uint16KeyValue {
	res := make([]Uint16KeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(uint16)
		res[i].Value = r[i].Second
	}
	return res
}
