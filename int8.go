package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Int8 struct {
	m OrderedMap
}

type Int8KeyValue struct {
	Key   int8
	Value interface{}
}

func cmpInt8(key1, key2 interface{}) int {
	if key1.(int8) == key2.(int8) {
		return 0
	} else if key1.(int8) > key2.(int8) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Int8) Get(key int8) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Int8) Put(key int8, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Int8) Delete(key int8) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Int8) Min() (int8, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(int8), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Int8) Max() (int8, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(int8), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Int8) PopMin() (int8, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(int8), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Int8) PopMax() (int8, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(int8), value
}

func (m *Int8) Keys() []int8 {
	r := m.m.Keys()
	res := make([]int8, len(r))
	for i := range r {
		res[i] = r[i].(int8)
	}
	return res
}

func (m *Int8) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Int8) RangeAll() []Int8KeyValue {
	r := m.m.RangeAll()
	return transformInt8(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Int8) RangeAllDesc() []Int8KeyValue {
	r := m.m.RangeAllDesc()
	return transformInt8(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int8) Range(minKey, maxKey int8) []Int8KeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformInt8(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int8) RangeDesc(minKey, maxKey int8) []Int8KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformInt8(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int8) RangeN(num int, key int8) []Int8KeyValue {
	r := m.m.RangeN(num, key)
	return transformInt8(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int8) RangeDescN(num int, key int8) []Int8KeyValue {
	r := m.m.RangeDescN(num, key)
	return transformInt8(r)
}

func (m *Int8) Len() int {
	return m.m.Len()
}

func (m *Int8) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Int8) String() string {
	return m.m.String()
}

func transformInt8(r []pair.Pair) []Int8KeyValue {
	res := make([]Int8KeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(int8)
		res[i].Value = r[i].Second
	}
	return res
}
