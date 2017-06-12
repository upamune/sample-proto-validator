package main

import (
	"github.com/upamune/sample-proto-validator/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type UserService struct {
}

func (us UserService) Create(ctx context.Context, req *api.Request) (*api.Response, error) {
	if err := req.Validate(); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%s", err)
	}

	res := &api.Response{}
	return res, nil
}

func NewServer() *grpc.Server {
	s := grpc.NewServer()
	api.RegisterUserSerivceServer(s, UserService{})
	return s
}
