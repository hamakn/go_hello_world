package main

import "fmt"

type I interface {
	Hoge()
}

type S struct {
}

func (s *S) Hoge() {
}

func main() {
	var i *I
	s := S{}
	fmt.Println(i)
	pi := hoge(&s)
	fmt.Println(pi)
}

func hoge(i I) *I {
	fmt.Println(i)
	fmt.Println(&i)
	return &i
}
