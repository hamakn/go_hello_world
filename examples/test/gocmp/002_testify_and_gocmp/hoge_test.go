package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type s struct {
	I int
}

func TestByTestify(t *testing.T) {
	s1 := s{1}
	s2 := s{2}

	require.Equal(t, s1, s2, "s2 should equal to s1")
}

func TestByGoCmp(t *testing.T) {
	s1 := s{1}
	s2 := s{2}

	RequireEqual(t, s1, s2, "s2 should equal to s1")
}
