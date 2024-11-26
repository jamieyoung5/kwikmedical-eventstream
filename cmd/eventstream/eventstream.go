package eventstream_server

import (
	"fmt"
	"github.com/jamieyoung5/kwikmedical-eventstream/pb"
	"github.com/jamieyoung5/kwikmedical-eventstream/pkg/eventstream"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const portEnvVar = "PORT"

func Start() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	port := fmt.Sprintf(":%s", os.Getenv(portEnvVar))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	eventStreamServer := eventstream.NewServer(logger)

	pb.RegisterEventStreamV1Server(grpcServer, eventStreamServer)

	log.Println("EventStreamV1 server is running on port " + os.Getenv(portEnvVar) + "...")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}
