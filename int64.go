package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Int64 struct {
	m OrderedMap
}

type Int64KeyValue struct {
	Key   int64
	Value interface{}
}

func cmpInt64(key1, key2 interface{}) int {
	if key1.(int64) == key2.(int64) {
		return 0
	} else if key1.(int64) > key2.(int64) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Int64) Get(key int64) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Int64) Put(key int64, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Int64) Delete(key int64) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Int64) Min() (int64, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(int64), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Int64) Max() (int64, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(int64), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Int64) PopMin() (int64, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(int64), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Int64) PopMax() (int64, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(int64), value
}

func (m *Int64) Keys() []int64 {
	r := m.m.Keys()
	res := make([]int64, len(r))
	for i := range r {
		res[i] = r[i].(int64)
	}
	return res
}

func (m *Int64) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Int64) RangeAll() []Int64KeyValue {
	r := m.m.RangeAll()
	return transformInt64(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Int64) RangeAllDesc() []Int64KeyValue {
	r := m.m.RangeAllDesc()
	return transformInt64(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int64) Range(minKey, maxKey int64) []Int64KeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformInt64(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int64) RangeDesc(minKey, maxKey int64) []Int64KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformInt64(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int64) RangeN(num int, key int64) []Int64KeyValue {
	r := m.m.RangeN(num, key)
	return transformInt64(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int64) RangeDescN(num int, key int64) []Int64KeyValue {
	r := m.m.RangeDescN(num, key)
	return transformInt64(r)
}

func (m *Int64) Len() int {
	return m.m.Len()
}

func (m *Int64) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Int64) String() string {
	return m.m.String()
}

func transformInt64(r []pair.Pair) []Int64KeyValue {
	res := make([]Int64KeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(int64)
		res[i].Value = r[i].Second
	}
	return res
}
