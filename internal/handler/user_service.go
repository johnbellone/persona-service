package handler

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/johnbellone/persona-service/proto/persona"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
)

type UserService struct{}

func (h *UserService) Create(ctx context.Context, req *pb.UserRequest, rsp *pb.User) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *UserService) Get(ctx context.Context, req *pb.UserRequest, rsp *pb.User) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *UserService) Update(ctx context.Context, req *pb.UserRequest, rsp *pb.User) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *UserService) Delete(ctx context.Context, req *pb.UserRequest, rsp *emptypb.Empty) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
