package main

import (
	"fmt"
	"time"
)

func main() {
	// 同時並列数2で処理をする
	maxConcurrency := 2
	sem := make(chan struct{}, maxConcurrency)

	for i := 0; i < 11; i++ {
		go func(x int) {
			sem <- struct{}{}
			defer func() { <-sem }()

			fmt.Println(x)
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Duration(6) * time.Second)
}
