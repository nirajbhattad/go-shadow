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

	// Create a listener on TCP port 5001
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoreRouterServer(grpcServer, shadowServer)

	log.Println("Starting server on port 5001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// ShadowServer implements the ShadowServer server
type ShadowServer struct {
	pb.UnimplementedCoreRouterServer
}

func (s ShadowServer) Send(ctx context.Context, req *pb.Transaction) (*emptypb.Empty, error) {
	message := req.GetMessage()
	ipAddress := "localhost"
	port := int(req.GetPortNumber())

	log.Printf("Received message to forward to tcp-proxy from core-router: %s, port %d", string(message), port)

	// Establish a socket connection
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipAddress, port))
	if err != nil {
		log.Printf("Failed to connect to tcp-proxy on %s:%d: %v", ipAddress, port, err)
		return nil, status.Errorf(codes.Internal, "Failed to establish socket connection with tcp-proxy")
	}
	defer conn.Close()

	// Forward the raw message over the socket connection
	_, err = conn.Write(message)
	if err != nil {
		log.Printf("Failed to send raw message to tcp-proxy: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to send raw message")
	}

	// Receive and print the response
	buffer := make([]byte, 1024)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response from tcp-proxy:", err.Error())
	}
	response := string(buffer[:bytesRead])
	log.Println("Response from tcp-proxy:", response)

	return &emptypb.Empty{}, nil
}
