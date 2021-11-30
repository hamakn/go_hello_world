package main

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	hello "google.golang.org/grpc/examples/helloworld/helloworld"
)

func TestProtobuf1(t *testing.T) {
	h1 := &hello.HelloRequest{Name: "hello"}
	h2 := &hello.HelloRequest{Name: "hello", XXX_sizecache: 10}

	// 普通XXX_sizecacheはcompareして欲しくない
	if diff := cmp.Diff(h1, h2); diff != "" {
		t.Errorf("diff (-want +got):\n%s", diff)
	}
}

func TestProtobuf2(t *testing.T) {
	h1 := &hello.HelloRequest{Name: "hello"}
	h2 := &hello.HelloRequest{Name: "hello", XXX_sizecache: 10}

	// proto.Equalを使えばいけるらしい
	if diff := cmp.Diff(h1, h2, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("diff (-want +got):\n%s", diff)
	}
}
