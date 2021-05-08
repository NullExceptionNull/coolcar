package main

import (
	rentalpb "coolcar/proto/rental/gen/go"
	"coolcar/server/rental/trip"
	"coolcar/server/shared/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//首先初始化 log
	logger, _ := zap.NewDevelopment()

	logger.Sugar().Fatal(
		server.RunGrpcServer(&server.GrpcServerConfig{
			Name:              "Rental",
			Addr:              ":8082",
			AuthPublicKeyFile: "shared/auth/public.key",
			Logger:            logger,
			RegisterFunc: func(g *grpc.Server) {
				rentalpb.RegisterTripServiceServer(g, &trip.Service{
					Logger: logger,
				})
			},
		}))
}
