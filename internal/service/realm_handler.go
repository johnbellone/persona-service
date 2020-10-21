package service

import (
	"context"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	ptypes "github.com/johnbellone/persona-service/internal/gen/persona/type"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RealmHandler struct{}

func (h *RealmHandler) Create(ctx context.Context, req *ptypes.Realm) (*pb.RealmResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *RealmHandler) Get(ctx context.Context, req *pb.RealmRequest) (*ptypes.Realm, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *RealmHandler) Update(ctx context.Context, req *ptypes.Realm) (*pb.RealmResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}

func (h *RealmHandler) Delete(ctx context.Context, req *pb.RealmRequest) (*pb.RealmResponse, error) {
	return nil, status.Error(codes.NotFound, "Not implemented")
}
