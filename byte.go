package orderedmap

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
// For example: if value := t.Get(key); value != nil { value found }
// O(logN)
func (m *Byte) Get(key byte) interface{} {
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
	return key.(byte), value
}

// Max returns the key-value to the maximum key, or nil if the tree is empty.
// For example: if key, value := t.Max(); key != nil { found }
// O(logN)
func (m *Byte) Max() (byte, interface{}) {
	key, value := m.m.Max()
	return key.(byte), value
}

// PopMin will delete the min node and return it.
// O(logN)
func (m *Byte) PopMin() (byte, interface{}) {
	key, value := m.m.PopMin()
	return key.(byte), value
}

// PopMax will delete the max node and return it.
// O(logN)
func (m *Byte) PopMax() (byte, interface{}) {
	key, value := m.m.PopMax()
	return key.(byte), value
}

// RangeAll traversals in ASC
// O(N)
func (m *Byte) RangeAll() []*ByteKeyValue {
	r := m.m.RangeAll()
	res := make([]*ByteKeyValue, len(r))
	for i, v := range r {
		res[i] = &ByteKeyValue{Key: v.First.(byte), Value: v.Second}
	}
	return res
}

// RangeAllDesc traversals in DESC
// O(N)
func (m *Byte) RangeAllDesc() []*ByteKeyValue {
	r := m.m.RangeAllDesc()
	res := make([]*ByteKeyValue, len(r))
	for i, v := range r {
		res[i] = &ByteKeyValue{Key: v.First.(byte), Value: v.Second}
	}
	return res
}

// Range traversals in [minKey, maxKey] in ASC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Byte) Range(minKey, maxKey byte) []*ByteKeyValue {
	r := m.m.Range(minKey, maxKey)
	res := make([]*ByteKeyValue, len(r))
	for i, v := range r {
		res[i] = &ByteKeyValue{Key: v.First.(byte), Value: v.Second}
	}
	return res
}

// RangeDesc traversals in [minKey, maxKey] in DESC
// MinKey & MaxKey are all closed interval.
// O(N)
func (m *Byte) RangeDesc(minKey, maxKey byte) []*ByteKeyValue {
	r := m.m.RangeDesc(minKey, maxKey)
	res := make([]*ByteKeyValue, len(r))
	for i, v := range r {
		res[i] = &ByteKeyValue{Key: v.First.(byte), Value: v.Second}
	}
	return res
}

func (m *Byte) Len() int {
	return m.m.Len()
}

func (m *Byte) IsEmpty() bool {
	return m.m.IsEmpty()
}
