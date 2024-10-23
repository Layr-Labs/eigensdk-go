package telemetry

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

func InitTelemetry(key string) {
	sync.OnceFunc(func() {
		telemetry := &Telemetry{} // TODO: args
		atomic.SwapPointer(&telemetrySingleton, unsafe.Pointer(telemetry))
	})()
}

func GetTelemetry() (*Telemetry, error) {
	telemetry := (*Telemetry)(atomic.LoadPointer(&telemetrySingleton))
	if telemetry == nil {
		return nil, fmt.Error()
	}
	return telemetry, nil
}
