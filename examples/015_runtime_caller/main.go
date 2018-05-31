package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtimeCaller1()
	runtimeCaller2()
}

func runtimeCaller1() {
	// stackは、0が上（最も最近詰まれたpcが0になる）

	// 0: 自分(=runtimeCaller1)自身
	fmt.Println("=== runtime.Caller(0) ===")
	pc, file, line, ok := runtime.Caller(0)
	fmt.Printf("pc: %v\n", pc)
	fmt.Printf("file: %v\n", file)
	fmt.Printf("line: %v\n", line)
	fmt.Printf("ok: %v\n", ok)

	// 1. 一つ上(=main)
	fmt.Println("=== runtime.Caller(1) ===")
	pc, file, line, ok = runtime.Caller(1)
	fmt.Printf("pc: %v\n", pc)
	fmt.Printf("file: %v\n", file)
	fmt.Printf("line: %v\n", line)
	fmt.Printf("ok: %v\n", ok)
}

func runtimeCaller2() {
	recorsiveFunc(0)
}

func recorsiveFunc(i int) int {
	fmt.Println(i)
	pcs := [20]uintptr{}
	runtime.Callers(0, pcs[:])
	fmt.Println(pcs)

	if i < 10 {
		return recorsiveFunc(i + 1)
	}
	return i
}
