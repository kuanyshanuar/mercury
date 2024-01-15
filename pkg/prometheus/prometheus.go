package prometheus

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics - represents metrics
type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
}

// Metrics struct
type metrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

// NewMetrics - creates metrics in prometheus.
func NewMetrics(address string, serviceName string) (Metrics, error) {
	var (
		prometheusMetrics metrics
	)

	prometheusMetrics.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: serviceName + "_hits_total",
		Help: "",
	})

	if err := prometheus.Register(prometheusMetrics.HitsTotal); err != nil {
		return nil, err
	}

	prometheusMetrics.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: serviceName + "_hits",
			Help: "The total amount of requests to the service",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(prometheusMetrics.Hits); err != nil {
		return nil, err
	}

	prometheusMetrics.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: serviceName + "_times",
			Help: "",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(prometheusMetrics.Times); err != nil {
		return nil, err
	}

	//go func() {
	//	router := gin.New()
	//	router.GET("/metrics", helpers.PrometheusHandler())
	//	log.Printf("Metrics server is running on port: %s", address)
	//	if err := router.Run(address); err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	return &prometheusMetrics, nil
}

func (m *metrics) IncHits(status int, method, path string) {
	m.HitsTotal.Inc()
	m.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

func (m *metrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	m.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}
