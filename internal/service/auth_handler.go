package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	"github.com/johnbellone/persona-service/internal/server"
	"github.com/pascaldekloe/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	config *server.Config
}

func NewAuthHandler(c *server.Config) *AuthHandler {
	return &AuthHandler{config: c}
}

func (h *AuthHandler) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	if req.Login == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid login or password")
	}

	// TODO: Check request information against the user store.

	c := jwt.Claims{
		Registered: jwt.Registered{
			Issuer:    "persona",
			Subject:   req.Login,
			Audiences: []string{"c"},
			Expires:   jwt.NewNumericTime(time.Now().Add(time.Hour * 2)),
			NotBefore: jwt.NewNumericTime(time.Now().Add(time.Second * -1)),
			Issued:    jwt.NewNumericTime(time.Now()),
			ID:        "x",
		},
		Set: map[string]interface{}{
			"usr": req.Login,
		},
	}

	token, err := c.RSASign("RS256", h.config.JwtPrivateKey)
	if err != nil {
		return nil, err
	}

	return &pb.AuthResponse{Token: string(token)}, nil
}

func (h *AuthHandler) Validate(ctx context.Context, req *wrappers.StringValue) (*pb.AuthResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (h *AuthHandler) Refresh(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}
