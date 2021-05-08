package main

import (
	"context"
	authpb "coolcar/proto/auth/gen/go"
	"coolcar/server/auth/auth"
	"coolcar/server/auth/dao"
	"coolcar/server/auth/wechat"
	"coolcar/server/shared/server"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"time"

	"google.golang.org/grpc"
)

const privateKey = "server/shared/private.key"

func main() {
	logger, _ := zap.NewDevelopment()

	ctx := context.Background()

	client := initMongo(ctx, logger)

	pkFile, err := os.Open(privateKey)
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}

	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}
	server.RunGrpcServer(&server.GrpcServerConfig{
		Name:   "auth",
		Addr:   ":8081",
		Logger: logger,
		//AuthPublicKeyFile: privateKey,
		RegisterFunc: func(g *grpc.Server) {
			authpb.RegisterAuthServiceServer(g, &auth.Service{
				Logger:         logger,
				Mongo:          dao.NewMongo(client.Database("coolcar")),
				TokenExpire:    7200,
				TokenGenerator: auth.NewJWTTokenGen("coolcar/auth", privKey),
				OPenIDResolver: &wechat.Service{
					AppID:  "wx61c9acfeed24fcd9",
					Secret: "75875abace4c259b4bdd93a9e506cd7c",
				},
			})
		},
	})
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
