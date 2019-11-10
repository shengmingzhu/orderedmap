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
