package metrics

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

type MetricsTestSuite struct {
	suite.Suite
	reg     *prometheus.Registry
	metrics *EigenMetrics
}

// this runs before all tests to initialize reg and metrics
func (suite *MetricsTestSuite) SetupTest() {
	logger := logging.NewNoopLogger()
	// create and start the metrics server
	suite.reg = prometheus.NewRegistry()
	suite.metrics = NewEigenMetrics("testavs", "localhost:9090", suite.reg, logger)
}

// here we just check that the metrics server is running and that the metrics are being emitted
func (suite *MetricsTestSuite) TestEigenMetricsServerIntegration() {
	errC := suite.metrics.Start(context.Background(), suite.reg)

	// we wait a little bit to make sure the server is running without errors
	time.Sleep(1 * time.Second)
	if (len(errC)) != 0 {
		err := <-errC
		if err != nil {
			suite.T().Fatal(err)
		}
	}

	suite.metrics.AddFeeEarnedTotal(1, "testtoken")
	suite.metrics.SetPerformanceScore(1)

	resp, err := http.Get("http://localhost:9090/metrics")
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)

	// We only check for "eigen_performance_score" since it's the only metric that doesn't have a label
	// the other metrics have labels (they are vectors) so they don't appear in the output unless we use them once at least
	assert.Contains(suite.T(), string(body), "eigen_fees_earned_total")
	assert.Contains(suite.T(), string(body), "eigen_performance_score")
}

// The below tests are very pedantic but at least show how avs teams can make use of
// the prometheus testutil package to test their metrics code in more involved scenarios
// See https://pkg.go.dev/github.com/prometheus/client_golang@v1.16.0/prometheus/testutil

func (suite *MetricsTestSuite) TestAddEigenFeeEarnedTotal() {
	suite.metrics.AddFeeEarnedTotal(1, "testtoken")
	assert.Equal(suite.T(), 1.0, testutil.ToFloat64(suite.metrics.feeEarnedTotal.WithLabelValues("testtoken")))
}

func (suite *MetricsTestSuite) TestSetEigenPerformanceScore() {
	suite.metrics.SetPerformanceScore(1)
	assert.Equal(suite.T(), 1.0, testutil.ToFloat64(suite.metrics.performanceScore))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMetricsTestSuite(t *testing.T) {
	suite.Run(t, new(MetricsTestSuite))
}
