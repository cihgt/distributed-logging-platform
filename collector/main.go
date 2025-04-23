package main

import (
	logging "github.com/cihgt/distributed-logging-platform/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Prometheus metrics available at :9090/metrics")
		log.Fatal(http.ListenAndServe(":9090", nil))
	}()

	// запуска REST API сервиса
	go startDebugREST()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	logging.RegisterLogServiceServer(s, &server{})

	log.Println("Log Collector gRPC Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
