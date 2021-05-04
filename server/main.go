package main

import (
	trippb "coolcar/proto/gen/go"
	trip "coolcar/server/tripservice"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)
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
