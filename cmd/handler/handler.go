package handler

import (
	eventstream_server "github.com/jamieyoung5/kwikmedical-eventstream/cmd/eventstream"
	"net/http"
)

func RunVercelApp(w http.ResponseWriter, r *http.Request) {
	go eventstream_server.Start()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("gRPC Server is running"))
}
