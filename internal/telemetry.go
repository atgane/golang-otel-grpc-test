package internal

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTrace() (*trace.TracerProvider, error) {
	// 1. trace를 어떻게 노출시키는 방법을 정의하는 exporter를 선언합니다.
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	// 2. 동기적으로 span에 대한 정보를 exporter로 전달하는 processor를 생성합니다.
	spanProcessor := trace.NewSimpleSpanProcessor(exporter)
	// 3. tracer provider를 선언합니다.
	provider := trace.NewTracerProvider()
	provider.RegisterSpanProcessor(spanProcessor)
	return provider, nil
}
