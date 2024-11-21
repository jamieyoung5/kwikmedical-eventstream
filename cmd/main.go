package main

import (
	"github.com/jamieyoung5/kwikmedical-eventstream/pb"
	"github.com/jamieyoung5/kwikmedical-eventstream/pkg/eventstream"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func RunVercelApp(w http.ResponseWriter, r *http.Request) {
	go main()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("gRPC Server is running"))
}

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
