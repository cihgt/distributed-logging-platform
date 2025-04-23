package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	healthCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_cpu_usage",
			Help: "CPU usage by service",
		},
		[]string{"service"},
	)

	healthMem = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_memory_usage",
			Help: "Memory usage by service",
		},
		[]string{"service"},
	)

	healthGoroutines = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_goroutines",
			Help: "Goroutines count",
		},
		[]string{"service"},
	)
)

func init() {
	prometheus.MustRegister(healthCPU, healthMem, healthGoroutines)
}
