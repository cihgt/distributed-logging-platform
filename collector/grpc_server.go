package main

import (
	"context"
	logging "github.com/cihgt/distributed-logging-platform/proto"
	"log"
)

type server struct {
	logging.LogServiceServer
}

func (s *server) ReportHealth(ctx context.Context, req *logging.HealthReport) (*logging.HealthResponse, error) {
	healthCPU.WithLabelValues(req.Service).Set(req.CpuUsage)
	healthMem.WithLabelValues(req.Service).Set(req.MemoryUsage)
	healthGoroutines.WithLabelValues(req.Service).Set(float64(req.Goroutines))

	log.Printf("[health][%s] CPU: %.2f, Mem: %.2f, GoR: %d\n",
		req.Service, req.CpuUsage, req.MemoryUsage, req.Goroutines)

	return &logging.HealthResponse{Status: "ok"}, nil
}
