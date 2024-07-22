package nodeapi

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

var logger = logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})
var testNodeApi = NewNodeApi("testAvs", "v0.0.1", "localhost:8080", logger)

// just making sure that the nodeapi starts without any errors
func TestStart(t *testing.T) {
	errC := testNodeApi.Start()
	select {
	case <-time.After(3 * time.Second):
		// consider it a pass if no errors received after 3 seconds
	case err := <-errC:
		assert.NoError(t, err)
	}
}

func TestNodeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/eigen/node", nil)
	w := httptest.NewRecorder()

	testNodeApi.nodeHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "{\"node_name\":\"testAvs\",\"node_version\":\"v0.0.1\",\"spec_version\":\"v0.0.1\"}\n", string(data))
}

func TestHealthHandler(t *testing.T) {

	tests := map[string]struct {
		nodeApi        *NodeApi
		wantStatusCode int
		wantBody       string
	}{
		"healthy": {
			nodeApi:        testNodeApi,
			wantStatusCode: http.StatusOK,
			wantBody:       "",
		},
		"partially healthy": {
			nodeApi: func() *NodeApi {
				nodeApi := NewNodeApi("testAvs", "v0.0.1", "localhost:8080", logger)
				nodeApi.UpdateHealth(PartiallyHealthy)
				return nodeApi
			}(),
			wantStatusCode: http.StatusPartialContent,
			wantBody:       "",
		},
		"unhealthy": {
			nodeApi: func() *NodeApi {
				nodeApi := NewNodeApi("testAvs", "v0.0.1", "localhost:8080", logger)
				nodeApi.UpdateHealth(Unhealthy)
				return nodeApi
			}(),
			wantStatusCode: http.StatusServiceUnavailable,
			wantBody:       "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/eigen/node/health", nil)
			w := httptest.NewRecorder()
			test.nodeApi.healthHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			data, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			assert.Equal(t, test.wantStatusCode, res.StatusCode)
			assert.Equal(t, test.wantBody, string(data))
		})
	}
}

func TestServicesHandler(t *testing.T) {

	tests := map[string]struct {
		nodeApi        *NodeApi
		wantStatusCode int
		wantBody       string
	}{
		"no services": {
			nodeApi:        testNodeApi,
			wantStatusCode: http.StatusOK,
			wantBody:       "{\"services\":[]}\n",
		},
		"one service": {
			nodeApi: func() *NodeApi {
				nodeApi := NewNodeApi("testAvs", "v0.0.1", "localhost:8080", logger)
				nodeApi.RegisterNewService("testServiceId", "testServiceName", "testServiceDescription", ServiceStatusUp)
				return nodeApi
			}(),
			wantStatusCode: http.StatusOK,
			wantBody:       "{\"services\":[{\"id\":\"testServiceId\",\"name\":\"testServiceName\",\"description\":\"testServiceDescription\",\"status\":\"Up\"}]}\n",
		},
		"two services": {
			nodeApi: func() *NodeApi {
				nodeApi := NewNodeApi("testAvs", "v0.0.1", "localhost:8080", logger)
				nodeApi.RegisterNewService("testServiceId", "testServiceName", "testServiceDescription", ServiceStatusUp)
				nodeApi.RegisterNewService("testServiceId2", "testServiceName2", "testServiceDescription2", ServiceStatusDown)
				return nodeApi
			}(),
			wantStatusCode: http.StatusOK,
			wantBody:       "{\"services\":[{\"id\":\"testServiceId\",\"name\":\"testServiceName\",\"description\":\"testServiceDescription\",\"status\":\"Up\"},{\"id\":\"testServiceId2\",\"name\":\"testServiceName2\",\"description\":\"testServiceDescription2\",\"status\":\"Down\"}]}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/eigen/node/services", nil)
			w := httptest.NewRecorder()
			test.nodeApi.servicesHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			data, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			assert.Equal(t, test.wantStatusCode, res.StatusCode)
			assert.Equal(t, test.wantBody, string(data))
		})
	}
}

func TestServiceHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/eigen/node/services/123/health", nil)
	w := httptest.NewRecorder()

	testNodeApi.serviceHealthHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, "", string(data))
}
