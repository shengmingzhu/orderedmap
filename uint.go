package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Uint struct {
	m OrderedMap
}

type UintKeyValue struct {
	Key   uint
	Value interface{}
}

func cmpUint(key1, key2 interface{}) int {
	if key1.(uint) == key2.(uint) {
		return 0
	} else if key1.(uint) > key2.(uint) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Uint) Get(key uint) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Uint) Put(key uint, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Uint) Delete(key uint) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Uint) Min() (uint, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(uint), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Uint) Max() (uint, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(uint), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Uint) PopMin() (uint, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(uint), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Uint) PopMax() (uint, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(uint), value
}

func (m *Uint) Keys() []uint {
	r := m.m.Keys()
	res := make([]uint, len(r))
	for i := range r {
		res[i] = r[i].(uint)
	}
	return res
}

func (m *Uint) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Uint) RangeAll() []UintKeyValue {
	r := m.m.RangeAll()
	return transformUint(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Uint) RangeAllDesc() []UintKeyValue {
	r := m.m.RangeAllDesc()
	return transformUint(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint) Range(minKey, maxKey uint) []UintKeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformUint(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint) RangeDesc(minKey, maxKey uint) []UintKeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformUint(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uint) RangeN(num int, key uint) []UintKeyValue {
	r := m.m.RangeN(num, key)
	return transformUint(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uint) RangeDescN(num int, key uint) []UintKeyValue {
	r := m.m.RangeDescN(num, key)
	return transformUint(r)
}

func (m *Uint) Len() int {
	return m.m.Len()
}

func (m *Uint) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Uint) String() string {
	return m.m.String()
}

func transformUint(r []pair.Pair) []UintKeyValue {
	res := make([]UintKeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(uint)
		res[i].Value = r[i].Second
	}
	return res
}
