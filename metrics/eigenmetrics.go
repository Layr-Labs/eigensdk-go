// Package metrics implements the avs node prometheus metrics spec: https://eigen.nethermind.io/docs/metrics/metrics-prom-spec
package metrics

import (
	"context"
	"errors"
	"net/http"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// EigenMetrics contains instrumented metrics that should be incremented by the avs node using the methods below
type EigenMetrics struct {
	ipPortAddress string
	logger        logging.Logger
	// metrics
	// fees are not yet turned on, so these should just be 0 for the time being
	feeEarnedTotal   *prometheus.CounterVec
	performanceScore prometheus.Gauge
}

var _ Metrics = (*EigenMetrics)(nil)

// Follows the structure from https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#hdr-A_Basic_Example
// TODO(samlaf): I think each avs runs in a separate docker bridge network.
// In order for prometheus to scrape the metrics does the address need to be 0.0.0.0:port to accept connections from other networks?
func NewEigenMetrics(avsName, ipPortAddress string, reg prometheus.Registerer, logger logging.Logger) *EigenMetrics {

	metrics := &EigenMetrics{
		feeEarnedTotal: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace:   types.EigenPromNamespace,
				Name:        "fees_earned_total",
				Help:        "The amount of fees earned in <token>",
				ConstLabels: prometheus.Labels{"avs_name": avsName},
			},
			[]string{"token"},
		),
		performanceScore: promauto.With(reg).NewGauge(
			prometheus.GaugeOpts{
				Namespace:   types.EigenPromNamespace,
				Name:        "performance_score",
				Help:        "The performance metric is a score between 0 and 100 and each developer can define their own way of calculating the score. The score is calculated based on the performance of the Node and the performance of the backing services.",
				ConstLabels: prometheus.Labels{"avs_name": avsName},
			},
		),
		ipPortAddress: ipPortAddress,
		logger:        logger,
	}

	metrics.initMetrics()
	return metrics
}

func (m *EigenMetrics) initMetrics() {
	// Performance score starts as 100, and goes down if node doesn't perform well
	m.performanceScore.Set(100)

	// TODO(samlaf): should we initialize the feeEarnedTotal? This would require the user to pass in a list of tokens for which to initialize the metric
	// same for rpcRequestDurationSeconds and rpcRequestTotal... we could initialize them to be 0 on every json-rpc... but is that really necessary?
}

// AddEigenFeeEarnedTotal adds the fee earned to the total fee earned metric
func (m *EigenMetrics) AddFeeEarnedTotal(amount float64, token string) {
	m.feeEarnedTotal.WithLabelValues(token).Add(amount)
}

// SetPerformanceScore sets the performance score of the node
func (m *EigenMetrics) SetPerformanceScore(score float64) {
	m.performanceScore.Set(score)
}

// Start creates an http handler for reg and starts the prometheus server in a goroutine, listening at m.ipPortAddress.
// reg needs to be the prometheus registry that was passed in the NewEigenMetrics constructor
func (m *EigenMetrics) Start(ctx context.Context, reg prometheus.Gatherer) <-chan error {
	m.logger.Infof("Starting metrics server at port %v", m.ipPortAddress)
	errC := make(chan error, 1)
	go func() {
		http.Handle("/metrics", promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{},
		))
		err := http.ListenAndServe(m.ipPortAddress, nil)
		if err != nil {
			errC <- errors.Join(errors.New("Prometheus server failed"), err)
		} else {
			errC <- nil
		}
	}()
	return errC
}
