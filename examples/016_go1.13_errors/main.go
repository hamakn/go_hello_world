package main

import (
	"errors"
	"fmt"
	"time"
)

type MyErrorInterface interface {
	GetErrorTime() time.Time
}

type MyError struct {
	m string
	t time.Time
}

func (e *MyError) Error() string {
	return e.m
}

func (e *MyError) GetErrorTime() time.Time {
	return e.t
}

type MyError2 struct {
	m string
	i int
}

func (e *MyError2) Error() string {
	return e.m
}

func (e *MyError2) GetI() int {
	return e.i
}

func main() {
	e := &MyError{"Error!!", time.Now()}
	fmt.Printf("元のエラー: %v\n", e)

	// wrap && unwrap
	wrappedError := fmt.Errorf("wrapping: %w", e)
	fmt.Printf("%%w によってwrapされたエラー: %v\n", wrappedError)
	fmt.Printf("unwrapされて元に戻ったエラー: %v\n", errors.Unwrap(wrappedError))

	fmt.Printf("Unwrapできないとnilになる（元のエラーはwrapされていない）: %v\n", errors.Unwrap(e))
	fmt.Printf("Unwrapできないとnilになる（wrapしたエラーを1回UnwrapしたやつをUnwrapしてもnil）: %v\n", errors.Unwrap(errors.Unwrap(wrappedError)))

	// As
	var myError *MyError
	if errors.As(wrappedError, &myError) {
		fmt.Println(myError.GetErrorTime())
	}

	// Is
	fmt.Println(errors.Is(wrappedError, myError))
	fmt.Println(errors.Is(wrappedError, &MyError{}))
	fmt.Println(errors.Is(errors.New("Error!!"), myError))
}
