package txmgr

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namespace = "txmgr"

type Metrics struct {
	Latency  *prometheus.SummaryVec
	GasUsed  prometheus.Gauge
	SpeedUps prometheus.Gauge
	TxQueue  prometheus.Gauge
	NumTx    *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer, subsystem string, logger logging.Logger) *Metrics {

	return &Metrics{
		Latency: promauto.With(reg).NewSummaryVec(
			prometheus.SummaryOpts{
				Namespace:  namespace,
				Subsystem:  subsystem,
				Name:       "latency_ms",
				Help:       "transaction confirmation latency summary in milliseconds",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.01, 0.99: 0.001},
			},
			[]string{"stage"},
		),
		GasUsed: promauto.With(reg).NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "gas_used",
				Help:      "gas used for onchain batch confirmation",
			},
		),
		SpeedUps: promauto.With(reg).NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "speedups_count",
				Help:      "number of times the gas price was increased",
			},
		),
		TxQueue: promauto.With(reg).NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "tx_queue_len",
				Help:      "number of transactions in transaction queue",
			},
		),
		NumTx: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "txs_total",
				Help:      "number of transactions processed",
			},
			[]string{"state"},
		),
	}
}

func (t *Metrics) ObserveLatency(stage string, latencyMs float64) {
	t.Latency.WithLabelValues(stage).Observe(latencyMs)
}

func (t *Metrics) UpdateGasUsed(gasUsed uint64) {
	t.GasUsed.Set(float64(gasUsed))
}

func (t *Metrics) UpdateSpeedUps(speedUps int) {
	t.SpeedUps.Set(float64(speedUps))
}

func (t *Metrics) UpdateTxQueue(txQueue int) {
	t.TxQueue.Set(float64(txQueue))
}

func (t *Metrics) IncrementTxnCount(state string) {
	t.NumTx.WithLabelValues(state).Inc()
}
