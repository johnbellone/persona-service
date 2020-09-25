package handler

import (
	"context"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/johnbellone/persona-service/proto/persona/api"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type AuthService struct{}

func (h *AuthService) Authenticate(ctx context.Context, req *pb.AuthRequest, rsp *pb.AuthResponse) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *AuthService) Refresh(ctx context.Context, req *pb.AuthRequest, rsp *pb.AuthResponse) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *AuthService) Validate(ctx context.Context, req *wrappers.StringValue, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		log.Trace("No metadata received.")
	}
	log.Debugf("metadata: %v\n", metadata)
	return nil
}
