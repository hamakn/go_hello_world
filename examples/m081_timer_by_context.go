package examples

// みんなのGo p.81 自前timeout by context package
// context packageはGo 1.7以降

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doSomething2(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.Done")
			wg.Done()
			return nil
		default:
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("Tick...")
		}
	}
}

func TimerByContext() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	done := make(chan error)

	wg.Add(1)
	go func() {
		done <- doSomething2(ctx)
	}()

	wg.Wait()
}
