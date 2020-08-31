package handler

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/johnbellone/persona-service/proto/persona"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
)

type RealmService struct{}

func (h *RealmService) Create(ctx context.Context, req *pb.RealmRequest, rsp *pb.Realm) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RealmService) Get(ctx context.Context, req *pb.RealmRequest, rsp *pb.Realm) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RealmService) Update(ctx context.Context, req *pb.RealmRequest, rsp *pb.Realm) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RealmService) Delete(ctx context.Context, req *pb.RealmRequest, rsp *emptypb.Empty) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
