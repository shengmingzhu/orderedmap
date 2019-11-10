package orderedmap

import (
	"strings"
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
// For example: if value := t.Get(key); value != nil { value found }
// O(logN)
func (m *String) Get(key string) interface{} {
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

// RangeAll traversals in ASC
// O(N)
func (m *String) RangeAll() []*StringKeyValue {
	r := m.m.RangeAll()
	res := make([]*StringKeyValue, len(r))
	for i, v := range r {
		res[i] = &StringKeyValue{Key: v.First.(string), Value: v.Second}
	}
	return res
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *String) RangeAllDesc() []*StringKeyValue {
	r := m.m.RangeAllDesc()
	res := make([]*StringKeyValue, len(r))
	for i, v := range r {
		res[i] = &StringKeyValue{Key: v.First.(string), Value: v.Second}
	}
	return res
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *String) Range(minKey, maxKey string) []*StringKeyValue {
	r := m.m.Range(minKey, maxKey)
	res := make([]*StringKeyValue, len(r))
	for i, v := range r {
		res[i] = &StringKeyValue{Key: v.First.(string), Value: v.Second}
	}
	return res
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *String) RangeDesc(minKey, maxKey string) []*StringKeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	res := make([]*StringKeyValue, len(r))
	for i, v := range r {
		res[i] = &StringKeyValue{Key: v.First.(string), Value: v.Second}
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
