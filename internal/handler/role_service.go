package handler

import (
	"context"

	pb "github.com/johnbellone/persona-service/proto/persona/api/v1"
	ptypes "github.com/johnbellone/persona-service/proto/persona/type"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type RoleService struct{}

func (h *RoleService) Create(ctx context.Context, req *pb.RoleRequest, rsp *status.Status) error {
	return nil
}

func (h *RoleService) Get(ctx context.Context, req *pb.RoleRequest, rsp *ptypes.Role) error {
	return nil
}

func (h *RoleService) Update(ctx context.Context, req *pb.RoleRequest, rsp *status.Status) error {
	return nil
}

func (h *RoleService) Delete(ctx context.Context, req *pb.RoleRequest, rsp *status.Status) error {
	return nil
}
