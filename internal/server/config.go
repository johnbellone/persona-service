package server

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

type Config struct {
	Environment   string
	Port          uint
	Verbose       bool
	Debug         bool
	DatabaseDSN   string
	SentryDSN     string
	Certificate   tls.Certificate
	TlsCertFile   string
	TlsKeyFile    string
	JwtPrivateKey *rsa.PrivateKey
}

func (c *Config) Setup() error {
	// Read the TLS server certificate and key from the command-line.
	cert, err := ioutil.ReadFile(c.TlsCertFile)
	if err != nil {
		return err
	}

	key, err := ioutil.ReadFile(c.TlsKeyFile)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(key)
	if block == nil {
		return errors.New("invalid pem format")
	}

	c.JwtPrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	// Create the server certificate from the TLS key pairs.
	c.Certificate, err = tls.X509KeyPair(cert, key)
	if err != nil {
		return err
	}

	return nil
}
