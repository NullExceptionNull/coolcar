package main

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	"coolcar/server/rental/trip"
	"coolcar/server/shared/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
)

const endpoints = "localhost:8085"

func main() {
	//首先初始化 log
	logger, _ := zap.NewDevelopment()

	logger.Sugar().Fatal(
		server.RunGrpcServer(&server.GrpcServerConfig{
			Name:              "Rental",
			Addr:              ":8088",
			AuthPublicKeyFile: "shared/auth/public.key",
			Logger:            logger,
			RegisterFunc: func(g *grpc.Server) {
				rentalpb.RegisterTripServiceServer(g, &trip.Service{
					Logger: logger,
				})
			},
		}))
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
