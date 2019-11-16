package orderedmap

import (
	"github.com/shengmingzhu/datastructures/pair"
	"github.com/shengmingzhu/datastructures/rbtree"
)

type Any OrderedMap

type OrderedMap interface {
	Get(interface{}) (interface{}, bool) // O(logN)
	Put(interface{}, interface{})        // O(logN)
	Delete(interface{})                  // O(logN)

	Keys() []interface{}   // O(N)
	Values() []interface{} // O(N)

	Min() (interface{}, interface{})    // O(logN)
	Max() (interface{}, interface{})    // O(logN)
	PopMin() (interface{}, interface{}) // O(logN). Call Min() and Delete(), and return Min
	PopMax() (interface{}, interface{}) // O(logN). Call Max() and Delete(), and return Max

	RangeAll() []pair.Pair                            // O(N)
	RangeAllDesc() []pair.Pair                        // O(N)
	Range(minKey, maxKey interface{}) []pair.Pair     // O(logN) + O(K), K is the number of keys between minKey and maxKey.
	RangeDesc(minKey, maxKey interface{}) []pair.Pair // O(logN) + O(K), K is the number of keys between minKey and maxKey.
	RangeN(num int, key interface{}) []pair.Pair      // O(logN) + O(K)
	RangeDescN(num int, key interface{}) []pair.Pair  // O(logN) + O(K)

	Len() int      // O(1)
	IsEmpty() bool // O(1)

	// String is very useful when debugging
	// Example: fmt.Println(t) will print as follows:
	/*
	 *                               [ 6]B
	 *           [ 3]B                                   [13]R
	 * [ 2]R               [ 5]R               [ 8]B               [15]B
	 *                                              [10]R
	 */
	// Deprecated: only for debugging, unstable function
	String() string
}

func NewAny(cmp rbtree.CmpFunc) Any {
	return rbtree.New(cmp)
}

func NewByte() *Byte {
	return &Byte{m: rbtree.New(cmpByte)}
}

func NewInt() *Int {
	return &Int{m: rbtree.New(cmpInt)}
}

func NewInt8() *Int8 {
	return &Int8{m: rbtree.New(cmpInt8)}
}

func NewInt16() *Int16 {
	return &Int16{m: rbtree.New(cmpInt16)}
}

func NewInt32() *Int32 {
	return &Int32{m: rbtree.New(cmpInt32)}
}

func NewInt64() *Int64 {
	return &Int64{m: rbtree.New(cmpInt64)}
}

func NewRune() *Rune {
	return &Rune{m: rbtree.New(cmpRune)}
}

func NewString() *String {
	return &String{m: rbtree.New(cmpString)}
}

func NewUint() *Uint {
	return &Uint{m: rbtree.New(cmpUint)}
}

func NewUint8() *Uint8 {
	return &Uint8{m: rbtree.New(cmpUint8)}
}

func NewUint16() *Uint16 {
	return &Uint16{m: rbtree.New(cmpUint16)}
}

func NewUint32() *Uint32 {
	return &Uint32{m: rbtree.New(cmpUint32)}
}

func NewUint64() *Uint64 {
	return &Uint64{m: rbtree.New(cmpUint64)}
}

func NewUintptr() *Uintptr {
	return &Uintptr{m: rbtree.New(cmpUintptr)}
}
