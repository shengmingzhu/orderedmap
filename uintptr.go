package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Uintptr struct {
	m OrderedMap
}

type UintptrKeyValue struct {
	Key   uintptr
	Value interface{}
}

func cmpUintptr(key1, key2 interface{}) int {
	if key1.(uintptr) == key2.(uintptr) {
		return 0
	} else if key1.(uintptr) > key2.(uintptr) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Uintptr) Get(key uintptr) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Uintptr) Put(key uintptr, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Uintptr) Delete(key uintptr) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Uintptr) Min() (uintptr, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(uintptr), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Uintptr) Max() (uintptr, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(uintptr), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Uintptr) PopMin() (uintptr, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(uintptr), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Uintptr) PopMax() (uintptr, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(uintptr), value
}

func (m *Uintptr) Keys() []uintptr {
	r := m.m.Keys()
	res := make([]uintptr, len(r))
	for i := range r {
		res[i] = r[i].(uintptr)
	}
	return res
}

func (m *Uintptr) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Uintptr) RangeAll() []UintptrKeyValue {
	r := m.m.RangeAll()
	return transformUintptr(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Uintptr) RangeAllDesc() []UintptrKeyValue {
	r := m.m.RangeAllDesc()
	return transformUintptr(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uintptr) Range(minKey, maxKey uintptr) []UintptrKeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformUintptr(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uintptr) RangeDesc(minKey, maxKey uintptr) []UintptrKeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformUintptr(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uintptr) RangeN(num int, key uintptr) []UintptrKeyValue {
	r := m.m.RangeN(num, key)
	return transformUintptr(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Uintptr) RangeDescN(num int, key uintptr) []UintptrKeyValue {
	r := m.m.RangeDescN(num, key)
	return transformUintptr(r)
}

func (m *Uintptr) Len() int {
	return m.m.Len()
}

func (m *Uintptr) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Uintptr) String() string {
	return m.m.String()
}

func transformUintptr(r []pair.Pair) []UintptrKeyValue {
	res := make([]UintptrKeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(uintptr)
		res[i].Value = r[i].Second
	}
	return res
}
