package trace

import "github.com/kitex-contrib/obs-opentelemetry/provider"

func TraceInit(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	return p
}
