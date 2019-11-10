package orderedmap

type Uint16 struct {
	m OrderedMap
}

type Uint16KeyValue struct {
	Key   uint16
	Value interface{}
}

func cmpUint16(key1, key2 interface{}) int {
	if key1.(uint16) == key2.(uint16) {
		return 0
	} else if key1.(uint16) > key2.(uint16) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value := t.Get(key); value != nil { value found }
// O(logN)
func (m *Uint16) Get(key uint16) interface{} {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Uint16) Put(key uint16, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Uint16) Delete(key uint16) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Uint16) Min() (uint16, interface{}) {
	key, value := m.m.Min()
	if key == nil {
		return 0, value
	}
	return key.(uint16), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Uint16) Max() (uint16, interface{}) {
	key, value := m.m.Max()
	if key == nil {
		return 0, value
	}
	return key.(uint16), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Uint16) PopMin() (uint16, interface{}) {
	key, value := m.m.PopMin()
	if key == nil {
		return 0, value
	}
	return key.(uint16), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Uint16) PopMax() (uint16, interface{}) {
	key, value := m.m.PopMax()
	if key == nil {
		return 0, value
	}
	return key.(uint16), value
}

// RangeAll traversals in ASC
// O(N)
func (m *Uint16) RangeAll() []*Uint16KeyValue {
	r := m.m.RangeAll()
	res := make([]*Uint16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint16KeyValue{Key: v.First.(uint16), Value: v.Second}
	}
	return res
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Uint16) RangeAllDesc() []*Uint16KeyValue {
	r := m.m.RangeAllDesc()
	res := make([]*Uint16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint16KeyValue{Key: v.First.(uint16), Value: v.Second}
	}
	return res
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint16) Range(minKey, maxKey uint16) []*Uint16KeyValue {
	r := m.m.Range(minKey, maxKey)
	res := make([]*Uint16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint16KeyValue{Key: v.First.(uint16), Value: v.Second}
	}
	return res
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint16) RangeDesc(minKey, maxKey uint16) []*Uint16KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	res := make([]*Uint16KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint16KeyValue{Key: v.First.(uint16), Value: v.Second}
	}
	return res
}

func (m *Uint16) Len() int {
	return m.m.Len()
}

func (m *Uint16) IsEmpty() bool {
	return m.m.IsEmpty()
}

// Deprecated: only for debugging, unstable function
func (m *Uint16) String() string {
	return m.m.String()
}
