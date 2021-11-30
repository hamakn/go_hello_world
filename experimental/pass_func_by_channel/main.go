package main

import (
	"fmt"
	"time"
)

// main.go 相当
func main() {
	ch := make(chan (func(i int) int))
	sub := NewSubscriber(ch)

	go func() {
		RunServer(ch)
	}()
	go func() {
		sub.Run()
	}()

	time.Sleep(5 * time.Second) // waitk
}

// subscriber 相当
type subscriber struct {
	ch chan func(i int) int
}

func NewSubscriber(ch chan (func(i int) int)) *subscriber {
	return &subscriber{
		ch: ch,
	}
}

func (s *subscriber) Run() {
	f := <-s.ch
	fmt.Println("ready")
	fmt.Println(f(99))
}

// server.RunGRPCServer 相当
func RunServer(ch chan (func(i int) int)) {
	f := func(i int) int {
		return i * i
	}
	time.Sleep(3 * time.Second)
	ch <- f
}
