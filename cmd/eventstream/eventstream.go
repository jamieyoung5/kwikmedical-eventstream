package eventstream_server

import (
	"fmt"
	"github.com/jamieyoung5/kwikmedical-eventstream/pb"
	"github.com/jamieyoung5/kwikmedical-eventstream/pkg/eventstream"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func Start() {
	var (
		err error
		lis net.Listener
	)

	for err != nil {
		port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

		lis, err = net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		
		time.Sleep(5 * time.Second)
	}

	grpcServer := grpc.NewServer()
	eventStreamServer := eventstream.NewServer()

	pb.RegisterEventStreamV1Server(grpcServer, eventStreamServer)

	log.Println("EventStreamV1 server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
