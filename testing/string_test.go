package orderedmap_test

import (
	"github.com/shengmingzhu/orderedmap"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

const (
	testCountString int = 1 << 4
	rangeLenString      = 1 << 10
)

func TestNewString(t *testing.T) {
	Convey("NewString and Put and Len", t, func() {
		m := orderedmap.NewString()
		hm := make(map[int]struct{}, testCountString)
		sl := make([]string, 0, testCountString)
		rand.Seed(time.Now().UnixNano())
		for i := int(0); i < testCountString; {
			key := int(rand.Int31n(0x40000000)) & (0x40000000 - 1)
			if _, ok := hm[key]; !ok {
				hm[key] = struct{}{}
				m.Put(strconv.Itoa(key), key<<1)
				sl = append(sl, strconv.Itoa(key))
				i++
			}
		}
		So(m.Len(), ShouldEqual, len(hm))
		sort.Slice(sl, func(i, j int) bool {
			return sl[i] < sl[j]
		})

		Convey("Get", func() {
			for k := range hm {
				So(m.Get(strconv.Itoa(k)), ShouldEqual, k<<1)
			}
		})

		Convey("Min", func() {
			key, _ := m.Min()
			So(key, ShouldEqual, sl[0])
		})

		Convey("Max", func() {
			key, _ := m.Max()
			So(key, ShouldEqual, sl[testCountString-1])
		})

		Convey("RangeAll", func() {
			pairs := m.RangeAll()
			So(len(pairs), ShouldEqual, len(sl))
			for i := range pairs {
				So(pairs[i].Key, ShouldEqual, sl[i])
				So(strconv.Itoa(pairs[i].Value.(int)>>1), ShouldEqual, pairs[i].Key)
			}
		})

		Convey("RangeAllDesc", func() {
			pairs := m.RangeAllDesc()
			So(len(pairs), ShouldEqual, len(sl))
			for i := range pairs {
				So(pairs[i].Key, ShouldEqual, sl[int(testCountString)-i-1])
				So(strconv.Itoa(pairs[i].Value.(int)>>1), ShouldEqual, pairs[i].Key)
			}
		})

		Convey("Range", func() {
			rand.Seed(time.Now().UnixNano())
			iKey1 := int(rand.Int63n(int64(testCountString)>>1)) & int((int64(testCountString)>>1)-1)
			iKey2 := int(rand.Int63n(int64(testCountString)>>1)) & int((int64(testCountString)>>1)-1)
			if iKey2 < 1 {
				iKey2 = 1
			}
			pairs := m.Range(sl[iKey1], sl[iKey1+iKey2])
			So(len(pairs), ShouldEqual, iKey2+1)
			for i := range pairs {
				So(pairs[i].Key, ShouldEqual, sl[int(iKey1)+i])
				So(strconv.Itoa(pairs[i].Value.(int)>>1), ShouldEqual, pairs[i].Key)
			}
		})

		Convey("RangeDesc", func() {
			rand.Seed(time.Now().UnixNano())
			iKey1 := int(rand.Int63n(int64(testCountString)>>1)) & int((int64(testCountString)>>1)-1)
			iKey2 := int(rand.Int63n(int64(testCountString)>>1)) & int((int64(testCountString)>>1)-1)
			if iKey2 < 1 {
				iKey2 = 1
			}
			pairs := m.RangeDesc(sl[iKey1], sl[iKey1+iKey2])
			So(len(pairs), ShouldEqual, iKey2+1)
			for i := range pairs {
				So(pairs[i].Key, ShouldEqual, sl[int(iKey1+iKey2)-i])
				So(strconv.Itoa(pairs[i].Value.(int)>>1), ShouldEqual, pairs[i].Key)
			}
		})

		Convey("RangeN", func() {
			rand.Seed(time.Now().UnixNano())
			iKey1 := 4
			iKey2 := 3
			pairs := m.RangeN(int(iKey2), sl[iKey1])
			So(len(pairs), ShouldEqual, iKey2)
			for i := range pairs {
				So(pairs[i].Key, ShouldEqual, sl[int(iKey1)+i])
			}
		})

		Convey("RangeDescN", func() {
			rand.Seed(time.Now().UnixNano())
			iKey1 := 4
			iKey2 := 3
			pairs := m.RangeDescN(int(iKey2), sl[iKey1])
			So(len(pairs), ShouldEqual, iKey2)
			for i := range pairs {
				So(pairs[i].Key, ShouldEqual, sl[int(iKey1)-i])
			}
		})

		Convey("String", func() {
			str := m.String()
			//fmt.Println()
			//fmt.Println(str)
			So(len(str), ShouldBeGreaterThan, 0)
		})

		Convey("PopMin", func() {
			for i := 0; i < 4; i++ {
				k, v := m.PopMin()
				So(k, ShouldEqual, sl[i])
				So(strconv.Itoa(v.(int)>>1), ShouldEqual, k)
				delete(hm, v.(int)>>1)
			}
		})

		Convey("PopMax", func() {
			for i := 0; i < 4; i++ {
				k, v := m.PopMax()
				So(k, ShouldEqual, sl[len(sl)-1-i])
				So(strconv.Itoa(v.(int)>>1), ShouldEqual, k)
				delete(hm, v.(int)>>1)
			}
		})

		Convey("Delete", func() {
			count := len(hm)
			for key := range hm {
				//fmt.Println()
				//fmt.Println(m)
				m.Delete(strconv.Itoa(key))
				count--
				So(m.Len(), ShouldEqual, count)
			}
		})

		Convey("EmptyMap", func() {
			m := orderedmap.NewString()
			v := m.Get("1")
			So(v, ShouldEqual, nil)
			m.Delete("2")
			k, v := m.Min()
			So(k, ShouldEqual, "")
			So(v, ShouldEqual, nil)
			k, v = m.Max()
			So(k, ShouldEqual, "")
			So(v, ShouldEqual, nil)
			k, v = m.PopMin()
			So(k, ShouldEqual, "")
			So(v, ShouldEqual, nil)
			k, v = m.PopMax()
			So(k, ShouldEqual, "")
			So(v, ShouldEqual, nil)
			res := m.RangeAll()
			So(len(res), ShouldEqual, 0)
			res = m.RangeAllDesc()
			So(len(res), ShouldEqual, 0)
			res = m.Range("1", "6")
			So(len(res), ShouldEqual, 0)
			res = m.RangeDesc("1", "10")
			So(len(res), ShouldEqual, 0)
		})
	})
}

func BenchmarkString_Put(b *testing.B) {
	b.StopTimer()
	m := orderedmap.NewString()
	sl := make([]string, 0, b.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		sl = append(sl, key)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Put(sl[i], struct{}{})
	}
}

func BenchmarkHashMapString_Put(b *testing.B) {
	b.StopTimer()
	m := make(map[string]struct{})
	sl := make([]string, 0, b.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		sl = append(sl, key)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m[sl[i]] = struct{}{}
	}
}

func BenchmarkString_Get(b *testing.B) {
	b.StopTimer()
	m := orderedmap.NewString()
	sl := make([]string, 0, b.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		sl = append(sl, key)
		m.Put(key, struct{}{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Get(sl[i])
	}
}

func BenchmarkHashMapString_Get(b *testing.B) {
	b.StopTimer()
	m := make(map[string]struct{})
	sl := make([]string, 0, b.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		sl = append(sl, key)
		m[key] = struct{}{}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = m[sl[i]]
	}
}

func BenchmarkString_Delete(b *testing.B) {
	b.StopTimer()
	m := orderedmap.NewString()
	sl := make([]string, 0, b.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		sl = append(sl, key)
		m.Put(key, struct{}{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(sl[i])
	}
}

func BenchmarkHashMapString_Delete(b *testing.B) {
	b.StopTimer()
	m := make(map[string]struct{})
	sl := make([]string, 0, b.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		sl = append(sl, key)
		m[key] = struct{}{}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		delete(m, sl[i])
	}
}

func BenchmarkString_RangeAll(b *testing.B) {
	b.StopTimer()
	m := orderedmap.NewString()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rangeLenString; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		m.Put(key, struct{}{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = m.RangeAll()
	}
}

func BenchmarkHashMapString_RangeAllNoSort(b *testing.B) {
	b.StopTimer()
	m := make(map[string]struct{})
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rangeLenString; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		m[key] = struct{}{}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]*orderedmap.StringKeyValue, 0)
		for key, value := range m {
			pair := &orderedmap.StringKeyValue{Key: key, Value: value}
			arr = append(arr, pair)
		}
	}
}

func BenchmarkHashMapString_RangeAllAndSort(b *testing.B) {
	b.StopTimer()
	m := make(map[string]struct{})
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rangeLenString; i++ {
		key := string(uint32(rand.Int31n(0x40000000)) & (0x40000000 - 1))
		m[key] = struct{}{}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		arr := make([]*orderedmap.StringKeyValue, 0)
		for key, value := range m {
			pair := &orderedmap.StringKeyValue{Key: key, Value: value}
			arr = append(arr, pair)
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Key < arr[j].Key
		})
	}
}
