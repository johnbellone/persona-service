package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PersonHandler struct{}

func (h *PersonHandler) Get(ctx context.Context, req *wrappers.StringValue) (*pb.PersonResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}
