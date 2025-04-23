package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	logging "github.com/cihgt/distributed-logging-platform/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	// Метрики
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Prometheus metrics available at :9090/metrics")
		log.Fatal(http.ListenAndServe(":9090", nil))
	}()

	// REST API
	go startDebugREST()

	// gRPC сервер
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	logging.RegisterLogServiceServer(s, &server{})

	go func() {
		log.Println("Log Collector gRPC Server started at :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 👇 gRPC клиент и вызов ReportHealth
	go func() {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("failed to connect to gRPC server: %v", err)
		}
		defer conn.Close()

		client := logging.NewLogServiceClient(conn)

		for {
			_, err := client.ReportHealth(context.Background(), &logging.HealthReport{
				Service: "health-agent",
			})
			if err != nil {
				log.Println("[health-agent] Failed to report health:", err)
			} else {
				log.Println("[health-agent] Health reported successfully")
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// Ждём навсегда
	select {}
}
