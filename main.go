package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	grpc_auth "github.com/johnbellone/persona-service/internal/auth"
	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	"github.com/johnbellone/persona-service/internal/service"

	log "github.com/sirupsen/logrus"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var (
	environment string
	serverPort  uint
	tlsCertFile string
	tlsKeyFile  string
	verbose     bool
	debug       bool
	databaseDsn string
	sentryDsn   string
	certificate tls.Certificate
)

func init() {
	flag.UintVar(&serverPort, "port", 9090, "Set the server port.")
	flag.BoolVar(&verbose, "verbose", false, "Turn on verbose logging.")
	flag.BoolVar(&debug, "debug", false, "Turn on debug logging.")
	flag.StringVar(&tlsCertFile, "tls-cert", "server.crt", "Set the path to TLS certificate.")
	flag.StringVar(&tlsKeyFile, "tls-key", "server.key", "Set the path to TLS key.")
	flag.StringVar(&databaseDsn, "database-dsn", "", "Set the database connection string.")
	flag.StringVar(&sentryDsn, "sentry-dsn", "", "Set the Sentry connection string.")
	flag.StringVar(&environment, "environment", "development", "Set the environment name.")
}

func main() {
	flag.Parse()

	switch {
	case debug:
		log.SetLevel(log.TraceLevel)
	case verbose:
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	switch {
	case environment == "production":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatal(err)
	}

	// Read the TLS server certificate and key from the command-line.
	cert, err := ioutil.ReadFile(tlsCertFile)
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile(tlsKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create the server certificate from the TLS key pairs.
	certificate, err = tls.X509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}

	entry := log.NewEntry(log.StandardLogger())
	grpc_logrus.ReplaceGrpcLogger(entry)

	s := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&certificate)),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(entry),
			grpc_auth.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(entry),
			grpc_recovery.StreamServerInterceptor(),
		),
	)

	// Add all of the service handlers to the gRPC server.
	pb.RegisterPersonServiceServer(s, new(service.PersonHandler))
	pb.RegisterAuthServiceServer(s, new(service.AuthHandler))
	pb.RegisterGroupServiceServer(s, new(service.GroupHandler))
	pb.RegisterRealmServiceServer(s, new(service.RealmHandler))
	pb.RegisterRoleServiceServer(s, new(service.RoleHandler))
	pb.RegisterUserServiceServer(s, new(service.UserHandler))

	// Turn on API reflection if we are in debug mode.
	if debug {
		reflection.Register(s)
	}

	log.Infof("ready on https://%v", ln.Addr())
	if err = s.Serve(ln); err != nil {
		log.Fatalf("Error %v", err)
	}
}
