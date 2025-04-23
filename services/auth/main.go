package main

//
//import (
//	"context"
//	"log"
//	"time"
//
//	logging "distributed-logging-platform/proto"
//	"google.golang.org/grpc"
//)
//
//func main() {
//	conn, err := grpc.Dial("collector:50051", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("failed to connect: %v", err)
//	}
//	defer conn.Close()
//
//	client := logging.NewLogStreamerClient(conn)
//	stream, err := client.StreamLogs(context.Background())
//	if err != nil {
//		log.Fatalf("failed to start log stream: %v", err)
//	}
//
//	for {
//		entry := &logging.LogEntry{
//			ServiceName: "auth",
//			Level:       "info",
//			Message:     "User login successful",
//			Timestamp:   time.Now().Format(time.RFC3339),
//		}
//		if err := stream.Send(entry); err != nil {
//			log.Printf("failed to send log: %v", err)
//		}
//		time.Sleep(2 * time.Second)
//	}
//}
