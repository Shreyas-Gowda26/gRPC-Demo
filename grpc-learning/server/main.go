package main

import (
	"context"
	"fmt"
	"net"

	pb "grpc-learning/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(
	ctx context.Context,
	req *pb.AddRequest,
) (*pb.AddResponse, error) {

	result := req.A + req.B

	return &pb.AddResponse{
		Result: result,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculatorServer(
		grpcServer,
		&server{},
	)

	fmt.Println("gRPC Server running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
