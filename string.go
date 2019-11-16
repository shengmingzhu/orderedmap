package orderedmap

import (
	"github.com/shengmingzhu/datastructures/pair"
	"strings"
	"unsafe"
)

type String struct {
	m OrderedMap
}

type StringKeyValue struct {
	Key   string
	Value interface{}
}

func cmpString(key1, key2 interface{}) int {
	return strings.Compare(key1.(string), key2.(string))
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *String) Get(key string) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *String) Put(key string, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *String) Delete(key string) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *String) Min() (string, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return "", value
	}
	return key.(string), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *String) Max() (string, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return "", value
	}
	return key.(string), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *String) PopMin() (string, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return "", value
	}
	return key.(string), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *String) PopMax() (string, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return "", value
	}
	return key.(string), value
}

func (m *String) Keys() []string {
	r := m.m.Keys()
	res := *(*[]string)(unsafe.Pointer(&r))
	for i := range r {
		res[i] = r[i].(string)
	}
	return res
}

func (m *String) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *String) RangeAll() []StringKeyValue {
	r := m.m.RangeAll()
	return transformString(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *String) RangeAllDesc() []StringKeyValue {
	r := m.m.RangeAllDesc()
	return transformString(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *String) Range(minKey, maxKey string) []StringKeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformString(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *String) RangeDesc(minKey, maxKey string) []StringKeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformString(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *String) RangeN(num int, key string) []StringKeyValue {
	r := m.m.RangeN(num, key)
	return transformString(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *String) RangeDescN(num int, key string) []StringKeyValue {
	r := m.m.RangeDescN(num, key)
	return transformString(r)
}

func transformString(r []pair.Pair) []StringKeyValue {
	res := *(*[]StringKeyValue)(unsafe.Pointer(&r))
	for i := range r {
		res[i].Key = r[i].First.(string)
	}
	return res
}

func (m *String) Len() int {
	return m.m.Len()
}

func (m *String) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *String) String() string {
	return m.m.String()
}
