package handler

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/johnbellone/persona-service/proto/persona"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
)

type RoleService struct{}

func (h *RoleService) Create(ctx context.Context, req *pb.RoleRequest, rsp *pb.Role) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RoleService) Get(ctx context.Context, req *pb.RoleRequest, rsp *pb.Role) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RoleService) Update(ctx context.Context, req *pb.RoleRequest, rsp *pb.Role) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RoleService) Delete(ctx context.Context, req *pb.RoleRequest, rsp *emptypb.Empty) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
