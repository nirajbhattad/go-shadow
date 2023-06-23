package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "go-shadow/proto"
)

func main() {
	var shadowServer ShadowServer

	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoreRouterServer(grpcServer, shadowServer)

	fmt.Println("Server started on port 5001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

// CalculatorServer implements the CalculatorService server
type ShadowServer struct {
	pb.UnimplementedCoreRouterServer
}

func (s ShadowServer) Send(ctx context.Context, req *pb.Transaction) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

// CalculatorServer implements the CalculatorService server
type CalculatorServer struct {
	pb.UnimplementedCalculatorServiceServer
}

// Add implements the Add RPC method
func (s CalculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.Num1 + req.Num2
	return &pb.AddResponse{Result: result}, nil
}
