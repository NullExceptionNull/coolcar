package main

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	"coolcar/server/rental/trip"
	"coolcar/server/rental/trip/client/car"
	"coolcar/server/rental/trip/client/poi"
	"coolcar/server/rental/trip/client/profile"
	"coolcar/server/rental/trip/dao"
	"coolcar/server/shared/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func main() {
	//首先初始化 log
	logger, _ := zap.NewDevelopment()

	client := initMongo(context.Background(), logger)

	logger.Sugar().Fatal(
		server.RunGrpcServer(&server.GrpcServerConfig{
			Name: "Rental",
			Addr: ":8082",
			//AuthPublicKeyFile: "shared/auth/public.key",
			Logger: logger,
			RegisterFunc: func(g *grpc.Server) {
				rentalpb.RegisterTripServiceServer(g, &trip.Service{
					Logger:         logger,
					CarManager:     &car.Manager{},
					ProfileManager: &profile.Manager{},
					POIManager:     &poi.Manager{},
					Mongo:          dao.NewMongo(client.Database("coolcar")),
				})
			},
		}))
}

func initMongo(ctx context.Context, logger *zap.Logger) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("Cannot connect to MongoDB")
	}
	return client

}
