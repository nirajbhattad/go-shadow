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
	message := req.GetMessage()
	ipAddress := "localhost" // Replace with the specific IP address from the config
	port := int(req.GetPortNumber())

	// establish socket connection
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipAddress, port))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to establish socket connection")
	}
	defer conn.Close()

	// send message over socket connection
	_, err = conn.Write(message)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to send raw message")
	}

	buffer := make([]byte, 1024)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
	}
	response := string(buffer[:bytesRead])
	fmt.Println("Response from server:", response)

	return &emptypb.Empty{}, nil
}


