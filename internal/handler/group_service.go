package handler

import (
	"context"

	pb "github.com/johnbellone/persona-service/proto/persona/api"
	ptypes "github.com/johnbellone/persona-service/proto/persona/type"
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/go-micro/v3/metadata"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type GroupService struct{}

func (h *GroupService) Create(ctx context.Context, req *pb.GroupRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *GroupService) Get(ctx context.Context, req *pb.GroupRequest, rsp *ptypes.Group) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *GroupService) Update(ctx context.Context, req *ptypes.Group, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *GroupService) Delete(ctx context.Context, req *pb.GroupRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
