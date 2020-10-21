package service

import (
	"context"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	ptypes "github.com/johnbellone/persona-service/internal/gen/persona/type"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct{}

func (h *UserHandler) Create(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *UserHandler) Get(ctx context.Context, req *pb.UserRequest) (*ptypes.User, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *UserHandler) Update(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *UserHandler) Delete(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}
