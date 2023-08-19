package services

import (
	// ...
	"context"
	"log"

	client "github.com/michael-m-truong/fitness-tracker/clients"
	pb "github.com/michael-m-truong/fitness-tracker/pb" // Import the generated gRPC code
	// ...
)

type IWorkoutService interface {
	CreateWorkout(req *pb.NewWorkout) (*pb.Workout, error)
}

type WorkoutResolverService struct{}

func (service WorkoutResolverService) CreateWorkout(req *pb.NewWorkout) (*pb.Workout, error) {
	// Get the gRPC client instance
	grpcClient := client.GetWorkoutClient()

	// Send the RPC and get the response
	//req.UserId = 21
	resp, err := grpcClient.CreateWorkout(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call gRPC: %v", err)
		return nil, err
	}

	// Process the response
	log.Printf("Received Workout Name: %s", resp.Workout.Name)
	log.Printf("Received Description: %s", resp.Workout.Description)
	log.Printf("Received UserID: %d", resp.Id)

	return resp, nil
}
