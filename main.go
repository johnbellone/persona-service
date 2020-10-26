package main

import (
	"context"
	"crypto/rsa"
	"flag"
	"fmt"
	"net"
	"time"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	"github.com/johnbellone/persona-service/internal/server"
	"github.com/johnbellone/persona-service/internal/service"
	"github.com/pascaldekloe/jwt"

	log "github.com/sirupsen/logrus"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var config *server.Config

func init() {
	config = new(server.Config)

	flag.UintVar(&config.Port, "port", 9090, "Set the server port.")
	flag.BoolVar(&config.Verbose, "verbose", false, "Turn on verbose logging.")
	flag.BoolVar(&config.Debug, "debug", false, "Turn on debug logging.")
	flag.StringVar(&config.TlsCertFile, "tls-cert", "server.crt", "Set the path to TLS certificate.")
	flag.StringVar(&config.TlsKeyFile, "tls-key", "server.key", "Set the path to TLS key.")
	flag.StringVar(&config.DatabaseDSN, "database-dsn", "", "Set the database connection string.")
	flag.StringVar(&config.SentryDSN, "sentry-dsn", "", "Set the Sentry connection string.")
	flag.StringVar(&config.Environment, "environment", "development", "Set the environment name.")
}

func JwtFromContext(ctx context.Context) (context.Context, error) {
	if bearer, err := grpc_auth.AuthFromMD(ctx, "bearer"); err == nil {
		claims, err := jwt.RSACheck([]byte(bearer), config.JwtPrivateKey.Public().(*rsa.PublicKey))
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "Bad token format")
		}

		if !claims.Valid(time.Now()) {
			return nil, status.Error(codes.Unauthenticated, "Token expired")
		}

		grpc_ctxtags.Extract(ctx).Set("auth.id", claims.ID)
		grpc_ctxtags.Extract(ctx).Set("auth.iss", claims.Issuer)
		grpc_ctxtags.Extract(ctx).Set("auth.sub", claims.Subject)
		grpc_ctxtags.Extract(ctx).Set("auth.exp", claims.Expires)
		log.Infof("%v", ctx)
		return context.WithValue(ctx, "jwt", claims.Raw), nil
	}
	return ctx, nil
}

func main() {
	flag.Parse()

	if err := config.Setup(); err != nil {
		log.Fatal(err)
	}

	switch {
	case config.Debug:
		log.SetLevel(log.TraceLevel)
	case config.Verbose:
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	switch {
	case config.Environment == "production":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	entry := log.NewEntry(log.StandardLogger())
	grpc_logrus.ReplaceGrpcLogger(entry)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&config.Certificate)),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(entry),
			grpc_auth.UnaryServerInterceptor(JwtFromContext),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(entry),
			grpc_auth.StreamServerInterceptor(JwtFromContext),
			grpc_recovery.StreamServerInterceptor(),
		),
	)

	// Add all of the service handlers to the gRPC server.
	pb.RegisterPersonServiceServer(s, service.NewPersonHandler(config))
	pb.RegisterAuthServiceServer(s, service.NewAuthHandler(config))
	pb.RegisterGroupServiceServer(s, service.NewGroupHandler(config))
	pb.RegisterRealmServiceServer(s, service.NewRealmHandler(config))
	pb.RegisterRoleServiceServer(s, service.NewRoleHandler(config))
	pb.RegisterUserServiceServer(s, service.NewUserHandler(config))

	// Turn on API reflection if we are in debug mode.
	if config.Debug {
		reflection.Register(s)
	}

	log.Infof("ready on https://%v", ln.Addr())
	if err = s.Serve(ln); err != nil {
		log.Fatalf("Error %v", err)
	}
}
