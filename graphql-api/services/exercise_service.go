package services

import (
	// ...
	"context"
	"log"

	client "github.com/michael-m-truong/fitness-tracker/clients"
	pb "github.com/michael-m-truong/fitness-tracker/pb" // Import the generated gRPC code
	// ...
)

type IExerciseService interface {
	CreateExercise(req *pb.NewExercise) (*pb.Exercise, error)
}

type ExerciseResolverService struct{}

func (service ExerciseResolverService) CreateExercise(req *pb.NewExercise) (*pb.Exercise, error) {
	// Get the gRPC client instance
	grpcClient := client.GetExerciseClient()

	// Send the RPC and get the response
	req.UserId = 21
	resp, err := grpcClient.CreateExercise(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call gRPC: %v", err)
		return nil, err
	}

	// Process the response
	log.Printf("Received Exercise Name: %s", resp.Exercise.Name)
	log.Printf("Received Muscle Group: %s", resp.Exercise.MuscleGroup)
	log.Printf("Received Description: %s", resp.Exercise.Description)
	log.Printf("Received UserID: %d", resp.Id)

	return resp, nil
}
