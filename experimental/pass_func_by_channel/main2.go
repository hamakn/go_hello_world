package main

import (
	"fmt"
	"time"
)

type i interface {
	f(i int) int
}

// main.go 相当
func main() {
	ch := make(chan i)
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
	ch chan i
}

func NewSubscriber(ch chan i) *subscriber {
	return &subscriber{
		ch: ch,
	}
}

func (s *subscriber) Run() {
	ii := <-s.ch
	fmt.Println("ready")
	fmt.Println(ii.f(99))
}

// server.RunGRPCServer 相当
type s struct{}

func (ss *s) f(i int) int {
	return i * i
}

func RunServer(ch chan i) {
	ss := &s{}
	time.Sleep(3 * time.Second)
	ch <- ss
}
