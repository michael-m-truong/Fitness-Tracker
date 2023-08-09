package services

import (
	// ...
	"context"
	"log"

	client "github.com/michael-m-truong/fitness-tracker/clients"
	pb "github.com/michael-m-truong/fitness-tracker/pb" // Import the generated gRPC code
	// ...
)

func CreateExercise(req *pb.Exercise) (*pb.Exercise, error) {
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
	log.Printf("Received Exercise Name: %s", resp.Name)
	log.Printf("Received Muscle Group: %s", resp.MuscleGroup)
	log.Printf("Received Description: %s", resp.Description)
	log.Printf("Received UserID: %d", resp.UserId)

	return resp, nil
}
