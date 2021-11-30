package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// merpay-apiのsrc/app/api/identify_qrcode_type.goがleakするんじゃね感あったので調べたやつ
// 結論としてはearly returnしてもleakしない
// timeoutしているしそりゃそうか??

func main() {
	printGoroutineCount()
}

func printGoroutineCount() {
	fmt.Println(runtime.NumGoroutine())
	for {
		runGoroutines(10)
		fmt.Println(runtime.NumGoroutine())
	}
}

func runGoroutines(count int) int {
	ch := make(chan int, count)

	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()

	for i := 0; i < count; i++ {
		go func(i int) {
			time.Sleep(100 * time.Millisecond)
			ch <- i
		}(i)
	}

	for {
		select {
		case v := <-ch:
			return v
		case <-ctx.Done():
			return -1
		}
	}
}
