package handler

import (
	"context"

	ptypes "github.com/johnbellone/persona-service/internal/gen/persona/type"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type RealmService struct{}

func (h *RealmService) Create(ctx context.Context, req *ptypes.Realm, rsp *status.Status) error {
	return nil
}

func (h *RealmService) Get(ctx context.Context, req *pb.RealmRequest, rsp *ptypes.Realm) error {
	return nil
}

func (h *RealmService) Update(ctx context.Context, req *ptypes.Realm, rsp *status.Status) error {
	return nil
}

func (h *RealmService) Delete(ctx context.Context, req *pb.RealmRequest, rsp *status.Status) error {
	return nil
}
