package main

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	"coolcar/server/rental/trip"
	"coolcar/server/shared/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

const endpoints = "localhost:8085"

func main() {
	//首先初始化 log
	logger, _ := zap.NewDevelopment()
	//grpc 端口
	listen, err := net.Listen("tcp", ":8085")

	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	interceptor, err := auth.Interceptor("server/shared/public.key")
	if err != nil {
		log.Fatalf("make Interceptor error :%v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor))

	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})

	go grpcServer()

	err = s.Serve(listen)

	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	go grpcServer()

}

func grpcServer() {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true,
		},
	}))

	err := rentalpb.RegisterTripServiceHandlerFromEndpoint(context.Background(), mux, endpoints, []grpc.DialOption{
		grpc.WithInsecure(),
	})
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
}

//type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
