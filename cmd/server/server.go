package main

import (
	"context"
	"log"
	"main/api"
	"main/internal"
	"net"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

var tracer trace.Tracer

func main() {
	// 4.1. trace provider를 가져옵니다.
	tp, err := internal.InitTrace()
	if err != nil {
		log.Fatal(err)
	}

	// 4.2. tracer를 생성합니다.
	tracer = tp.Tracer("server-go")

	l, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}

	// 4.3. data server를 호스팅합니다.
	s := new(DataServer)
	gs := grpc.NewServer()
	api.RegisterDataServer(gs, s)
	if err := gs.Serve(l); err != nil {
		log.Fatal(err)
	}
}

type DataServer struct {
	api.DataServer
}

// 4.4. data server 메서드 get을 구현합니다.
func (d *DataServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	_, span := tracer.Start(ctx, "receive message")
	defer span.End()

	res := &api.GetResponse{Key: "hi"}
	return res, nil
}
