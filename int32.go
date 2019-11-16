package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Int32 struct {
	m OrderedMap
}

type Int32KeyValue struct {
	Key   int32
	Value interface{}
}

func cmpInt32(key1, key2 interface{}) int {
	if key1.(int32) == key2.(int32) {
		return 0
	} else if key1.(int32) > key2.(int32) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Int32) Get(key int32) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Int32) Put(key int32, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Int32) Delete(key int32) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Int32) Min() (int32, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(int32), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Int32) Max() (int32, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(int32), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Int32) PopMin() (int32, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(int32), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Int32) PopMax() (int32, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(int32), value
}

func (m *Int32) Keys() []int32 {
	r := m.m.Keys()
	res := make([]int32, len(r))
	for i := range r {
		res[i] = r[i].(int32)
	}
	return res
}

func (m *Int32) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Int32) RangeAll() []Int32KeyValue {
	r := m.m.RangeAll()
	return transformInt32(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Int32) RangeAllDesc() []Int32KeyValue {
	r := m.m.RangeAllDesc()
	return transformInt32(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int32) Range(minKey, maxKey int32) []Int32KeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformInt32(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int32) RangeDesc(minKey, maxKey int32) []Int32KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformInt32(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int32) RangeN(num int, key int32) []Int32KeyValue {
	r := m.m.RangeN(num, key)
	return transformInt32(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int32) RangeDescN(num int, key int32) []Int32KeyValue {
	r := m.m.RangeDescN(num, key)
	return transformInt32(r)
}

func (m *Int32) Len() int {
	return m.m.Len()
}

func (m *Int32) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Int32) String() string {
	return m.m.String()
}

func transformInt32(r []pair.Pair) []Int32KeyValue {
	res := make([]Int32KeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(int32)
		res[i].Value = r[i].Second
	}
	return res
}
