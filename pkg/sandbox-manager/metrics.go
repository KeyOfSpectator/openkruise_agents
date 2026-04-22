package sandbox_manager // Shared with api.go

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	// SandboxCreationLatency tracks the time from request to return
	SandboxCreationLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "sandbox_creation_latency_ms",
			Help:    "Latency of sandbox creation in milliseconds",
			Buckets: prometheus.ExponentialBuckets(10, 2, 10), // 10ms to ~10s
		},
	)

	// SandboxCreationResponses tracks total requests and failures
	SandboxCreationResponses = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sandbox_creation_responses",
			Help: "Total number of sandbox creation requests and their results",
		},
		[]string{"result"}, // "success" or "failure"
	)

	// SandboxPauseLatency tracks the time of sandbox pause operations
	SandboxPauseLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "sandbox_pause_latency_ms",
			Help:    "Latency of sandbox pause operations in milliseconds",
			Buckets: prometheus.ExponentialBuckets(10, 2, 10),
		},
	)

	// SandboxPauseResponses tracks total pause requests and their results
	SandboxPauseResponses = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sandbox_pause_responses",
			Help: "Total number of sandbox pause requests and their results",
		},
		[]string{"result"},
	)

	// SandboxResumeLatency tracks the time of sandbox resume operations
	SandboxResumeLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "sandbox_resume_latency_ms",
			Help:    "Latency of sandbox resume operations in milliseconds",
			Buckets: prometheus.ExponentialBuckets(10, 2, 10),
		},
	)

	// SandboxResumeResponses tracks total resume requests and their results
	SandboxResumeResponses = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sandbox_resume_responses",
			Help: "Total number of sandbox resume requests and their results",
		},
		[]string{"result"},
	)

	// SandboxDeleteResponses tracks total delete requests and their results
	SandboxDeleteResponses = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sandbox_delete_responses",
			Help: "Total number of sandbox delete requests and their results",
		},
		[]string{"result"},
	)
)

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(SandboxCreationLatency, SandboxCreationResponses,
		SandboxPauseLatency, SandboxPauseResponses,
		SandboxResumeLatency, SandboxResumeResponses,
		SandboxDeleteResponses)
}
