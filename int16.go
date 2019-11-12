package orderedmap

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
// For example: if value := t.Get(key); value != nil { value found }
// O(logN)
func (m *Int16) Get(key int16) interface{} {
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

// RangeAll traversals in ASC
// O(N)
func (m *Int16) RangeAll() []*Int16KeyValue {
	r := m.m.RangeAll()
	res := make([]*Int16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int16KeyValue{Key: v.First.(int16), Value: v.Second}
	}
	return res
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Int16) RangeAllDesc() []*Int16KeyValue {
	r := m.m.RangeAllDesc()
	res := make([]*Int16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int16KeyValue{Key: v.First.(int16), Value: v.Second}
	}
	return res
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int16) Range(minKey, maxKey int16) []*Int16KeyValue {
	r := m.m.Range(minKey, maxKey)
	res := make([]*Int16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int16KeyValue{Key: v.First.(int16), Value: v.Second}
	}
	return res
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Int16) RangeDesc(minKey, maxKey int16) []*Int16KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	res := make([]*Int16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int16KeyValue{Key: v.First.(int16), Value: v.Second}
	}
	return res
}

// RangeN get num key-values which >= key in ASC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int16) RangeN(num int, key int16) []*Int16KeyValue {
	r := m.m.RangeN(num, key)
	res := make([]*Int16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int16KeyValue{Key: v.First.(int16), Value: v.Second}
	}
	return res
}

// RangeDescN get num key-values which <= key in DESC
// Pair.First: Key, Pair.Second: Value
// O(N)
func (m *Int16) RangeDescN(num int, key int16) []*Int16KeyValue {
	r := m.m.RangeDescN(num, key)
	res := make([]*Int16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Int16KeyValue{Key: v.First.(int16), Value: v.Second}
	}
	return res
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
