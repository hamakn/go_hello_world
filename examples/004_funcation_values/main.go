package main

import (
	"fmt"
	"math"
)

// 2. 関数はクロージャなので、変数はレキシカル変数になる
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		i, j = j, i+j
		return i
	}
}

func main() {
	// 1. 関数を値として扱える
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Print("hypot(3, 4): ")
	fmt.Println(hypot(3, 4))

	// 2. レキシカル変数は(当然)変数ごとに保持される
	pos, neg := adder(), adder()
	fmt.Print("adder: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("(%d, %d)", pos(i), neg(-2*i))
	}
	fmt.Println("")

	fmt.Print("fibonacci: ")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	fmt.Println("")
}
