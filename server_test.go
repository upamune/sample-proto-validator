package main

import (
	"fmt"
	"net"
	"testing"

	"github.com/upamune/sample-proto-validator/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type rpcError struct {
	code codes.Code
	desc string
}

func (e rpcError) Error() string {
	return fmt.Sprintf("rpc error: code = %d desc = %s", e.code, e.desc)
}

func TestValidator(t *testing.T) {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		t.Fatal(err)
	}

	s := NewServer()
	go func() {
		if err := s.Serve(l); err != nil {
			t.Fatal(err)
		}
	}()

	cases := []struct {
		Request      api.Request
		IsError bool
		ExpectedCode codes.Code
	}{
		{
			Request: api.Request{
				User: &api.User{
					Id:     "serizawa",
					Age:    22,
					Height: 163,
				},
			},
			IsError: false,
		},
		{
			Request:      api.Request{},
			IsError: true,
			ExpectedCode: codes.InvalidArgument,
		},
		{
			Request: api.Request{
				User: &api.User{
					Id:     "toolonguserid",
					Age:    30,
					Height: 180,
				},
			},
			IsError: true,
			ExpectedCode: codes.InvalidArgument,
		},
		{
			Request: api.Request{
				User: &api.User{
					Id:     "serizawa",
					Age:    -1,
					Height: 180,
				},
			},
			IsError: true,
			ExpectedCode: codes.InvalidArgument,
		},
		{
			Request: api.Request{
				User: &api.User{
					Id:     "serizawa",
					Age:    22,
					Height: 0,
				},
			},
			IsError: false,
		},
		{
			Request: api.Request{
				User: &api.User{
					Id:     "serizawa",
					Age:    22,
					Height: 250.1,
				},
			},
			IsError: true,
			ExpectedCode: codes.InvalidArgument,
		},
	}

	for _, c := range cases {
		conn, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			t.Fatal(err)
		}
		client := api.NewUserSerivceClient(conn)
		ctx := context.Background()
		_, err = client.Create(ctx, &c.Request)
		if err == nil && !c.IsError {
			continue
		}

		if c.ExpectedCode != grpc.Code(err) {
			t.Errorf("c.ExpectedCode (%d) != grpc.Code(err) (%d)", c.ExpectedCode, grpc.Code(err))
		}

	}

}
