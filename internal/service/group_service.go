package handler

import (
	"context"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	ptypes "github.com/johnbellone/persona-service/internal/gen/persona/type"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type GroupService struct{}

func (h *GroupService) Create(ctx context.Context, req *pb.GroupRequest, rsp *status.Status) error {
	return nil
}

func (h *GroupService) Get(ctx context.Context, req *pb.GroupRequest, rsp *ptypes.Group) error {
	return nil
}

func (h *GroupService) Update(ctx context.Context, req *ptypes.Group, rsp *status.Status) error {
	return nil
}

func (h *GroupService) Delete(ctx context.Context, req *pb.GroupRequest, rsp *status.Status) error {
	return nil
}
