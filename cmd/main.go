package main

import (
	"google.golang.org/grpc"
	"kwikmedical-eventstream/gen"
	"kwikmedical-eventstream/pkg/eventstream"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	eventStreamServer := eventstream.NewServer()

	gen.RegisterEventStreamV1Server(grpcServer, eventStreamServer)

	log.Println("EventStreamV1 server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
