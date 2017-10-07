package examples

// みんなのGo p.70 go-humanize

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
)

func GoHumanize() {
	if len(os.Args) <= 1 {
		fmt.Println("No os.Args[1]")
		return
	}
	name := os.Args[1]
	s, _ := os.Stat(name)
	fmt.Printf(
		"%s: %s\n",
		name,
		humanize.Bytes(uint64(s.Size())),
	)
}
