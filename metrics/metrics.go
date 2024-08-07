package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics is the interface for the EigenMetrics server
// it only wraps 2 of the 6 methods required by the spec
// (https://docs.eigenlayer.xyz/eigenlayer/avs-guides/spec/metrics/metrics-prom-spec)
// the others are implemented by the economics and rpc_calls collectors,
// and need to be registered with the metrics server prometheus registry
type Metrics interface {
	AddFeeEarnedTotal(amount float64, token string)
	SetPerformanceScore(score float64)
	Start(ctx context.Context, reg prometheus.Gatherer) <-chan error
}
