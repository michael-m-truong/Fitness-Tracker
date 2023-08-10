package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	pb "github.com/michael-m-truong/exercise-grpc/pb" // Import the generated gRPC code

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
	fmt.Printf("Received Exercise: %+v\n", req)

	db := getDB() // Access the singleton database instance
	// Prepare the INSERT query with placeholders
	insertQuery := `
		INSERT INTO exercise (name, description, muscle_group, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	// Execute the INSERT query with parameters and retrieve the inserted exercise's ID
	var insertedID int
	err := db.QueryRowContext(ctx, insertQuery, req.Name, req.Description, req.MuscleGroup, req.UserId).Scan(&insertedID)
	if err != nil {
		return nil, err
	}
	exercise := &pb.Exercise{
		Id:       int32(insertedID),
		Exercise: req,
	}

	// Update the request with the inserted ID

	return exercise, nil
}

func main() {
	flag.Parse()
	// docker run -d --name postgres-container -e POSTGRES_PASSWORD=mysecretpassword -v <volume_name>:/var/lib/postgresql/data -p 5433:5432 postgres:latest

	// Initialize the database connection
	if err := initDB(); err != nil {
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
