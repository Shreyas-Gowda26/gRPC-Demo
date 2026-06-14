package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-learning/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	resp, err := client.Add(
		context.Background(),
		&pb.AddRequest{
			A: 10,
			B: 20,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", resp.Result)
}
