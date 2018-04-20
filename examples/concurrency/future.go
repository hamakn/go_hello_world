package main

import (
	"fmt"
	"time"
)

type futureInt struct {
	result int
	done   chan struct{}
}

func (f *futureInt) Get() int {
	<-f.done // block until close channel
	return f.result
}

func getFutureInt() *futureInt {
	f := &futureInt{
		done: make(chan struct{}),
	}
	go func() {
		time.Sleep(time.Second)
		f.result = 1
		close(f.done)
	}()
	return f
}

func main() {
	// Getの方がgoroutineの終了より先に呼ばれるので終わるまでblockする
	f1 := getFutureInt()
	fmt.Println("call Get")
	fmt.Println(f1.Get())

	// Getの方がgoroutineの終了より後に呼ばれるので値が即座に返ってくる
	f2 := getFutureInt()
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("call Get")
	fmt.Println(f2.Get())
}
