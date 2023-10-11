package internal

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTrace() (*trace.TracerProvider, error) {
	// 1.1. trace를 어떻게 노출시키는 방법을 정의하는 exporter를 선언합니다.
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	// 1.2. 동기적으로 span에 대한 정보를 exporter로 전달하는 processor를 생성합니다.
	spanProcessor := trace.NewSimpleSpanProcessor(exporter)
	// 1.3. tracer provider를 선언합니다.
	provider := trace.NewTracerProvider()
	provider.RegisterSpanProcessor(spanProcessor)
	otel.SetTracerProvider(provider)
	// 5.1. trace 전파를 위해 global propagator를 설정합니다.
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return provider, nil
}
