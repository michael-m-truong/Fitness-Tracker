package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/michael-m-truong/fitness-tracker/graph/model"
	"github.com/michael-m-truong/fitness-tracker/pb"
	services "github.com/michael-m-truong/fitness-tracker/services"
)

// CreateExercise is the resolver for the createExercise field.
func (r *mutationResolver) CreateExercise(ctx context.Context, input model.NewExercise) (*model.Exercise, error) {
	// Initialize a new Exercise object with the data provided in the input.
	newExercise := model.Exercise{
		Name:        input.Name,
		Description: input.Description,
		MuscleGroup: input.MuscleGroup,
	}

	// Set a fixed value for the User field.
	user := model.User{
		Name: "Mike",
	}
	newExercise.User = &user

	// If the EquipmentInput is provided in the input, create a new Equipment object and assign it to the Exercise.
	if input.Equipment != nil {
		newEquipment := model.Equipment{
			Name: input.Equipment.Name,
			User: &user, // Set the User field for the Equipment to the same fixed user.
		}
		newExercise.Equipment = &newEquipment
	}

	// Here you may want to perform additional logic, such as saving the newExercise and its associated Equipment
	// to a database.
	req := &pb.Exercise{
		Name: input.Name,
		//Description: *input.Description,
		MuscleGroup: input.MuscleGroup,
	}

	if input.Description != nil {
		req.Description = *input.Description
	}

	services.CreateExercise(req)
	// resp, err := services.CreateExercise(req)
	// if err != nil {
	// 		// Handle the error
	// 		return nil, err
	// }

	// Return the newly created exercise along with no error (nil).
	return &newExercise, nil
}

// CreateEquipment is the resolver for the createEquipment field.
func (r *mutationResolver) CreateEquipment(ctx context.Context, input model.NewEquipment) (*model.Equipment, error) {
	panic(fmt.Errorf("not implemented: CreateEquipment - createEquipment"))
}

// CreateWorkout is the resolver for the createWorkout field.
func (r *mutationResolver) CreateWorkout(ctx context.Context, input model.NewWorkout) (*model.Workout, error) {
	panic(fmt.Errorf("not implemented: CreateWorkout - createWorkout"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	// Create a User request from the input
	userReq := &pb.User{
		Username: input.Username,
		Password: input.Password,
		// Add other fields as needed
	}

	// Call the CreateUser service
	newUserResp, err := services.CreateUser(userReq)
	if err != nil {
		return "", err
	}

	// Return a success message or relevant information
	return fmt.Sprintf("User %s created with ID %d", newUserResp.Username, newUserResp.UserId), nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	// Create a User request from the input
	userReq := &pb.User{
		Username: input.Username,
		Password: input.Password,
	}

	// Call the Login service
	accessTokenResp, err := services.Login(userReq)
	if err != nil {
		// Handle authentication errors
		return "", err
	}

	// Return the access token
	return accessTokenResp.Token, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Exercises is the resolver for the exercises field.
func (r *queryResolver) Exercises(ctx context.Context) ([]*model.Exercise, error) {
	var exercises []*model.Exercise
	dummyDescription := "our dummy link"

	dummyExercise := model.Exercise{
		Name:        "our dummy link",
		Description: &dummyDescription,
		Equipment:   &model.Equipment{ID: "equipment-id-1", Name: "Dumbbells"},
		MuscleGroup: "our dummy link",
		User:        &model.User{Name: "admin"},
	}
	exercises = append(exercises, &dummyExercise)
	return exercises, nil
}

// Workouts is the resolver for the workouts field.
func (r *queryResolver) Workouts(ctx context.Context) ([]*model.Workout, error) {
	panic(fmt.Errorf("not implemented: Workouts - workouts"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
