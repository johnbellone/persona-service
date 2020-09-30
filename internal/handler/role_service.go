package handler

import (
	"context"

	pb "github.com/johnbellone/persona-service/proto/persona/api"
	ptypes "github.com/johnbellone/persona-service/proto/persona/type"
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/go-micro/v3/metadata"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type RoleService struct{}

func (h *RoleService) Create(ctx context.Context, req *pb.RoleRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RoleService) Get(ctx context.Context, req *pb.RoleRequest, rsp *ptypes.Role) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RoleService) Update(ctx context.Context, req *pb.RoleRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RoleService) Delete(ctx context.Context, req *pb.RoleRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
