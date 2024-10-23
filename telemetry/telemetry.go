package telemetry

import (
	"os"

	"github.com/posthog/posthog-go"
)

func config() {
	client, _ := posthog.NewWithConfig(
		os.Getenv("POSTHOG_API_KEY"),
		posthog.Config{
			PersonalApiKey: "your personal API key", // Optional, but much more performant.  If this token is not supplied, then fetching feature flag values will be slower.
			Endpoint:       "https://us.i.posthog.com",
		},
	)
	defer client.Close()
	// run commands

}

func captureEvent(client *posthog.Client) {
	(*client).Enqueue(posthog.Capture{
		DistinctId: "distinct_id_of_the_user",
		Event:      "user_signed_up",
	})
}

/*
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
*/
