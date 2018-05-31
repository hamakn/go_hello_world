package main

// run test
// $ go test ./m125_gotest/examples/...
//
// run test and benchmark
// $ go test ./m125_gotest/examples/... -bench .

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}

func ExampleSum() {
	fmt.Println(sum(1, 2))
	// Output: 3
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1, 2)
	}
}
