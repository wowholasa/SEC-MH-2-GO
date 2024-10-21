package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	pb "github.com/wowholasa/SEC-MH-2-GO/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Struct to define a patient
type patient struct {
	pb.UnimplementedPatientShareSendingServiceServer
	patientID            int64
	addressPort          string
	otherPatientAdresses map[int64]string
	serverAddress        string
	initialShares        []int64
	receivedShares       []int64
	input                int64
}

// Function to calculate the initial shares of the patient's private input
func (p *patient) calculateInitialShares() {
	share1 := rand.Int63()
	share2 := rand.Int63()

	share3 := p.input - share1 - share2

	p.initialShares = []int64{share1, share2, share3}
	return
}

// Function to aggregate the shares received from the other patients
func (p *patient) aggregateShares() int64 {
	var aggregation int64
	for _, share := range p.receivedShares {
		aggregation += share
	}
	return aggregation
}

// Function to load the TLS credentials for the patient
// Load's the patient's certificate and private key and defines the CAs which the patient trusts when connecting to the server and other patients
func loadTLSCredentials(patientID int64) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	CACert, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(CACert) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	certFile := fmt.Sprintf("cert/client%d-cert.pem", patientID)
	keyFile := fmt.Sprintf("cert/client%d-key.pem", patientID)
	clientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load client key pair: %s", err)
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

// Function to start the patient server with TLS credentials
func (p *patient) startPatientServer(wg *sync.WaitGroup) {
	defer wg.Done()

	listen, err := net.Listen("tcp", p.addressPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tlsCredentials, err := loadTLSCredentials(p.patientID)
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	pb.RegisterPatientShareSendingServiceServer(grpcServer, p)

	log.Printf("Patient %d server listening on %s\n", p.patientID, p.addressPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// Function to receive the share from another patient
// If the patient has received all 3 shares, it aggregates them and sends the aggregation to the hospital
func (p *patient) SendShare(ctx context.Context, msg *pb.Share) (*pb.Acknowledge, error) {
	log.Printf("Received share from patient %d\n", msg.ShareOfSecret)
	p.receivedShares = append(p.receivedShares, msg.ShareOfSecret)

	if len(p.receivedShares) == 3 {
		aggregation := p.aggregateShares()
		p.sendAggregationToHospital(context.Background(), aggregation)
	}

	return &pb.Acknowledge{Message: "Received Share, and added it to list."}, nil
}

// Function to send the share to another patient
func (p *patient) sendShareToOtherPatient(ctx context.Context, share int64, otherPatientID int64) {
	tlsCredentials, err := loadTLSCredentials(p.patientID)
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	connection, err := grpc.Dial(p.otherPatientAdresses[otherPatientID], grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer connection.Close()

	client := pb.NewPatientShareSendingServiceClient(connection)
	ack, err := client.SendShare(ctx, &pb.Share{ShareOfSecret: share})
	if err != nil {
		log.Fatalf("Failed to send share: %v", err)
	}

	log.Printf("Received Acknowledgement from Client %d: %s\n", otherPatientID, ack.Message)
}

// Function to send the aggregation to the hospital
func (p *patient) sendAggregationToHospital(ctx context.Context, aggregation int64) {
	tlsCredentials, err := loadTLSCredentials(p.patientID)
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	connection, err := grpc.Dial(p.serverAddress, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer connection.Close()

	client := pb.NewAggregationSendingServiceClient(connection)
	ack, err := client.SendAggregation(ctx, &pb.Aggregation{Aggregation: aggregation})
	if err != nil {
		log.Fatalf("Client %d failed to send aggregation: %v", p.patientID, err)
	}

	log.Printf("Received Acknowledgement from Hospital: %s\n", ack.Message)
}

// Main function to start the patient
func main() {
	// Parse the command line arguments
	patientID := flag.Int("id", -1, "Patient ID")
	input := flag.Int64("input", -1, "Input value")
	flag.Parse()

	port := 5455 + int64(*patientID)

	// Define the addresses of the other patients
	otherPatients := map[int64]string{
		0: "localhost:5455",
		1: "localhost:5456",
		2: "localhost:5457",
	}

	// Remove the patient's own address from the list of other patients
	delete(otherPatients, int64(*patientID))

	// Create the patient struct
	patient := &patient{
		patientID:            int64(*patientID),
		addressPort:          fmt.Sprintf("localhost:%d", port),
		otherPatientAdresses: otherPatients,
		serverAddress:        "localhost:5454",
		initialShares:        []int64{},
		receivedShares:       []int64{},
		input:                *input,
	}

	// Calculate the initial shares of the patient's private input
	patient.calculateInitialShares()

	// Start the patient server with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go patient.startPatientServer(&wg)

	// Wait for the patient server to start and give time for other clients to start
	time.Sleep(10 * time.Second)

	// Send the shares to the other patients
	// Preferably, this would be done in a loop, but I couldn't get it to work
	if patient.patientID == 0 {
		patient.receivedShares = append(patient.receivedShares, patient.initialShares[0])
		time.Sleep(10 * time.Second) // This is a way to make sure that the other patients are ready to receive the shares
		patient.sendShareToOtherPatient(context.Background(), patient.initialShares[1], 1)
		patient.sendShareToOtherPatient(context.Background(), patient.initialShares[2], 2)
	} else if patient.patientID == 1 {
		patient.receivedShares = append(patient.receivedShares, patient.initialShares[1])
		time.Sleep(10 * time.Second) // This is a way to make sure that the other patients are ready to receive the shares
		patient.sendShareToOtherPatient(context.Background(), patient.initialShares[0], 0)
		patient.sendShareToOtherPatient(context.Background(), patient.initialShares[2], 2)
	} else if patient.patientID == 2 {
		patient.receivedShares = append(patient.receivedShares, patient.initialShares[2])
		patient.sendShareToOtherPatient(context.Background(), patient.initialShares[0], 0)
		patient.sendShareToOtherPatient(context.Background(), patient.initialShares[1], 1)
	}

	//// I just really wish I could get this to work
	// patient.receivedShares = append(patient.receivedShares, patient.initialShares[int(*patientID)])
	// for i := 0; i < 3; i++ {
	// 	if i == int(*patientID) {
	// 		continue
	// 	}
	// 	patient.sendShareToOtherPatient(context.Background(), patient.initialShares[i], int64(i))
	// }

	// if len(patient.receivedShares) == 3 {
	// 	aggregation := patient.aggregateShares()
	// 	patient.sendAggregationToHospital(context.Background(), aggregation)
	// }

	wg.Wait()
}
