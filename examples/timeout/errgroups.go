package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func sleepFunc(ctx context.Context, d time.Duration) error {
	end := time.Now().Add(d)
	for {
		select {
		case <-ctx.Done():
			return errors.New("timeout!!")
		default:
			if time.Now().After(end) {
				return nil
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// 全体のtimeout時間を制御したい
func main() {
	fmt.Println("timeout_no_errgroup()")
	timeout_no_errgroup()
	fmt.Println("timeout_errgroup()")
	timeout_errgroup()
	fmt.Println("timeout_errgroup_with_subfuncs()")
	timeout_errgroup_with_subfuncs()
}

func timeout_no_errgroup() {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	err := sleepFunc(ctx, 100*time.Millisecond)
	fmt.Println(err) // OK (100 < 250)
	err = sleepFunc(ctx, 100*time.Millisecond)
	fmt.Println(err) // also OK (200 < 250)
	err = sleepFunc(ctx, 100*time.Millisecond)
	fmt.Println(err) // timeout (300 > 250)
	err = sleepFunc(ctx, 100*time.Millisecond)
	fmt.Println(err) // also timeout (400 > 250)
}

func timeout_errgroup() {
	parentCtx := context.Background()

	// Action1: OK (100 < 250)
	g, ctx := errgroup.WithContext(parentCtx)
	ctx, cancel := context.WithTimeout(ctx, 250*time.Millisecond)
	defer cancel()

	g.Go(func() error {
		return sleepFunc(ctx, 100*time.Millisecond)
	})
	g.Go(func() error {
		return nil
	})

	fmt.Println(g.Wait())

	// Action2: OK (200 < 250)
	d := remainingDuration(ctx)
	g, ctx = errgroup.WithContext(parentCtx)
	ctx, cancel = context.WithTimeout(ctx, d)
	defer cancel()

	g.Go(func() error {
		return sleepFunc(ctx, 100*time.Millisecond)
	})
	g.Go(func() error {
		return nil
	})

	fmt.Println(g.Wait())

	// Action3: error (300 > 250)
	d = remainingDuration(ctx)
	g, ctx = errgroup.WithContext(parentCtx)
	ctx, cancel = context.WithTimeout(ctx, d)
	defer cancel()

	g.Go(func() error {
		return sleepFunc(ctx, 100*time.Millisecond)
	})
	g.Go(func() error {
		return nil
	})

	fmt.Println(g.Wait())
}

func remainingDuration(ctx context.Context) time.Duration {
	if deadline, ok := ctx.Deadline(); ok {
		return time.Until(deadline)
	}
	return 0
}

func timeout_errgroup_with_subfuncs() {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	err := subfunc_errgroup(ctx)
	fmt.Println(err)

	err = subfunc_errgroup(ctx)
	fmt.Println(err)

	err = subfunc_errgroup(ctx)
	fmt.Println(err)
}

func subfunc_errgroup(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return sleepFunc(ctx, 100*time.Millisecond)
	})
	g.Go(func() error {
		return nil
	})

	return g.Wait()
}
