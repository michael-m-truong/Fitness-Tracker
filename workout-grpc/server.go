package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	pb "github.com/michael-m-truong/workout-grpc/pb" // Import the generated gRPC code
	repository "github.com/michael-m-truong/workout-grpc/repositories"

	// Import the dotenv package
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type workoutServer struct {
	pb.UnimplementedWorkoutServiceServer
}

func (s *workoutServer) CreateWorkout(ctx context.Context, req *pb.NewWorkout) (*pb.Workout, error) {

	insertedID, err := repository.CreateWorkout(ctx, req)
	if err != nil {
		return nil, err
	}

	err = repository.AddExercisesToWorkout(ctx, *insertedID, req.ExerciseIds)
	if err != nil {
		return nil, err
	}

	workout := &pb.Workout{
		Id:      int32(*insertedID),
		Workout: req,
	}

	return workout, nil
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

	// Register the workoutServer with the gRPC server
	pb.RegisterWorkoutServiceServer(grpcServer, &workoutServer{})

	log.Printf("Starting gRPC server on :%d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
