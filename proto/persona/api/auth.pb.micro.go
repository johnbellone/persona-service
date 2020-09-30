// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: persona/api/auth.proto

package persona_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	status "google.golang.org/genproto/googleapis/rpc/status"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
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

// Api Endpoints for AuthService service

func NewAuthServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuthService service

type AuthService interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
	Refresh(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
	Validate(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*status.Status, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Authenticate(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "AuthService.Authenticate", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Refresh(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "AuthService.Refresh", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Validate(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*status.Status, error) {
	req := c.c.NewRequest(c.name, "AuthService.Validate", in)
	out := new(status.Status)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceHandler interface {
	Authenticate(context.Context, *AuthRequest, *AuthResponse) error
	Refresh(context.Context, *AuthRequest, *AuthResponse) error
	Validate(context.Context, *wrappers.StringValue, *status.Status) error
}

func RegisterAuthServiceHandler(s server.Server, hdlr AuthServiceHandler, opts ...server.HandlerOption) error {
	type authService interface {
		Authenticate(ctx context.Context, in *AuthRequest, out *AuthResponse) error
		Refresh(ctx context.Context, in *AuthRequest, out *AuthResponse) error
		Validate(ctx context.Context, in *wrappers.StringValue, out *status.Status) error
	}
	type AuthService struct {
		authService
	}
	h := &authServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthService{h}, opts...))
}

type authServiceHandler struct {
	AuthServiceHandler
}

func (h *authServiceHandler) Authenticate(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.AuthServiceHandler.Authenticate(ctx, in, out)
}

func (h *authServiceHandler) Refresh(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.AuthServiceHandler.Refresh(ctx, in, out)
}

func (h *authServiceHandler) Validate(ctx context.Context, in *wrappers.StringValue, out *status.Status) error {
	return h.AuthServiceHandler.Validate(ctx, in, out)
}
