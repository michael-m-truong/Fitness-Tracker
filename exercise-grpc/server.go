package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	pb "github.com/michael-m-truong/exercise-grpc/pb" // Import the generated gRPC code
	repository "github.com/michael-m-truong/exercise-grpc/repositories"

	// Import the dotenv package
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type exerciseServer struct {
	pb.UnimplementedExerciseServiceServer
}

func (s *exerciseServer) CreateExercise(ctx context.Context, req *pb.NewExercise) (*pb.Exercise, error) {

	insertedID, err := repository.CreateExercise(ctx, req)
	if err != nil {
		return nil, err
	}

	exercise := &pb.Exercise{
		Id:       int32(*insertedID),
		Exercise: req,
	}

	return exercise, nil
}

func (s *exerciseServer) CheckExerciseExists(ctx context.Context, req *pb.ExerciseIds) (*pb.ExerciseExistence, error) {

	exist, err := repository.CheckExerciseExists(ctx, req.ExerciseIds)
	if err != nil {
		return nil, err
	}

	exerciseExistence := &pb.ExerciseExistence{
		Exists: exist,
	}

	return exerciseExistence, nil
}

func main() {
	flag.Parse()
	// docker run -d --name postgres-container -e POSTGRES_PASSWORD=mysecretpassword -v <volume_name>:/var/lib/postgresql/data -p 5433:5432 postgres:latest

	// Initialize the database connection
	if _, err := repository.GetDB(); err != nil {
		log.Fatal(err)
	}

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
