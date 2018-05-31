package main

import "fmt"

func variadicAdder(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func arrayAdder(args []int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	nums := []int{10, 20, 30}
	fmt.Println(variadicAdder(nums...))
	fmt.Println(arrayAdder(nums))
}
