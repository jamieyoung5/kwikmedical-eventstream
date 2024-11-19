package main

import (
	"github.com/jamieyoung5/kwikmedical-eventstream/pb"
	"github.com/jamieyoung5/kwikmedical-eventstream/pkg/eventstream"
	"google.golang.org/grpc"
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

	pb.RegisterEventStreamV1Server(grpcServer, eventStreamServer)

	log.Println("EventStreamV1 server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}