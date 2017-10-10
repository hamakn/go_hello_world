package examples

import "fmt"

func printIntSlice(label string, x []int) {
	fmt.Printf(
		"%s: len=%d cap=%d %v\n",
		label,
		len(x),
		cap(x),
		x,
	)
}

func SliceExample() {
	slice := []int{1, 1, 2, 3, 5, 8}
	printIntSlice("`slice := []int{1, 1, 2, 3, 5, 8}`", slice)

	// slicing
	printIntSlice("`slice[1:4]`", slice[1:4])
	printIntSlice("`slice[:3]`", slice[:3])
	printIntSlice("`slice[4:]`", slice[4:])

	// make([]int, len, cap)
	printIntSlice("`make([]int, 5)`", make([]int, 5))
	printIntSlice("`make([]int, 0, 5)`", make([]int, 0, 5))

	// sliceの初期値はnil
	var z []int
	if z == nil {
		fmt.Println("`var z []int`: z == nil!!")
	}

	// range
	for i, v := range slice {
		fmt.Printf("slice[%d]: %d\n", i, v)
	}

	// append
	slice = append(slice, 13)
	printIntSlice("`append(slice, 13)`", slice)

	// 多次元配列(5x6)
	height := 5
	width := 6
	field := make([][]uint, height)
	for i := 0; i < height; i++ {
		field[i] = make([]uint, width)
	}
	field[0][1] = 1
	field[4][5] = 2
	fmt.Println(field)
}
