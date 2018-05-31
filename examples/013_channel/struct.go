package main

// https://twitter.com/sonatard/status/985836534677913601
// chan boolよりも chan struct{}を使った方が良い理由は確保するメモリサイズの違いだと思っていたけど試したことなかったので試してみました
// 想像通りstruct{}{}は値として0byteでした
// https://play.golang.org/p/QWE4g8H1tdz #golangtokyo

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		c <- struct{}{}
		close(c)
	}()

	select {
	case x, opened := <-c:
		fmt.Println(x)
		fmt.Println(opened)
	}
}
