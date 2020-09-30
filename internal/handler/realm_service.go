package handler

import (
	"context"

	pb "github.com/johnbellone/persona-service/proto/persona/api"
	ptypes "github.com/johnbellone/persona-service/proto/persona/type"
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/go-micro/v3/metadata"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type RealmService struct{}

func (h *RealmService) Create(ctx context.Context, req *ptypes.Realm, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RealmService) Get(ctx context.Context, req *pb.RealmRequest, rsp *ptypes.Realm) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)

	return nil
}

func (h *RealmService) Update(ctx context.Context, req *ptypes.Realm, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *RealmService) Delete(ctx context.Context, req *pb.RealmRequest, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
