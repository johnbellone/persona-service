package handler

import (
	"context"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/johnbellone/persona-service/proto/persona/api/v1"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type AuthService struct{}

func (h *AuthService) Authenticate(ctx context.Context, req *pb.AuthRequest, rsp *pb.AuthResponse) error {
	return nil
}

func (h *AuthService) Refresh(ctx context.Context, req *pb.AuthRequest, rsp *pb.AuthResponse) error {
	return nil
}

func (h *AuthService) Validate(ctx context.Context, req *wrappers.StringValue, rsp *status.Status) error {
	return nil
}
