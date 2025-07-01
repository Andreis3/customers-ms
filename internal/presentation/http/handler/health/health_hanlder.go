package health

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/andreis3/customers-ms/internal/presentation/http/transport"
)

type HealthCheckResponse struct {
	Status    string            `json:"status"`
	Timestamp string            `json:"timestamp"`
	System    SystemInformation `json:"system"`
	Component ComponentInfo     `json:"component"`
}
type SystemInformation struct {
	Version          string `json:"version"`
	GoroutinesCount  int    `json:"goroutines_count"`
	TotalAllocBytes  uint64 `json:"total_alloc_bytes"`
	HeapObjectsCount uint64 `json:"heap_objects_count"`
	AllocBytes       uint64 `json:"alloc_bytes"`
	HealAllocBytes   uint64 `json:"heal_alloc_bytes"`
}
type ComponentInfo struct {
	ServiceName string `json:"service_name"`
}

func HealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		systemInfo := getSystemInformation()
		response := HealthCheckResponse{
			Status:    http.StatusText(http.StatusOK),
			Timestamp: time.Now().Format(time.RFC3339),
			System:    systemInfo,
			Component: ComponentInfo{
				ServiceName: "customers-ms",
			},
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to serialize JSON response")
			return
		}
		w.Header().Set(transport.ContentType, transport.ApplicationJSON)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonResponse)
	})
}

func getSystemInformation() SystemInformation {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return SystemInformation{
		Version:          runtime.Version(),
		GoroutinesCount:  runtime.NumGoroutine(),
		TotalAllocBytes:  memStats.TotalAlloc,
		HeapObjectsCount: memStats.HeapObjects,
		AllocBytes:       memStats.Alloc,
		HealAllocBytes:   memStats.HeapAlloc,
	}
}
