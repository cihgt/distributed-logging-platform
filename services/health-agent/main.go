package main

import (
	"context"
	pb "github.com/cihgt/distributed-logging-platform/proto"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	host := os.Getenv("LOG_SERVICE_HOST")
	if host == "" {
		host = "collector:50051"
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("не смог подрубиться к gRPC: %v", err)
	}
	defer conn.Close()

	client := pb.NewLogServiceClient(conn)

	for {
		_, err := client.ReportHealth(context.Background(), &pb.HealthReport{
			Service:     "auth",
			CpuUsage:    14.5,
			MemoryUsage: 27.8,
			Goroutines:  32,
			Uptime:      uint64(time.Now().Unix()),
		})
		if err != nil {
			log.Printf("Ошибка при отправке health: %v", err)
		} else {
			log.Println("Health метрика отправлена")
		}
		time.Sleep(5 * time.Second)
	}
}
