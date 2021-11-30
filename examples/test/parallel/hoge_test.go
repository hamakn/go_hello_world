package main

import (
	"fmt"
	"testing"
	"time"
)

func sleepFunc(name string) {
	fmt.Printf(">>> %v start\n", name)
	time.Sleep(5 * time.Second)
	fmt.Printf(">>> %v end\n", name)
}

func TestHoge(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
	}{
		"a": {
			//
		},
		"b": {
			//
		},
	}

	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			fmt.Println(tt)
			sleepFunc(name)
		})
	}
}

func TestFuga(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
	}{
		"c": {
			//
		},
		"d": {
			//
		},
	}

	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			fmt.Println(tt)
			sleepFunc(name)
		})
	}
}
