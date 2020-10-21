package service

import (
	"context"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthHandler struct{}

func (h *AuthHandler) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *AuthHandler) Refresh(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	_, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}
	return nil, status.Error(codes.NotFound, "Not implemented")
}
