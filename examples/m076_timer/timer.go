package main

// みんなのGo p.76 自前timeout

import (
	"fmt"
	"time"
)

func doSomething() error {
	time.Sleep(4 * time.Second)
	return nil
}

func main() {
	timer := time.NewTimer(3 * time.Second)
	done := make(chan error)
	// done := make(chan AsyncResult)

	go func() {
		done <- doSomething()
		// r, err = doSomething()
		// done <- AsyncResult{r, err}
	}()

	// timerの方が終了が早ければtimeout
	// doSomethingの方が終了が早ければOKかerrになる
	select {
	case <-timer.C:
		fmt.Println("Timeout!!")
	case err := <-done:
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("OK")
		// case result := <-done:
		// 	if result.err != nil {
		// 		...
		// 	}
		// 	...
	}
}
