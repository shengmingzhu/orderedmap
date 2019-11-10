# orderedmap
An efficient ordered map, the basic data structure is RB-Tree coding by Golang.

# Example
```
	m := orderedmap.NewInt()
	fmt.Println(m.IsEmpty()) // true
	m.Put(1, 11)
	m.Put(2, 22)
	m.Put(3, 33)
	m.Put(4, 44)
	m.Put(5, 55)
	m.Delete(2) // remove {2,22}
	fmt.Println(m.Len()) //  4
	if v := m.Get(1); v != nil {
		fmt.Println(v.(int)) // 11
	}
	k, v := m.Min()
	fmt.Println(k, v.(int)) // 1 11
	k, v = m.Max()
	fmt.Println(k, v.(int)) // 5 55
	_ = m.RangeAll()        // [{1,11},{3,33},{4,44},{5,55}]
	_ = m.RangeAllDesc()    // [{5,55},{4,44},{3,33},{1,11}]
	_ = m.Range(3, 4)       // [{3,33},{4,44}]
	_ = m.RangeDesc(3, 4)   // [{4,44},{3,33}]
	_, _ = m.PopMin()       // remove and return {1,11}
	_, _ = m.PopMax()       // remove and return {5,55}
	fmt.Println(m.Len())    // 2
```

# Testing
The complete function and performance test cases are placed in the testing folder in the repository.
All functional and performance tests can be performed using the following instructions:
```
$ cd $GOPATH/src/github.com/shengmingzhu/orderedmap/testing
$ go test
$ go test -bench=. -timeout=20m
```

Due to the large number of test cases, performance testing will take a long time, so it is recommended to call only for one type, such as:
```
$ go test -bench=. int8_test.go
........................................................................................................................................................
152 total assertions

goos: windows
goarch: amd64
BenchmarkInt8_Put-4                             20000000                93.4 ns/op
BenchmarkHashMapInt8_Put-4                      30000000                40.1 ns/op

BenchmarkInt8_Get-4                             20000000                69.8 ns/op
BenchmarkHashMapInt8_Get-4                      30000000                40.3 ns/op

BenchmarkInt8_Delete-4                          300000000                6.28 ns/op
BenchmarkHashMapInt8_Delete-4                   500000000                3.10 ns/op

BenchmarkInt8_RangeAll-4                          200000             11094 ns/op
BenchmarkHashMapInt8_RangeAllNoSort-4             200000              5155 ns/op
BenchmarkHashMapInt8_RangeAllAndSort-4            200000              7761 ns/op

PASS
ok      command-line-arguments  106.073s
```
All the performance tests are based on the comparison between orderdmap and golang's native HashMap. For type uint8, orderedmap has little performance loss, and for other types, if the data volume is very large, such as 1 million key-values, the single time consumption of orderdmap will be about 10 times of map, after all, map is O(1). You can run the test cases in testing by yourself.
