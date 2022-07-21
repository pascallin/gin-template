package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	labels             map[string]string
	WebsocketConnCount prometheus.Gauge
)

const prefix = "service_instance_"

func setupPrometheusCollectors() {
	// Setup the Prometheus collectors.
	WebsocketConnCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name:        prefix + "websocket_conn_count",
		Help:        "The number of live room count.",
		ConstLabels: labels,
	})
}
