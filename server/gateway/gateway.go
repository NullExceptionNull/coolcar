package main

import (
	"context"
	authpb "coolcar/proto/auth/gen/go"
	rentalpb "coolcar/proto/rental/gen/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

func main() {
	logger, _ := zap.NewDevelopment()

	c := context.Background()

	ctx, cancelFunc := context.WithCancel(c)

	defer cancelFunc()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true,
		},
	}))

	serverConfig := []*struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         "localhost:8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		}, {
			name:         "rental",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}

	for _, config := range serverConfig {
		_ = config.registerFunc(ctx, mux, config.addr, []grpc.DialOption{
			grpc.WithInsecure(),
		})
	}

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		logger.Fatal("gateway run error", zap.Error(err))
	}
}
