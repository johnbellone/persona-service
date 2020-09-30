package handler

import (
	"context"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/johnbellone/persona-service/proto/persona/api"
	"github.com/micro/go-micro/v3/metadata"
	"github.com/micro/micro/v3/service/logger"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type AuthService struct{}

func (h *AuthService) Authenticate(ctx context.Context, req *pb.AuthRequest, rsp *pb.AuthResponse) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		logger.Trace("No metadata received.")
	}
	logger.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *AuthService) Refresh(ctx context.Context, req *pb.AuthRequest, rsp *pb.AuthResponse) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		logger.Trace("No metadata received.")
	}
	logger.Debugf("metadata: %v\n", metadata)
	return nil
}

func (h *AuthService) Validate(ctx context.Context, req *wrappers.StringValue, rsp *status.Status) error {
	metadata, ok := metadata.FromContext(ctx)
	if !ok {
		logger.Trace("No metadata received.")
	}
	logger.Debugf("metadata: %v\n", metadata)
	return nil
}
