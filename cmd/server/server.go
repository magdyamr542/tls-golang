package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load server certificate and private key
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		panic(err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}

	// Create a certificate pool and add the CA certificate to it
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	// Create an HTTP handler that returns a simple message
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS!")
	})

	// Create a server with the TLS configuration and handler
	server := &http.Server{
		Addr:      ":1234",
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	// Start the server
	log.Printf("server is up on https://localhost:1234\n")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
