package main

import (
	"context"
	"flag"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	return
}

var (
	Environment string
	Port        uint
	TlsCertPath string
	TlsInsecure bool
	TlsKeyPath  string
	Verbose     bool
	Debug       bool
	DatabaseUrl string
)

func init() {
	flag.UintVar(&Port, "port", 9090, "Set the server port.")
	flag.BoolVar(&Verbose, "verbose", false, "Turn on verbose logging.")
	flag.BoolVar(&Debug, "debug", false, "Turn on debug logging.")
	flag.BoolVar(&TlsInsecure, "tls-insecure", false, "Enable insecure mode.")
	flag.StringVar(&TlsCertPath, "tls-cert-path", "", "Set the path to TLS certificate.")
	flag.StringVar(&TlsKeyPath, "tls-key-path", "", "Set the path to TLS key.")
	flag.StringVar(&DatabaseUrl, "database-url", "", "Set the database connection string.")
	flag.StringVar(&Environment, "environment", "development", "Set the environment name.")

	if Debug {
		log.SetLevel(log.TraceLevel)
	} else if Verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	if Environment == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
}

func main() {
	_, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		os.Exit(2)
	}

}
