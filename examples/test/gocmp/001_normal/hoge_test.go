package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type s1 struct {
	I int
}

func TestGoCmp1(t *testing.T) {
	s11 := s1{1}
	s12 := s1{2}

	if diff := cmp.Diff(s11, s12); diff != "" {
		t.Errorf("diff (-want +got):\n%s", diff)
	}

	// 特定のfieldを無視して比較できる
	if diff := cmp.Diff(s11, s12, cmpopts.IgnoreFields(s12, "I")); diff != "" {
		t.Errorf("diff (-want +got):\n%s", diff)
	}
}

type s2 struct {
	i int
}

func TestGoCmp2(t *testing.T) {
	s21 := s2{1}
	s22 := s2{2}

	// cmp.AllowUnexported を付けるとunexported fieldもtestできる
	// 指定しなければpanic
	if diff := cmp.Diff(s21, s22, cmp.AllowUnexported(s2{})); diff != "" {
		t.Errorf("diff (-want +got):\n%s", diff)
	}

	// cmpopts.IgnoreUnexported を付けるとunexported fieldはtestされない
	if diff := cmp.Diff(s21, s22, cmpopts.IgnoreUnexported(s2{})); diff != "" {
		t.Errorf("diff (-want +got):\n%s", diff)
	}
}
