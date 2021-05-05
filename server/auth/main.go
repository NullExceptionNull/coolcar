package main

import (
	"context"
	authpb "coolcar/proto/auth/gen/go"
	"coolcar/server/auth/auth"
	"coolcar/server/auth/dao"
	"coolcar/server/auth/wechat"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

const (
	endpoints  = "localhost:8083"
	issuer     = "coolcar/auth"
	privateKey = ``
)

func main() {

	logger, _ := zap.NewDevelopment()

	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	go startGRPCGateway()

	ctx := context.Background()

	client := initMongo(ctx, err, logger)

	s := grpc.NewServer()

	bytes, err := ioutil.ReadFile("server/auth/private.key")

	if err != nil {
		log.Fatalf("open file error : %v", err)
	}
	pem, _ := jwt.ParseRSAPrivateKeyFromPEM(bytes)

	authpb.RegisterAuthServiceServer(s, &auth.Service{
		Logger:         logger,
		Mongo:          dao.NewMongo(client.Database("coolcar")),
		TokenGenerator: auth.NewJWTTokenGen(issuer, pem),
		TokenExpire:    7200,
		OPenIDResolver: &wechat.Service{
			AppID:  "wx61c9acfeed24fcd9",
			Secret: "75875abace4c259b4bdd93a9e506cd7c",
		},
	})

	err = s.Serve(listen)

	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
}

func initMongo(ctx context.Context, err error, logger *zap.Logger) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("Cannot connect to MongoDB")
	}
	return client
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	marshalOptions := protojson.MarshalOptions{}

	marshalOptions.UseEnumNumbers = true

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseEnumNumbers: true,
					UseProtoNames:  true,
				},
			}),
	)

	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		c, mux, endpoints, []grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", mux)
}
