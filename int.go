package orderedmap

type Int struct {
	m OrderedMap
}

type IntKeyValue struct {
	Key   int
	Value interface{}
}

func cmpInt(key1, key2 interface{}) int {
	return key1.(int) - key2.(int)
}

// Get returns the value to key, or nil if not found.
// For example: if value := t.Get(key); value != nil { value found }
// O(logN)
func (m *Int) Get(key int) interface{} {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Int) Put(key int, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Int) Delete(key int) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Int) Min() (int, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(int), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Int) Max() (int, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(int), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Int) PopMin() (int, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(int), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Int) PopMax() (int, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(int), value
}

// RangeAll traversals in ASC
// O(N)
func (m *Int) RangeAll() []*IntKeyValue {
	r := m.m.RangeAll()
	res := make([]*IntKeyValue, len(r))
	for i, v := range r {
		res[i] = &IntKeyValue{Key: v.First.(int), Value: v.Second}
	}
	return res
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Int) RangeAllDesc() []*IntKeyValue {
	r := m.m.RangeAllDesc()
	res := make([]*IntKeyValue, len(r))
	for i, v := range r {
		res[i] = &IntKeyValue{Key: v.First.(int), Value: v.Second}
	}
	return res
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int) Range(minKey, maxKey int) []*IntKeyValue {
	r := m.m.Range(minKey, maxKey)
	res := make([]*IntKeyValue, len(r))
	for i, v := range r {
		res[i] = &IntKeyValue{Key: v.First.(int), Value: v.Second}
	}
	return res
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int) RangeDesc(minKey, maxKey int) []*IntKeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	res := make([]*IntKeyValue, len(r))
	for i, v := range r {
		res[i] = &IntKeyValue{Key: v.First.(int), Value: v.Second}
	}
	return res
}

func (m *Int) Len() int {
	return m.m.Len()
}

func (m *Int) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Int) String() string {
	return m.m.String()
}
