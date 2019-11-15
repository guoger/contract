package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/hyperledger/fabric-protos-go/peer"
	"google.golang.org/grpc"
)

var server = flag.String("server", "localhost:7070", "chaincode server address")

func main() {
	flag.Parse()
	log.Println("Dialing", *server)
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial %s: %v", *server, err)
	}
	defer conn.Close()

	client, err := peer.NewChaincodeClient(conn).Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to grpc server: %v", err)
	}

	msg, err := client.Recv()
	if err != nil {
		log.Fatalf("Failed to receive chaincode message: %v", err)
	}
	log.Printf("Received: %v", msg)

	err = client.Send(&peer.ChaincodeMessage{Type: peer.ChaincodeMessage_REGISTERED})
	if err != nil {
		log.Fatalf("Failed to send REGISTERED: %v", err)
	}

	err = client.Send(&peer.ChaincodeMessage{Type: peer.ChaincodeMessage_READY})
	if err != nil {
		log.Fatalf("Failed to send chaincode message: %v", err)
	}

	err = client.Send(&peer.ChaincodeMessage{Type: peer.ChaincodeMessage_INIT})
	if err != nil {
		log.Fatalf("Failed to send chaincode message: %v", err)
	}

	msg, err = client.Recv()
	if err != nil {
		log.Fatalf("Failed to receive chaincode message: %v", err)
	}
	log.Printf("Received: %v", msg)

	err = client.Send(&peer.ChaincodeMessage{Type: peer.ChaincodeMessage_TRANSACTION})
	if err != nil {
		log.Fatalf("Failed to send chaincode message: %v", err)
	}

	msg, err = client.Recv()
	if err != nil {
		log.Fatalf("Failed to receive chaincode message: %v", err)
	}
	log.Printf("Received: %v", msg)

	time.Sleep(4 * time.Second)

	err = client.Send(&peer.ChaincodeMessage{Type: peer.ChaincodeMessage_TRANSACTION})
	if err != nil {
		log.Fatalf("Failed to send chaincode message: %v", err)
	}

	msg, err = client.Recv()
	if err != nil {
		log.Fatalf("Failed to receive chaincode message: %v", err)
	}
	log.Printf("Received: %v", msg)

	time.Sleep(6 * time.Second)

	err = client.Send(&peer.ChaincodeMessage{Type: peer.ChaincodeMessage_TRANSACTION})
	if err != nil {
		log.Fatalf("Failed to send chaincode message: %v", err)
	}

	msg, err = client.Recv()
	if err != nil {
		log.Fatalf("Failed to receive chaincode message: %v", err)
	}
	log.Printf("Received: %v", msg)

	msg, err = client.Recv()
	if err != nil {
		log.Fatalf("Failed to receive chaincode message: %v", err)
	}
	log.Printf("Received: %v", msg)

	os.Exit(0)
}
