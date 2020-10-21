package service

import (
	"context"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	ptypes "github.com/johnbellone/persona-service/internal/gen/persona/type"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoleHandler struct{}

func (h *RoleHandler) Create(ctx context.Context, req *pb.RoleRequest) (*pb.RoleResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *RoleHandler) Get(ctx context.Context, req *pb.RoleRequest) (*ptypes.Role, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *RoleHandler) Update(ctx context.Context, req *pb.RoleRequest) (*pb.RoleResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *RoleHandler) Delete(ctx context.Context, req *pb.RoleRequest) (*pb.RoleResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}
