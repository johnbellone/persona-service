package handler

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/johnbellone/persona-service/proto/persona"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
)

type GroupService struct{}

func (h *GroupService) Create(ctx context.Context, req *pb.GroupRequest, rsp *pb.Group) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *GroupService) Get(ctx context.Context, req *pb.GroupRequest, rsp *pb.Group) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *GroupService) Update(ctx context.Context, req *pb.GroupRequest, rsp *pb.Group) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *GroupService) Delete(ctx context.Context, req *pb.GroupRequest, rsp *emptypb.Empty) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
