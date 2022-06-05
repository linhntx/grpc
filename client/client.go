package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/linhntx/gprc/grpc_chatserver"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Enter server IP:Port :::")
	reader := bufio.NewReader(os.Stdin)
	serverID, err := reader.ReadString('\n')

	if err != nil {
		log.Printf("Failed to read from console :: %v", err)

	}
	serverID = strings.Trim(serverID, "\r\n")

	log.Println("Connecting: " + serverID)

	//connect to grpc server
	con, err := grpc.Dial(serverID, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to gRPC server :: %v", err)
	}

	defer con.Close()

	//client ChatServer to create a stream
	client := grpc_chatserver.NewServicesClient(con)

	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect ChatService :: %v", err)
	}
	fmt.Println("stream", stream)

	// implement comunication with gRPC server
	// chat := clientHandle{stream: stream}
	

	//	blocker
	bl := make(chan bool)
	<-bl 
	
	
}

type clientHandle struct {
	stream 		grpc_chatserver.Services_ChatServiceClient
	clientname	string
}