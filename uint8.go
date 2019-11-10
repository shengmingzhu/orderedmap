package orderedmap

type Uint8 struct {
	m OrderedMap
}

type Uint8KeyValue struct {
	Key   uint8
	Value interface{}
}

func cmpUint8(key1, key2 interface{}) int {
	if key1.(uint8) == key2.(uint8) {
		return 0
	} else if key1.(uint8) > key2.(uint8) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value := t.Get(key); value != nil { value found }
// O(logN)
func (m *Uint8) Get(key uint8) interface{} {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Uint8) Put(key uint8, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Uint8) Delete(key uint8) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Uint8) Min() (uint8, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(uint8), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Uint8) Max() (uint8, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(uint8), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Uint8) PopMin() (uint8, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(uint8), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Uint8) PopMax() (uint8, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(uint8), value
}

// RangeAll traversals in ASC
// O(N)
func (m *Uint8) RangeAll() []*Uint8KeyValue {
	r := m.m.RangeAll()
	res := make([]*Uint8KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint8KeyValue{Key: v.First.(uint8), Value: v.Second}
	}
	return res
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Uint8) RangeAllDesc() []*Uint8KeyValue {
	r := m.m.RangeAllDesc()
	res := make([]*Uint8KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint8KeyValue{Key: v.First.(uint8), Value: v.Second}
	}
	return res
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint8) Range(minKey, maxKey uint8) []*Uint8KeyValue {
	r := m.m.Range(minKey, maxKey)
	res := make([]*Uint8KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint8KeyValue{Key: v.First.(uint8), Value: v.Second}
	}
	return res
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint8) RangeDesc(minKey, maxKey uint8) []*Uint8KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	res := make([]*Uint8KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint8KeyValue{Key: v.First.(uint8), Value: v.Second}
	}
	return res
}

func (m *Uint8) Len() int {
	return m.m.Len()
}

func (m *Uint8) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Uint8) String() string {
	return m.m.String()
}
