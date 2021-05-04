package main

import (
	authpb "coolcar/proto/auth/gen/go"
	"coolcar/server/auth/auth"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()

	authpb.RegisterAuthServiceServer(s, &auth.Service{})

	err = s.Serve(listen)

	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
}
