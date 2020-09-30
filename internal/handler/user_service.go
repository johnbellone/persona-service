package handler

import (
	"context"

	pb "github.com/johnbellone/persona-service/proto/persona/api"
	ptypes "github.com/johnbellone/persona-service/proto/persona/type"
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/go-micro/v3/metadata"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type UserService struct{}

func (h *UserService) Create(ctx context.Context, req *pb.UserRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *UserService) Get(ctx context.Context, req *pb.UserRequest, rsp *ptypes.User) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *UserService) Update(ctx context.Context, req *pb.UserRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *UserService) Delete(ctx context.Context, req *pb.UserRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
