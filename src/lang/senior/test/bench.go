package test

// MakeSliceWithoutAlloc 不预分配
func MakeSliceWithoutAlloc() []int {
	var newSlice []int
	for i := 0; i < 10000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}

// MakeSliceWithPreAlloc 预分配
func MakeSliceWithPreAlloc() []int {
	var newSlice []int
	newSlice = make([]int, 100000)
	for i := 0; i < 10000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}
