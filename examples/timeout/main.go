package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("boooom!!")
			return
		default:
			fmt.Println("tick...")
			time.Sleep(1 * time.Second)
		}
	}
}
