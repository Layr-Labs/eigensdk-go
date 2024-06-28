package txmgr

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namespace = "txmgr"

type Metrics struct {
	broadcastLatencyMs    prometheus.Summary
	confirmationLatencyMs prometheus.Summary
	gasUsedGwei           prometheus.Summary
	speedUps              prometheus.Histogram
	processingTxCount     prometheus.Gauge
	processedTxsTotal     *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer, subsystem string, logger logging.Logger) *Metrics {

	return &Metrics{
		// TODO: we only observe latency of txs that were successfully broadcasted or confirmed.
		// We should also observe latency of txs that failed to broadcast or confirm.
		broadcastLatencyMs: promauto.With(reg).NewSummary(
			prometheus.SummaryOpts{
				Namespace:  namespace,
				Subsystem:  subsystem,
				Name:       "broadcast_latency_ms",
				Help:       "transaction confirmation latency summary in milliseconds",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.01, 0.99: 0.001},
			},
		),
		confirmationLatencyMs: promauto.With(reg).NewSummary(
			prometheus.SummaryOpts{
				Namespace:  namespace,
				Subsystem:  subsystem,
				Name:       "confirmation_latency_ms",
				Help:       "total transaction confirmation latency summary in milliseconds",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.01, 0.99: 0.001},
			},
		),
		gasUsedGwei: promauto.With(reg).NewSummary(
			prometheus.SummaryOpts{
				Namespace:  namespace,
				Subsystem:  subsystem,
				Name:       "gas_used_total",
				Help:       "total gas used to submit each transaction onchain",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.01, 0.99: 0.001},
			},
		),
		speedUps: promauto.With(reg).NewHistogram(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "speedups_total",
				Help:      "number of times a transaction's gas price was increased",
				Buckets:   prometheus.LinearBuckets(0, 1, 10),
			},
		),
		processingTxCount: promauto.With(reg).NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "processing_tx_count",
				Help:      "number of transactions currently being processed",
			},
		),
		processedTxsTotal: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "processed_txs_total",
				Help:      "number of transactions processed by state (success, error)",
			},
			[]string{"state"},
		),
	}
}

func (t *Metrics) ObserveBroadcastLatencyMs(latencyMs int64) {
	t.broadcastLatencyMs.Observe(float64(latencyMs))
}

func (t *Metrics) ObserveConfirmationLatencyMs(latencyMs int64) {
	t.confirmationLatencyMs.Observe(float64(latencyMs))
}

func (t *Metrics) ObserveGasUsedWei(gasUsedWei uint64) {
	gasUsedGwei := gasUsedWei / 1e9
	t.gasUsedGwei.Observe(float64(gasUsedGwei))
}

func (t *Metrics) ObserveSpeedups(speedUps int) {
	t.speedUps.Observe(float64(speedUps))
}

func (t *Metrics) IncrementProcessingTxCount() {
	t.processingTxCount.Inc()
}

func (t *Metrics) DecrementProcessingTxCount() {
	t.processingTxCount.Dec()
}

func (t *Metrics) IncrementProcessedTxsTotal(state string) {
	t.processedTxsTotal.WithLabelValues(state).Inc()
}
