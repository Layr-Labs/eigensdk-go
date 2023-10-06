package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

type NoopMetrics struct{}

// enforce that NoopMetrics implements Metrics interface
var _ Metrics = (*NoopMetrics)(nil)

func NewNoopMetrics() *NoopMetrics {
	return &NoopMetrics{}
}

func (m *NoopMetrics) AddFeeEarnedTotal(amount float64, token string) {}
func (m *NoopMetrics) SetPerformanceScore(score float64)              {}
func (m *NoopMetrics) Start(ctx context.Context, reg prometheus.Gatherer) <-chan error {
	return nil
}
