package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// errgroupで1つも待たないとどうなるんだっけ??
func main() {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
