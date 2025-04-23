package main

import (
	pb "distributed-logging-platform/proto"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var lastHealthReports = sync.Map{}

func startDebugREST() {
	http.HandleFunc("/debug/health", func(w http.ResponseWriter, r *http.Request) {
		reports := make(map[string]*pb.HealthReport)
		lastHealthReports.Range(func(key, value interface{}) bool {
			reports[key.(string)] = value.(*pb.HealthReport)
			return true
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reports)
	})

	log.Println("🧠 Debug API: http://localhost:9091/debug/health")
	http.ListenAndServe(":9091", nil)
}
