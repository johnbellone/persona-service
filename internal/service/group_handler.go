package service

import (
	"context"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	ptypes "github.com/johnbellone/persona-service/internal/gen/persona/type"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GroupHandler struct{}

func (h *GroupHandler) Create(ctx context.Context, req *pb.GroupRequest) (*pb.GroupResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *GroupHandler) Get(ctx context.Context, req *pb.GroupRequest) (*ptypes.Group, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *GroupHandler) Update(ctx context.Context, req *ptypes.Group) (*pb.GroupResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *GroupHandler) Delete(ctx context.Context, req *pb.GroupRequest) (*pb.GroupResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}
