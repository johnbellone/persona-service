package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	pb "github.com/johnbellone/persona-service/internal/gen/persona/api/v1"
	"github.com/johnbellone/persona-service/internal/service"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var (
	Environment string
	Port        uint
	TlsCertFile string
	TlsKeyFile  string
	Verbose     bool
	Debug       bool
	DatabaseUrl string
	Certificate tls.Certificate
)

func init() {
	flag.UintVar(&Port, "port", 9090, "Set the server port.")
	flag.BoolVar(&Verbose, "verbose", false, "Turn on verbose logging.")
	flag.BoolVar(&Debug, "debug", false, "Turn on debug logging.")
	flag.StringVar(&TlsCertFile, "tls-cert", "server.crt", "Set the path to TLS certificate.")
	flag.StringVar(&TlsKeyFile, "tls-key", "server.key", "Set the path to TLS key.")
	flag.StringVar(&DatabaseUrl, "database-url", "", "Set the database connection string.")
	flag.StringVar(&Environment, "environment", "development", "Set the environment name.")
}

func main() {
	flag.Parse()

	switch {
	case Debug:
		log.SetLevel(log.TraceLevel)
	case Verbose:
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	switch {
	case Environment == "production":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		log.Fatal(err)
	}

	// Read the TLS server certificate and key from the command-line.
	cert, err := ioutil.ReadFile(TlsCertFile)
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile(TlsKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create the server certificate from the TLS key pairs.
	Certificate, err = tls.X509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}

	entry := logrus.NewEntry(log.StandardLogger())
	grpc_logrus.ReplaceGrpcLogger(entry)

	s := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&Certificate)),
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(entry),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(entry),
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
	if Debug {
		reflection.Register(s)
	}

	log.Infof("ready to rock on https://%v", ln.Addr())
	if err = s.Serve(ln); err != nil {
		log.Fatalf("Error %v", err)
	}
}
