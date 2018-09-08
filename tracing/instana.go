// +build opentracing,instana

package tracing

import (
	"strings"

	instana "github.com/instana/golang-sensor"
	opentracing "github.com/opentracing/opentracing-go"
)

func init() {
	tracers["instana"] = initInstanaTracer
}

func initInstanaTracer(opts []string) (opentracing.Tracer, error) {
	serviceName := defServiceName

	for _, o := range opts {
		parts := strings.SplitN(o, "=", 2)
		switch parts[0] {
		case "service-name":
			if len(parts) > 1 {
				serviceName = parts[1]
			}
		}
	}

	return instana.NewTracerWithOptions(&instana.Options{
		Service:  serviceName,
		LogLevel: instana.Error,
	}), nil
}
