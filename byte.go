package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Byte struct {
	m OrderedMap
}

type ByteKeyValue struct {
	Key   byte
	Value interface{}
}

func cmpByte(key1, key2 interface{}) int {
	if key1.(byte) == key2.(byte) {
		return 0
	} else if key1.(byte) > key2.(byte) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Byte) Get(key byte) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Byte) Put(key byte, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Byte) Delete(key byte) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Byte) Min() (byte, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(byte), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Byte) Max() (byte, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(byte), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Byte) PopMin() (byte, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(byte), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Byte) PopMax() (byte, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(byte), value
}

func (m *Byte) Keys() []byte {
	r := m.m.Keys()
	res := make([]byte, len(r))
	for i := range r {
		res[i] = r[i].(byte)
	}
	return res
}

func (m *Byte) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Byte) RangeAll() []ByteKeyValue {
	r := m.m.RangeAll()
	return transformByte(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Byte) RangeAllDesc() []ByteKeyValue {
	r := m.m.RangeAllDesc()
	return transformByte(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Byte) Range(minKey, maxKey byte) []ByteKeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformByte(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Byte) RangeN(num int, key byte) []ByteKeyValue {
	r := m.m.RangeN(num, key)
	return transformByte(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Byte) RangeDescN(num int, key byte) []ByteKeyValue {
	r := m.m.RangeDescN(num, key)
	return transformByte(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Byte) RangeDesc(minKey, maxKey byte) []ByteKeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformByte(r)
}

func transformByte(r []pair.Pair) []ByteKeyValue {
	res := make([]ByteKeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(byte)
		res[i].Value = r[i].Second
	}
	return res
}

func (m *Byte) Len() int {
	return m.m.Len()
}

func (m *Byte) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Byte) String() string {
	return m.m.String()
}
