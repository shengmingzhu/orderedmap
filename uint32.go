package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Uint32 struct {
	m OrderedMap
}

type Uint32KeyValue struct {
	Key   uint32
	Value interface{}
}

func cmpUint32(key1, key2 interface{}) int {
	if key1.(uint32) == key2.(uint32) {
		return 0
	} else if key1.(uint32) > key2.(uint32) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Uint32) Get(key uint32) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Uint32) Put(key uint32, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Uint32) Delete(key uint32) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Uint32) Min() (uint32, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(uint32), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Uint32) Max() (uint32, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(uint32), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Uint32) PopMin() (uint32, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(uint32), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Uint32) PopMax() (uint32, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(uint32), value
}

func (m *Uint32) Keys() []uint32 {
	r := m.m.Keys()
	res := make([]uint32, len(r))
	for i := range r {
		res[i] = r[i].(uint32)
	}
	return res
}

func (m *Uint32) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Uint32) RangeAll() []Uint32KeyValue {
	r := m.m.RangeAll()
	return transformUint32(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Uint32) RangeAllDesc() []Uint32KeyValue {
	r := m.m.RangeAllDesc()
	return transformUint32(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint32) Range(minKey, maxKey uint32) []Uint32KeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformUint32(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint32) RangeDesc(minKey, maxKey uint32) []Uint32KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformUint32(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uint32) RangeN(num int, key uint32) []Uint32KeyValue {
	r := m.m.RangeN(num, key)
	return transformUint32(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uint32) RangeDescN(num int, key uint32) []Uint32KeyValue {
	r := m.m.RangeDescN(num, key)
	return transformUint32(r)
}

func (m *Uint32) Len() int {
	return m.m.Len()
}

func (m *Uint32) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Uint32) String() string {
	return m.m.String()
}

func transformUint32(r []pair.Pair) []Uint32KeyValue {
	res := make([]Uint32KeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(uint32)
		res[i].Value = r[i].Second
	}
	return res
}
