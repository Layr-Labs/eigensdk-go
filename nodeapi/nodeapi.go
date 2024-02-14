package nodeapi

// Implementation of https://eigen.nethermind.io/docs/spec/api/#api-versioning

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

const (
	baseUrl = "/eigen"
	// Spec version is the version of the avs node spec that this node is implementing
	// see https://eigen.nethermind.io/docs/spec/api/#api-versioning
	specSemVer = "v0.0.1"
)

type NodeHealth int

const (
	Healthy          NodeHealth = iota
	PartiallyHealthy            // either initializing or some backing services are not healthy
	Unhealthy
)

type ServiceStatus string

const (
	ServiceStatusUp           ServiceStatus = "Up"
	ServiceStatusDown         ServiceStatus = "Down"
	ServiceStatusInitializing ServiceStatus = "Initializing"
)

type nodeService struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Status      ServiceStatus `json:"status"`
}

type NodeApi struct {
	avsNodeName   string
	avsNodeSemVer string
	health        NodeHealth
	nodeServices  []nodeService
	ipPortAddr    string
	logger        logging.Logger
}

func NewNodeApi(avsNodeName, avsNodeSemVer, IpPortAddr string, logger logging.Logger) *NodeApi {
	nodeApi := &NodeApi{
		avsNodeName:   avsNodeName,
		avsNodeSemVer: avsNodeSemVer,
		health:        Healthy,
		nodeServices:  []nodeService{},
		ipPortAddr:    IpPortAddr,
		logger:        logger,
	}
	return nodeApi
}

func (api *NodeApi) UpdateHealth(health NodeHealth) {
	api.health = health
}

func (api *NodeApi) RegisterNewService(serviceId, serviceName, serviceDescription string, serviceStatus ServiceStatus) {
	api.nodeServices = append(api.nodeServices, nodeService{
		Id:          serviceId,
		Name:        serviceName,
		Description: serviceDescription,
		Status:      serviceStatus,
	})
}

func (api *NodeApi) UpdateServiceStatus(serviceId string, serviceStatus ServiceStatus) error {
	for i, service := range api.nodeServices {
		if service.Id == serviceId {
			api.nodeServices[i].Status = serviceStatus
			return nil
		}
	}
	return fmt.Errorf("Service with serviceId %v not found", serviceId)
}

func (api *NodeApi) DeregisterService(serviceId string) error {
	for i, service := range api.nodeServices {
		if service.Id == serviceId {
			api.nodeServices = append(api.nodeServices[:i], api.nodeServices[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Service with serviceId %v not found", serviceId)
}

// Start starts the node api server in a goroutine
func (api *NodeApi) Start() <-chan error {
	api.logger.Infof("Starting node api server at address %v", api.ipPortAddr)

	mux := http.NewServeMux()
	httpServer := http.Server{
		Addr:    api.ipPortAddr,
		Handler: mux,
	}

	mux.HandleFunc(baseUrl+"/node", api.nodeHandler)
	mux.HandleFunc(baseUrl+"/node/health", api.healthHandler)
	mux.HandleFunc(baseUrl+"/node/services", api.servicesHandler)
	// Note: You'll need to extract the service_ID from the URL
	// /node/services/{service_ID}/health
	mux.HandleFunc(baseUrl+"/node/services/", api.serviceHealthHandler)

	errChan := run(api.logger, &httpServer)
	return errChan
}

// https://eigen.nethermind.io/docs/metrics/metrics-api#get-eigennodespec-version
func (api *NodeApi) specVersionHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"spec_version": specSemVer,
	}
	err := jsonResponse(w, response)
	if err != nil {
		api.logger.Error("Error in specVersionHandler", "err", err)
	}
}

// https://eigen.nethermind.io/docs/metrics/metrics-api#get-eigennodeversion
func (api *NodeApi) nodeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"node_name":    api.avsNodeName,
		"spec_version": specSemVer,
		"node_version": api.avsNodeSemVer,
	}
	err := jsonResponse(w, response)
	if err != nil {
		api.logger.Error("Error in nodeVersionHandler", "err", err)
	}
}

// https://eigen.nethermind.io/docs/metrics/metrics-api#get-eigennodehealth
func (api *NodeApi) healthHandler(w http.ResponseWriter, r *http.Request) {
	switch api.health {
	case Healthy:
		// 200 - Node is healthy
		w.WriteHeader(http.StatusOK)
	case PartiallyHealthy:
		// 206 - Node is partially healthy. It is either initializing or some backing services are not healthy.
		w.WriteHeader(http.StatusPartialContent)
	case Unhealthy:
		// 503 - Node is unhealthy or having issues.
		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		// we still return unhealthy if we don't know the health status, to prevent caller from hanging
		w.WriteHeader(http.StatusServiceUnavailable)
		api.logger.Error("Unknown health status", "health", api.health)
	}
}

// https://eigen.nethermind.io/docs/metrics/metrics-api#get-eigennodeservices
func (api *NodeApi) servicesHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string][]nodeService{
		"services": api.nodeServices,
	}
	err := jsonResponse(w, response)
	if err != nil {
		api.logger.Error("Error in servicesHandler", "err", err)
	}
}

// https://eigen.nethermind.io/docs/metrics/metrics-api#get-eigennodeservicesservice_idhealth
func (api *NodeApi) serviceHealthHandler(w http.ResponseWriter, r *http.Request) {
	suffix, found := strings.CutPrefix(r.URL.Path, "/eigen/node/services/")
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	serviceId, found := strings.CutSuffix(suffix, "/health")
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// find the service
	for _, service := range api.nodeServices {
		if service.Id == serviceId {
			switch service.Status {
			case ServiceStatusUp:
				w.WriteHeader(http.StatusOK)
			case ServiceStatusDown:
				w.WriteHeader(http.StatusServiceUnavailable)
			case ServiceStatusInitializing:
				w.WriteHeader(http.StatusPartialContent)
			default:
				w.WriteHeader(http.StatusServiceUnavailable)
				api.logger.Error("Unknown service status", "serviceStatus", service.Status)
			}
			return
		}
	}

	// if the service has not been registered, return 404
	w.WriteHeader(http.StatusNotFound)
}

func run(logger logging.Logger, httpServer *http.Server) <-chan error {
	errChan := make(chan error, 1)
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()

		logger.Info("shutdown signal received")

		defer func() {
			stop()
			close(errChan)
		}()

		if err := httpServer.Shutdown(context.Background()); err != nil {
			errChan <- err
		}
		logger.Info("shutdown completed")
	}()

	go func() {
		logger.Info("node api server running", "addr", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func jsonResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	return err
}
