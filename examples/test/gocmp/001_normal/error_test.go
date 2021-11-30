package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestErrorCmp(t *testing.T) {
	err := errors.New("hoge")

	if diff := cmp.Diff(err, err, cmpopts.EquateErrors()); diff != "" { // OK
		t.Errorf("diff (-want +got):\n%s", diff)
	}

	// errorsの場合は同じものじゃないとダメ
	if diff := cmp.Diff(err, errors.New("hoge"), cmpopts.EquateErrors()); diff != "" { // NG
		t.Errorf("diff (-want +got):\n%s", diff)
	}
	// =>
	//  TestErrorCmp: error_test.go:22: diff (-want +got):
	//        (*errors.errorString)(
	//      -       e"hoge",
	//      +       e"hoge",
	//        )

	// 別messageだと当然diffが出る
	if diff := cmp.Diff(err, errors.New("fuga"), cmpopts.EquateErrors()); diff != "" { // NG
		t.Errorf("diff (-want +got):\n%s", diff)
	}
	// =>
	//  TestErrorCmp: error_test.go:26: diff (-want +got):
	//        (*errors.errorString)(
	//      -       e"hoge",
	//      +       e"fuga",
	//        )

	// comparerを指定しないとpanic
	if diff := cmp.Diff(err, err); diff != "" { // OK
		t.Errorf("diff (-want +got):\n%s", diff)
	}
	// panic: cannot handle unexported field at {*errors.errorString}.s:
	//         "errors".errorString
	// consider using a custom Comparer; if you control the implementation of type, you can also consider using an Exporter, AllowUnexported, or cmpopts.IgnoreUnexported [recovered]
	//         panic: cannot handle unexported field at {*errors.errorString}.s:
	//         "errors".errorString
	// consider using a custom Comparer; if you control the implementation of type, you can also consider using an Exporter, AllowUnexported, or cmpopts.IgnoreUnexported
}

func TestErrorCmpGRPCStatus(t *testing.T) {
	st := status.New(codes.Internal, "internal error")

	// gRPC statusだと別物でも通る
	if diff := cmp.Diff(st.Err(), status.New(codes.Internal, "internal error").Err(), cmpopts.EquateErrors()); diff != "" { // OK
		t.Errorf("diff (-want +got):\n%s", diff)
	}

	// grpc statusならcomparer いらない
	if diff := cmp.Diff(st.Err(), status.New(codes.Internal, "internal error").Err()); diff != "" { // OK
		t.Errorf("diff (-want +got):\n%s", diff)
	}

	// messageやcodeが違えばdiffになる
	if diff := cmp.Diff(st.Err(), status.New(codes.Unknown, "internal error2").Err(), cmpopts.EquateErrors()); diff != "" { // NG
		t.Errorf("diff (-want +got):\n%s", diff)
	}
	// =>
	//  TestErrorCmpGRPCStatus: error_test.go:45: diff (-want +got):
	//        (*status.Error)(
	//      -       e"rpc error: code = Internal desc = internal error",
	//      +       e"rpc error: code = Unknown desc = internal error2",
	//        )
}
