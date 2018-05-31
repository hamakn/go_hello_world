package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	arr := []int{1, 2, 3, 4, 5}
	res := []int{}

	for _, i := range arr {
		go func(x int) {
			mu.Lock()
			defer mu.Unlock()

			res = append(res, x*x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(res)
}
