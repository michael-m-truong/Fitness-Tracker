package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/michael-m-truong/exercise-grpc/pb" // Import the generated gRPC code
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type exerciseServer struct {
	pb.UnimplementedExerciseServiceServer
}

func (s *exerciseServer) AddExercise(ctx context.Context, req *pb.Exercise) (*pb.Exercise, error) {
	// Implement the logic to handle the incoming exercise request and save it to the database
	fmt.Printf("Received Exercise: %+v\n", req)
	return req, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register the exerciseServer with the gRPC server
	pb.RegisterExerciseServiceServer(grpcServer, &exerciseServer{})

	log.Printf("Starting gRPC server on :%d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
