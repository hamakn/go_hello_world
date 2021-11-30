package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func RequireEqual(t *testing.T, want, got interface{}, msg string, options ...cmp.Option) {
	t.Helper()
	if diff := cmp.Diff(want, got, options...); diff != "" {
		t.Fatalf("RequireEqual unmatched:\n message: %s\n (-want +got):\n%s", msg, diff)
	}
}

// TODO:
// func RequireNil
// func RequireNotNil
// func RequireError
