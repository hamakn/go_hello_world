package examples

import (
	"flag"
	"fmt"
)

var string_flag *string = flag.String(
	"string_flag",
	"default string flag",
	"help message for string_flag",
)

var int_flag *int = flag.Int(
	"int_flag",
	0,
	"help message for int_flag",
)

func ParseOptionsExample() {
	fmt.Println("use command line options like: --string_flag aaa --int_flag 42")
	flag.Parse()
	fmt.Println(*string_flag)
	fmt.Println(*int_flag)
}
