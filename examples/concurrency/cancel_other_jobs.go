package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// 1つでもerrorがあるか、全部終わったら処理を抜ける
func main() {
	// errorが起きないので、3秒待つ
	hoge([]int{1, 3})
	// errorが起きるので1,3,5秒待つことなく抜ける
	hoge([]int{1, 2, 3, 4, 5})
}

func hoge(arr []int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan error)
	len := len(arr)
	done := 0

	for _, i := range arr {
		i := i
		go func() {
			var err error
			if i%2 == 0 {
				err = errors.New(fmt.Sprintf("error: %v", i))
			} else {
				// block
				time.Sleep(time.Duration(i) * time.Second)
			}
			select {
			case c <- err:
			case <-ctx.Done():
			}
		}()
	}

	for {
		select {
		case err := <-c:
			done++
			if err != nil {
				fmt.Println(err)
				return
			}
			if len == done {
				fmt.Println("all OK")
				return
			}
		}
	}
}
