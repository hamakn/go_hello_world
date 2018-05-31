package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// breakは必要ない
	fmt.Print("Go runs os on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s", os)
	}

	fmt.Print("When's Monday?: ")
	today := time.Now().Weekday()
	switch time.Monday {
	case (today + 0) % 7:
		fmt.Println("Today")
	case (today + 1) % 7:
		fmt.Println("Tommorow")
	case (today + 2) % 7:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	// 条件のないswitchは、if-then-elseをシンプルに書ける
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
