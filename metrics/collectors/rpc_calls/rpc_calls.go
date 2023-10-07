package rpccalls

import (
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Collector contains instrumented metrics that should be incremented by the avs node using the methods below
// it is used by the eigensdk's instrumented_client, but can also be used by avs teams to instrument their own client
// if it differs from ours.
type Collector struct {
	rpcRequestDurationSeconds *prometheus.HistogramVec
	rpcRequestTotal           *prometheus.CounterVec
}

// NewCollector returns an rpccalls Collector that collects metrics for json-rpc calls
func NewCollector(avsName string, reg prometheus.Registerer) *Collector {
	return &Collector{
		rpcRequestDurationSeconds: promauto.With(reg).NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace:   types.EigenPromNamespace,
				Name:        "rpc_request_duration_seconds",
				Help:        "Duration of json-rpc <method> in seconds",
				ConstLabels: prometheus.Labels{"avs_name": avsName},
			},
			// client_version is the client name and its current version. We don't separate them because
			// this string is returned as a single string by the web3_clientVersion jsonrpc call
			// which doesn't follow any standardized format so not possible to parse it...
			// https://ethereum.org/en/developers/docs/apis/json-rpc/#web3_clientversion
			[]string{"method", "client_version"},
		),
		rpcRequestTotal: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace:   types.EigenPromNamespace,
				Name:        "rpc_request_total",
				Help:        "Total number of json-rpc <method> requests",
				ConstLabels: prometheus.Labels{"avs_name": avsName},
			},
			[]string{"method", "client_version"},
		),
	}
}

// ObserveRPCRequestDurationSeconds observes the duration of a json-rpc request
func (c *Collector) ObserveRPCRequestDurationSeconds(duration float64, method, clientVersion string) {
	c.rpcRequestDurationSeconds.With(prometheus.Labels{
		"method":         method,
		"client_version": clientVersion,
	}).Observe(duration)
}

// AddRPCRequestTotal adds a json-rpc request to the total number of requests
func (c *Collector) AddRPCRequestTotal(method, clientVersion string) {
	c.rpcRequestTotal.With(prometheus.Labels{
		"method":         method,
		"client_version": clientVersion,
	}).Inc()
}
