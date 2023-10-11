package internal

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitTrace() (*trace.TracerProvider, error) {
	ctx := context.Background()

	// 6.1. opentelemetry collector로 전달하는 exporter를 선언합니다.
	conn, err := grpc.DialContext(ctx, "localhost:4317", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
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
