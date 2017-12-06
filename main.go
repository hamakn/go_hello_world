package main

import (
	"fmt"
	"math"

	"github.com/hamakn/go_hello_world/examples"
)

func myPrintln(str string) {
	fmt.Println("")
	fmt.Println(str)
}

func main() {
	myPrintln("== 001 Hello ==")
	examples.Hello()

	myPrintln("== 002 Newton ==")
	fmt.Println(examples.MySqrt(2))
	fmt.Println(math.Sqrt(2))

	myPrintln("== 003 Struct ==")
	examples.StructExample()

	myPrintln("== 004 Function values ==")
	examples.FunctionValuesExample()

	myPrintln("== 005 Variadic function ==")
	examples.VariadicFunctionExample()

	myPrintln("== 006 Switch ==")
	examples.SwitchExample()

	myPrintln("== 007 Parse options ==")
	examples.ParseOptionsExample()

	myPrintln("== 008 JSON ==")
	examples.JSONExample()

	myPrintln("== 009 Slice ==")
	examples.SliceExample()

	myPrintln("== 010 Map ==")
	examples.MapExample()

	// fmt.Println("== m034 Httpd ==")
	// examples.Httpd()

	myPrintln("== m070 go-humanize ==")
	examples.GoHumanize()

	myPrintln("== m076 Timer ==")
	examples.Timer()

	myPrintln("== m081 TimerByContext ==")
	examples.TimerByContext()

	myPrintln("== 011 HttpClient ==")
	examples.HttpClientExample()
}
