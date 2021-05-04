package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	tripClient := trippb.NewTripServiceClient(client)

	resp, err := tripClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip456",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
