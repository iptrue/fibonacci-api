package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type Metrics struct {
	RequestCount    *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
	ErrorCount      *prometheus.CounterVec
}

func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		RequestCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "request_count_total",
				Help:      "Total number of requests",
			},
			[]string{"method", "endpoint"},
		),
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "request_duration_seconds",
				Help:      "Request duration in seconds",
			},
			[]string{"method", "endpoint"},
		),
		ErrorCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "error_count_total",
				Help:      "Total number of errors",
			},
			[]string{"method", "endpoint", "error_code"},
		),
	}
}

func (m *Metrics) Register() {
	prometheus.MustRegister(m.RequestCount)
	prometheus.MustRegister(m.RequestDuration)
	prometheus.MustRegister(m.ErrorCount)
}

func ServeMetrics(port string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Printf("Prometheus metrics server running on port %s\n", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("Failed to start Prometheus metrics server: %v", err)
		}
	}()
}
