package orderedmap

type Uint64 struct {
	m OrderedMap
}

type Uint64KeyValue struct {
	Key   uint64
	Value interface{}
}

func cmpUint64(key1, key2 interface{}) int {
	if key1.(uint64) == key2.(uint64) {
		return 0
	} else if key1.(uint64) > key2.(uint64) {
		return 1
	} else {
		return -1
	}
}

// Get returns the value to key, or nil if not found.
// For example: if value := t.Get(key); value != nil { value found }
// O(logN)
func (m *Uint64) Get(key uint64) interface{} {
	return m.m.Get(key)
}

// Put stores the key-value pair into RbTree.
// 1. If there is already a same key in RbTree, it will replace the value.
// 2. Otherwise, it will insert a new node with the key-value.
// O(logN)
func (m *Uint64) Put(key uint64, value interface{}) {
	m.m.Put(key, value)
}

// O(logN)
func (m *Uint64) Delete(key uint64) {
	m.m.Delete(key)
}

// Min returns the key-value to the minimum key, or nil if the tree is empty.
// For example: if key, value := t.Min(key); key != nil { found }
// O(logN)
func (m *Uint64) Min() (uint64, interface{}) {
	key, value := m.m.Min()
	return key.(uint64), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Uint64) Max() (uint64, interface{}) {
	key, value := m.m.Max()
	return key.(uint64), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Uint64) PopMin() (uint64, interface{}) {
	key, value := m.m.PopMin()
	return key.(uint64), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Uint64) PopMax() (uint64, interface{}) {
	key, value := m.m.PopMax()
	return key.(uint64), value
}

// RangeAll traversals in ASC
// O(N)
func (m *Uint64) RangeAll() []*Uint64KeyValue {
	r := m.m.RangeAll()
	res := make([]*Uint64KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint64KeyValue{Key: v.First.(uint64), Value: v.Second}
	}
	return res
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Uint64) RangeAllDesc() []*Uint64KeyValue {
	r := m.m.RangeAllDesc()
	res := make([]*Uint64KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint64KeyValue{Key: v.First.(uint64), Value: v.Second}
	}
	return res
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint64) Range(minKey, maxKey uint64) []*Uint64KeyValue {
	r := m.m.Range(minKey, maxKey)
	res := make([]*Uint64KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint64KeyValue{Key: v.First.(uint64), Value: v.Second}
	}
	return res
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Uint64) RangeDesc(minKey, maxKey uint64) []*Uint64KeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	res := make([]*Uint64KeyValue, len(r))
	for i, v := range r {
		res[i] = &Uint64KeyValue{Key: v.First.(uint64), Value: v.Second}
	}
	return res
}

func (m *Uint64) Len() int {
	return m.m.Len()
}

func (m *Uint64) IsEmpty() bool {
	return m.m.IsEmpty()
}
