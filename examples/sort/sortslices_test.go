package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestSortSlices(t *testing.T) {
	arr1 := []string{"a", "b"}
	arr2 := []string{"b", "a"}

	fmt.Println("==")
	d1 := cmp.Diff(arr1, arr2, cmpopts.SortSlices(func(s1, s2 string) bool { return s1 > s2 }))
	fmt.Println(d1)

	fmt.Println("==")
	d2 := cmp.Diff(arr1, arr2)
	fmt.Println(d2)

	fmt.Println("==")
	d3 := cmp.Diff(arr2, arr1, cmpopts.SortSlices(func(s1, s2 string) bool { return s1 > s2 }))
	fmt.Println(d3)

	fmt.Println("==")
	d4 := cmp.Diff(arr1, arr2, cmpopts.SortSlices(func(s1, s2 int) bool { return s1 > s2 }))
	fmt.Println(d4)
}
