package service

import (
	"context"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	ptypes "github.com/johnbellone/persona-service/internal/gen/persona/type"
	"github.com/johnbellone/persona-service/internal/server"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	config *server.Config
}

func NewUserHandler(c *server.Config) *UserHandler {
	return &UserHandler{config: c}
}

func (h *UserHandler) Create(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(""), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	_ = &ptypes.User{
		Password: string(hash),
	}

	return &pb.UserResponse{}, nil
}

func (h *UserHandler) Get(ctx context.Context, req *pb.UserRequest) (*ptypes.User, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (h *UserHandler) Update(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (h *UserHandler) Delete(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}
