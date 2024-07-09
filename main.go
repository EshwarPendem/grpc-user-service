package main

import (
	"log"
	"net"

	"github.com/eshwarpendem/grpc-user-service/proto"
	"github.com/eshwarpendem/grpc-user-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8081"

func main() {
	// Create a new gRPC server and register the User service
	server := grpc.NewServer()
	userServiceServer := service.NewUserServiceServer()
	proto.RegisterUserServiceServer(server, userServiceServer)

	//This line is for triggering the endpoints using terminal tools like grpcurl.
	reflection.Register(server)
	
	// Start the server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start listener %v", err)
	}
	log.Printf("Server started at %v", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
