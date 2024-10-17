package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/wowholasa/SEC-MH-2-GO/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type hospital struct {
	pb.UnimplementedAggregationSendingServiceServer
	receivedShares []int64
}

func (h *hospital) SendAggregation(ctx context.Context, msg *pb.Aggregation) (*pb.Acknowledge, error) {
	fmt.Println("Received aggregation from patient")
	fmt.Println("Aggregation: ", msg.Aggregation)

	h.receivedShares = append(h.receivedShares, msg.Aggregation)
	return &pb.Acknowledge{Message: "Received Aggregation, and added it to list."}, nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, fmt.Errorf("could not load server key pair: %s", err)
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func (h *hospital) SumAggregations() int64 {
	var sum int64
	for _, share := range h.receivedShares {
		sum += share
	}
	return sum
}

func main() {
	listen, err := net.Listen("tcp", ":5454") // Listen on port 5454
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAggregationSendingServiceServer(grpcServer, &hospital{})

	fmt.Println("Hospital server running on port :5454")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
