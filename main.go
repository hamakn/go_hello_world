package main

import (
	"fmt"

	"github.com/hamakn/go_hello_world/examples"
)

func main() {
	fmt.Println("== 001 Hello ==")
	examples.Hello()

	// fmt.Println("== m034 Httpd ==")
	// examples.Httpd()

	fmt.Println("== m070 go-humanize ==")
	examples.GoHumanize()

	fmt.Println("== m076 Timer ==")
	examples.Timer()

	fmt.Println("== m081 TimerByContext ==")
	examples.TimerByContext()
}
