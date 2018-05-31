package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	errGroup1()
	errGroup2()
}

// Waitは全ての実行後、最初のerrorを返す
func errGroup1() {
	var eg errgroup.Group

	for i := 0; i < 5; i++ {
		i := i
		eg.Go(func() error {
			if i%2 == 1 {
				return errors.New(fmt.Sprintf("error: %v", i))
			} else {
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("no error")
	}
}

// Waitは全ての実行後、最初のerrorを返すのは変わらないが、
// 最初のerrorが入った時点でcancelが呼ばれるので、他のgoroutineがctx.Done()の処理により抜けてくる
// ただし、ctx.Done()で明示的に抜けるようにコードを書かないと当然抜けてこない
func errGroup2() {
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := 0; i < 5; i++ {
		i := i
		eg.Go(func() error {
			if i%2 == 1 {
				return errors.New(fmt.Sprintf("error: %v", i))
			}
			for j := 0; j < i*i; j++ {
				select {
				case <-ctx.Done():
					return nil
				default:
					time.Sleep(time.Second)
				}
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("no error")
	}
}
