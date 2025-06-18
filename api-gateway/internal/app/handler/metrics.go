package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func initPrometheus() (*prometheus.CounterVec, *prometheus.HistogramVec) {
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: []float64{0.1, 0.5, 1, 2, 5},
		},
		[]string{"method", "path"},
	)

	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)

	return httpRequestsTotal, httpRequestDuration
}

func prometheusMiddleware(cv *prometheus.CounterVec, hv *prometheus.HistogramVec) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()

		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		c.Next()

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		cv.WithLabelValues(
			c.Request.Method,
			path,
			http.StatusText(status),
		).Inc()

		hv.WithLabelValues(
			c.Request.Method,
			path,
		).Observe(duration)
	}
}
