package orderedmap

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
// For example: if value := t.Get(key); value != nil { value found }
// O(logN)
func (m *Int8) Get(key int8) interface{} {
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

// RangeAll traversals in ASC
// O(N)
func (m *Int8) RangeAll() []*Int8KeyValue {
	r := m.m.RangeAll()
	res := make([]*Int8KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int8KeyValue{Key: v.First.(int8), Value: v.Second}
	}
	return res
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Int8) RangeAllDesc() []*Int8KeyValue {
	r := m.m.RangeAllDesc()
	res := make([]*Int8KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int8KeyValue{Key: v.First.(int8), Value: v.Second}
	}
	return res
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int8) Range(minKey, maxKey int8) []*Int8KeyValue {
	r := m.m.Range(minKey, maxKey)
	res := make([]*Int8KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int8KeyValue{Key: v.First.(int8), Value: v.Second}
	}
	return res
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int8) RangeDesc(minKey, maxKey int8) []*Int8KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	res := make([]*Int8KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int8KeyValue{Key: v.First.(int8), Value: v.Second}
	}
	return res
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
