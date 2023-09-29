package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	AddFeeEarnedTotal(amount float64, token string)
	SetPerformanceScore(score float64)
	ObserveRPCRequestDurationSeconds(duration float64, method, client, version string)
	AddRPCRequestTotal(method, client, version string)
	Start(ctx context.Context, reg prometheus.Gatherer) <-chan error
}
