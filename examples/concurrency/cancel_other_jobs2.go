package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// cancel_other_jobsを汎用っぽくしてみたやつ
func main() {
	// errorが起きないので、3秒待つ
	err := ExecAllOrAbortOnError(
		context.Background(),
		[]func() error{
			func() error {
				time.Sleep(time.Duration(1) * time.Second)
				return nil
			},
			func() error {
				time.Sleep(time.Duration(3) * time.Second)
				return nil
			},
		},
	)
	fmt.Println(err)

	// errorが起きるので、10秒待つことなく抜ける
	err = ExecAllOrAbortOnError(
		context.Background(),
		[]func() error{
			func() error {
				time.Sleep(time.Duration(10) * time.Second)
				return nil
			},
			func() error {
				return errors.New("error!!")
			},
		},
	)
	fmt.Println(err)
}

func ExecAllOrAbortOnError(ctx context.Context, funcs []func() error) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan error)
	jobs := len(funcs)
	done := 0

	for _, f := range funcs {
		f := f
		go func() {
			err := f()
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
				return err
			}
			if jobs == done {
				return nil
			}
		}
	}
}
