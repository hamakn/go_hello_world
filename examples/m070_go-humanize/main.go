package main

// みんなのGo p.70 go-humanize

import (
	"flag"
	"fmt"
	"os"

	humanize "github.com/dustin/go-humanize"
)

var filename_flag *string = flag.String(
	"filename",
	"",
	"filename to show filesize on m070_go-humanize",
)

func main() {
	flag.Parse()
	if *filename_flag == "" {
		fmt.Println("No option --filename")
		return
	}
	s, _ := os.Stat(*filename_flag)
	fmt.Printf(
		"%s: %s\n",
		*filename_flag,
		humanize.Bytes(uint64(s.Size())),
	)
}
