// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: persona/realm.proto

package persona

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RealmService service

func NewRealmServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "RealmService.Create",
			Path:    []string{"/api/v1/realms"},
			Method:  []string{"POST"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "RealmService.Get",
			Path:    []string{"/api/v1/realms/{name}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "RealmService.Update",
			Path:    []string{"/api/v1/realms/{name}"},
			Method:  []string{"PUT"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "RealmService.Delete",
			Path:    []string{"/api/v1/realms/{name}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
	}
}

// Client API for RealmService service

type RealmService interface {
	Create(ctx context.Context, in *Realm, opts ...client.CallOption) (*status.Status, error)
	Get(ctx context.Context, in *RealmRequest, opts ...client.CallOption) (*Realm, error)
	Update(ctx context.Context, in *Realm, opts ...client.CallOption) (*status.Status, error)
	Delete(ctx context.Context, in *RealmRequest, opts ...client.CallOption) (*status.Status, error)
}

type realmService struct {
	c    client.Client
	name string
}

func NewRealmService(name string, c client.Client) RealmService {
	return &realmService{
		c:    c,
		name: name,
	}
}

func (c *realmService) Create(ctx context.Context, in *Realm, opts ...client.CallOption) (*status.Status, error) {
	req := c.c.NewRequest(c.name, "RealmService.Create", in)
	out := new(status.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmService) Get(ctx context.Context, in *RealmRequest, opts ...client.CallOption) (*Realm, error) {
	req := c.c.NewRequest(c.name, "RealmService.Get", in)
	out := new(Realm)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmService) Update(ctx context.Context, in *Realm, opts ...client.CallOption) (*status.Status, error) {
	req := c.c.NewRequest(c.name, "RealmService.Update", in)
	out := new(status.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmService) Delete(ctx context.Context, in *RealmRequest, opts ...client.CallOption) (*status.Status, error) {
	req := c.c.NewRequest(c.name, "RealmService.Delete", in)
	out := new(status.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RealmService service

type RealmServiceHandler interface {
	Create(context.Context, *Realm, *status.Status) error
	Get(context.Context, *RealmRequest, *Realm) error
	Update(context.Context, *Realm, *status.Status) error
	Delete(context.Context, *RealmRequest, *status.Status) error
}

func RegisterRealmServiceHandler(s server.Server, hdlr RealmServiceHandler, opts ...server.HandlerOption) error {
	type realmService interface {
		Create(ctx context.Context, in *Realm, out *status.Status) error
		Get(ctx context.Context, in *RealmRequest, out *Realm) error
		Update(ctx context.Context, in *Realm, out *status.Status) error
		Delete(ctx context.Context, in *RealmRequest, out *status.Status) error
	}
	type RealmService struct {
		realmService
	}
	h := &realmServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RealmService.Create",
		Path:    []string{"/api/v1/realms"},
		Method:  []string{"POST"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RealmService.Get",
		Path:    []string{"/api/v1/realms/{name}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RealmService.Update",
		Path:    []string{"/api/v1/realms/{name}"},
		Method:  []string{"PUT"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RealmService.Delete",
		Path:    []string{"/api/v1/realms/{name}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&RealmService{h}, opts...))
}

type realmServiceHandler struct {
	RealmServiceHandler
}

func (h *realmServiceHandler) Create(ctx context.Context, in *Realm, out *status.Status) error {
	return h.RealmServiceHandler.Create(ctx, in, out)
}

func (h *realmServiceHandler) Get(ctx context.Context, in *RealmRequest, out *Realm) error {
	return h.RealmServiceHandler.Get(ctx, in, out)
}

func (h *realmServiceHandler) Update(ctx context.Context, in *Realm, out *status.Status) error {
	return h.RealmServiceHandler.Update(ctx, in, out)
}

func (h *realmServiceHandler) Delete(ctx context.Context, in *RealmRequest, out *status.Status) error {
	return h.RealmServiceHandler.Delete(ctx, in, out)
}
