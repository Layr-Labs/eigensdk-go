package rpccalls

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestRpcCallsCollector(t *testing.T) {
	reg := prometheus.NewRegistry()
	rpcCallsCollector := NewCollector("testavs", "testname", reg)

	rpcCallsCollector.ObserveRPCRequestDurationSeconds(1, "testmethod", "testclient/testversion")
	// TODO(samlaf): not sure how to test histogram values.. there's no mention of histograms in
	//               https://pkg.go.dev/github.com/prometheus/client_golang/prometheus/testutil
	// assert.Equal(t, 1.0, testutil.ToFloat64(suite.metrics.rpcRequestDurationSeconds))

	rpcCallsCollector.AddRPCRequestTotal("testmethod", "testclient/testversion")
	assert.Equal(t, 1.0, testutil.ToFloat64(rpcCallsCollector.rpcRequestTotal.WithLabelValues("testmethod", "testclient/testversion")))
}
