package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Int16 struct {
	m OrderedMap
}

type Int16KeyValue struct {
	Key   int16
	Value interface{}
}

func cmpInt16(key1, key2 interface{}) int {
	if key1.(int16) == key2.(int16) {
		return 0
	} else if key1.(int16) > key2.(int16) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Int16) Get(key int16) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Int16) Put(key int16, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Int16) Delete(key int16) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Int16) Min() (int16, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(int16), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Int16) Max() (int16, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(int16), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Int16) PopMin() (int16, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(int16), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Int16) PopMax() (int16, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(int16), value
}

func (m *Int16) Keys() []int16 {
	r := m.m.Keys()
	res := make([]int16, len(r))
	for i := range r {
		res[i] = r[i].(int16)
	}
	return res
}

func (m *Int16) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Int16) RangeAll() []Int16KeyValue {
	r := m.m.RangeAll()
	return transformInt16(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Int16) RangeAllDesc() []Int16KeyValue {
	r := m.m.RangeAllDesc()
	return transformInt16(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int16) Range(minKey, maxKey int16) []Int16KeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformInt16(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int16) RangeDesc(minKey, maxKey int16) []Int16KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformInt16(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int16) RangeN(num int, key int16) []Int16KeyValue {
	r := m.m.RangeN(num, key)
	return transformInt16(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int16) RangeDescN(num int, key int16) []Int16KeyValue {
	r := m.m.RangeDescN(num, key)
	return transformInt16(r)
}

func (m *Int16) Len() int {
	return m.m.Len()
}

func (m *Int16) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Int16) String() string {
	return m.m.String()
}

func transformInt16(r []pair.Pair) []Int16KeyValue {
	res := make([]Int16KeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(int16)
		res[i].Value = r[i].Second
	}
	return res
}
