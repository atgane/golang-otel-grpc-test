package main

import (
	"context"
	"fmt"
	"log"
	"main/api"
	"main/internal"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func main() {
	// 2.1. trace provider를 가져옵니다.
	tp, err := internal.InitTrace()
	if err != nil {
		log.Fatal(err)
	}

	// 2.2. tracer를 생성합니다.
	tracer = tp.Tracer("client-go")
	ctx := context.Background()
	someFunc1(ctx)
}

// 2.3. 컨텍스트를 전파합니다.
func someFunc1(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "some func1")
	defer span.End()

	fmt.Println("call func1")
	someFunc2(ctx)
}

func someFunc2(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "some func2")
	defer span.End()

	fmt.Println("call func2")
	callServer(ctx)
}

func callServer(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "call server")
	defer span.End()

	// 3.2. gRPC client를 생성합니다.
	conn, err := internal.CreateClient(ctx, ":7777")
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}
	defer conn.Close()

	// 3.3. 요청을 전달합니다.
	res, err := api.NewDataClient(conn).Get(ctx, &api.GetRequest{Key: "hello"})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}

	fmt.Println(res.Key)
}
