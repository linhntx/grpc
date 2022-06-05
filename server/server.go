package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	address := "0.0.0.0:50051"

	//init listener
	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	log.Printf("Server is listening on %v....", address)

	//gRPC server instance
	grpcServer := grpc.NewServer()

	//grpc listen and serve
	err = grpcServer.Serve(listen)
	
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
}