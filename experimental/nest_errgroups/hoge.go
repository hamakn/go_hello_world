package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	r1, r2, r3, r4, err := hoge()
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(err)
}

func hoge() (int, int, int, int, error) {
	var (
		r1 int
		r2 int
		r3 int
		r4 int
	)

	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	g.Go(func() error {
		r1 = 1
		return timeoutFunc(ctx, 1)
	})

	g.Go(func() error {
		r2 = 2
		return timeoutFunc(ctx, 2)
	})

	g.Go(func() error {
		var err error
		r3, r4, err = fuga(ctx)
		return err
	})

	if err := g.Wait(); err != nil {
		return 0, 0, 0, 0, err
	}

	return r1, r2, r3, r4, nil
}

func fuga(ctx context.Context) (int, int, error) {
	var (
		r3 int
		r4 int
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		r3 = 3
		return timeoutFunc(ctx, 3)
	})

	g.Go(func() error {
		r4 = 4
		return timeoutFunc(ctx, 4)
	})

	if err := g.Wait(); err != nil {
		return 0, 0, err
	}

	return r3, r4, nil
}

func sleepFunc(ms time.Duration) {
	time.Sleep(ms * time.Millisecond)
}

func timeoutFunc(ctx context.Context, waitCount int) error {
	c := 0

	for {
		select {
		case <-ctx.Done():
			return errors.New("timeout!!")
		default:
			if waitCount < c {
				return nil
			}
			time.Sleep(100 * time.Millisecond)
			c++
		}
	}
}
