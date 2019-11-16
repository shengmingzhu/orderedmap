package orderedmap

import "github.com/shengmingzhu/datastructures/pair"

type Rune struct {
	m OrderedMap
}

type RuneKeyValue struct {
	Key   rune
	Value interface{}
}

func cmpRune(key1, key2 interface{}) int {
	if key1.(rune) == key2.(rune) {
		return 0
	} else if key1.(rune) > key2.(rune) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value, ok := t.Get(key); ok { value found }
// O(logN)
func (m *Rune) Get(key rune) (interface{}, bool) {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Rune) Put(key rune, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Rune) Delete(key rune) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Rune) Min() (rune, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(rune), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Rune) Max() (rune, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(rune), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Rune) PopMin() (rune, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(rune), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Rune) PopMax() (rune, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(rune), value
}

func (m *Rune) Keys() []rune {
	r := m.m.Keys()
	res := make([]rune, len(r))
	for i := range r {
		res[i] = r[i].(rune)
	}
	return res
}

func (m *Rune) Values() []interface{} {
	return m.m.Values()
}

// RangeAll traversals in ASC
// O(N)
func (m *Rune) RangeAll() []RuneKeyValue {
	r := m.m.RangeAll()
	return transformRune(r)
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Rune) RangeAllDesc() []RuneKeyValue {
	r := m.m.RangeAllDesc()
	return transformRune(r)
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Rune) Range(minKey, maxKey rune) []RuneKeyValue {
	r := m.m.Range(minKey, maxKey)
	return transformRune(r)
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Rune) RangeDesc(minKey, maxKey rune) []RuneKeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	return transformRune(r)
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Rune) RangeN(num int, key rune) []RuneKeyValue {
	r := m.m.RangeN(num, key)
	return transformRune(r)
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Rune) RangeDescN(num int, key rune) []RuneKeyValue {
	r := m.m.RangeDescN(num, key)
	return transformRune(r)
}

func (m *Rune) Len() int {
	return m.m.Len()
}

func (m *Rune) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Rune) String() string {
	return m.m.String()
}

func transformRune(r []pair.Pair) []RuneKeyValue {
	res := make([]RuneKeyValue, len(r))
	for i := range r {
		res[i].Key = r[i].First.(rune)
		res[i].Value = r[i].Second
	}
	return res
}
