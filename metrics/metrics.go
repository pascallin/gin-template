package metrics

import (
	"sync"
	"time"

	"github.com/pascallin/gin-template/model"
	"github.com/pascallin/gin-template/ws"
)

const systemMetricsPollingInterval = 10 * time.Second

// CollectedMetrics stores different collected + timestamped values.
type CollectedMetrics struct {
	m sync.Mutex `json:"-"`

	model.Status
}

// Metrics is the shared Metrics instance.
var metrics *CollectedMetrics

func Start() {
	setupPrometheusCollectors()

	metrics = new(CollectedMetrics)

	go func() {
		for range time.Tick(systemMetricsPollingInterval) {
			handlePolling()
		}
	}()
}

func handlePolling() {
	metrics.m.Lock()
	defer metrics.m.Unlock()

	collectSystem()
}

func collectSystem() {
	WebsocketConnCount.Set(float64(ws.GetWebsocketConnectionCount()))
}

// GetMetrics will return the collected metrics.
func GetMetrics() *CollectedMetrics {
	return metrics
}
