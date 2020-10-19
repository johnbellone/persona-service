package handler

import (
	"context"

	pb "github.com/johnbellone/persona-service/proto/persona/api/v1"
	ptypes "github.com/johnbellone/persona-service/proto/persona/type"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type UserService struct{}

func (h *UserService) Create(ctx context.Context, req *pb.UserRequest, rsp *status.Status) error {
	return nil
}

func (h *UserService) Get(ctx context.Context, req *pb.UserRequest, rsp *ptypes.User) error {
	return nil
}

func (h *UserService) Update(ctx context.Context, req *pb.UserRequest, rsp *status.Status) error {
	return nil
}

func (h *UserService) Delete(ctx context.Context, req *pb.UserRequest, rsp *status.Status) error {
	return nil
}
