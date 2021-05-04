package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	trip "coolcar/server/tripservice"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const endpoints = "localhost:8081"

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})

	err = s.Serve(listen)

	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux()

	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c, mux, endpoints, []grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", mux)
}
