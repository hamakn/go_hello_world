package main

// $ go test ./examples/

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
