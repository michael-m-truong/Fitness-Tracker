// resolver_test.go
package graph

import (
	"context"
	"testing"

	// Import your models package
	"github.com/michael-m-truong/fitness-tracker/graph/model"
	auth "github.com/michael-m-truong/fitness-tracker/middleware"
	"github.com/michael-m-truong/fitness-tracker/pb"
	stub "github.com/michael-m-truong/fitness-tracker/stubs"
	"github.com/stretchr/testify/assert" // Import the testify assertion library
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ptrString(s string) *string {
	return &s
}

func TestCreateUser(t *testing.T) {
	// Stubbed data and context
	stubContext := context.Background() // Create a stubed context
	stubInput := model.NewUser{         // Stub resolver arguments
		Username: "TestUser",
		Password: "123",
	}
	expectedCharacter := "randomaccesstoken"

	// Create a new instance of your resolver
	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	// Invoke the resolver function
	result, err := resolver.Mutation().CreateUser(stubContext, stubInput)

	// Assertions
	assert.NoError(t, err)                     // No error should be returned
	assert.NotNil(t, result)                   // Result should not be nil
	assert.Equal(t, expectedCharacter, result) // Check if the result matches the expected character
}

func TestAuth(t *testing.T) {
	// Convert the gRPC user response to the appropriate struct
	authUser := &pb.User{
		Username: "TestUser",
	}

	ctx := context.Background()

	// Put user information in the context
	stubContext := context.WithValue(ctx, auth.UserCtxKey, authUser)

	// Retrieve user information from the context
	auth_user := auth.ForContext(stubContext)

	assert.NotNil(t, auth_user)
}

// TODO: do access denied for CreateExercise
func TestCreateExerciseFail(t *testing.T) {

	stubContext := context.Background()

	stubEquipment := model.EquipmentInput{Name: "Barbell"}

	stubInput := model.NewExercise{
		Name:        "Push-Up",
		Description: ptrString("Focus on form"),
		MuscleGroup: "Chest",
		Equipment:   &stubEquipment,
	}
	//expectedID := "123" // Stubbed ID

	resolver := &Resolver{
		ExerciseService: stub.StubExerciseService{},
	}

	result, err := resolver.Mutation().CreateExercise(stubContext, stubInput)

	assert.Equal(t, &model.Exercise{}, result)
	assert.Error(t, err, "Expected an error")
	assert.EqualError(t, err, "access denied", "Expected error message: access denied")
}

// TODO: edge cases for CreateUser
func TestCreateUserFail(t *testing.T) {
	// Stubbed data and context
	stubContext := context.Background() // Create a stubed context
	stubInput := model.NewUser{         // Stub resolver arguments
		Username: "TestUser",
		Password: "123",
	}
	// expectedCharacter := "randomaccesstoken"

	// Create a new instance of your resolver
	resolver := &Resolver{
		AuthService: stub.StubFailAuthService{},
	}

	// Invoke the resolver function
	result, err := resolver.Mutation().CreateUser(stubContext, stubInput)

	// Assertions
	assert.Error(t, err)        // Error should be returned
	assert.NotNil(t, result)    // Result should not be nil
	assert.Equal(t, "", result) // Check if the result matches the expected character
	assert.Equal(t, err, status.Errorf(codes.Internal, "Failed to create user"))
}

func TestCreateExercise(t *testing.T) {
	authUser := &pb.User{
		Username: "TestUser",
	}

	ctx := context.Background()

	// Put user information in the context
	stubContext := context.WithValue(ctx, auth.UserCtxKey, authUser)

	stubEquipment := model.EquipmentInput{Name: "Barbell"}

	stubInput := model.NewExercise{
		Name:        "Push-Up",
		Description: ptrString("Focus on form"),
		MuscleGroup: "Chest",
		Equipment:   &stubEquipment,
	}
	expectedID := "123" // Stubbed ID

	resolver := &Resolver{
		ExerciseService: stub.StubExerciseService{},
	}

	result, err := resolver.Mutation().CreateExercise(stubContext, stubInput)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedID, result.ID)
}

func TestCreateEquipment(t *testing.T) {
	stubContext := context.Background()
	stubInput := model.NewEquipment{
		Name: "Dumbbells",
	}
	//expectedID := "equipment-id-1"

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	assert.Panics(t, func() {
		resolver.Mutation().CreateEquipment(stubContext, stubInput)
	})

	//result, err := resolver.Mutation().CreateEquipment(stubContext, stubInput)

	//assert.Error(t, err)  // Function is not implemented, expecting error
	//assert.Nil(t, result) // Result should be nil
	//assert.Equal(t, expectedID, result.ID)
}

func TestCreateWorkout(t *testing.T) {
	authUser := &pb.User{
		Username: "TestUser",
	}

	ctx := context.Background()

	// Put user information in the context
	stubContext := context.WithValue(ctx, auth.UserCtxKey, authUser)

	stubInput := model.NewWorkout{
		Name: "Upper Body Routine",
		//TODO: Change it to a new type exercise where you just input ID or something
		ExerciseIds: []int{1, 2, 3},
		// Add more exercises as needed
	}

	//expectedID := "workout-id-1"

	resolver := &Resolver{
		AuthService:     stub.StubAuthService{},
		ExerciseService: stub.StubExerciseService{},
	}

	assert.Panics(t, func() {
		resolver.Mutation().CreateWorkout(stubContext, stubInput)
	})

	//assert.Error(t, err)  // Function is not implemented, expecting error
	//assert.Nil(t, result) // Result should be nil
	//assert.Equal(t, expectedID, result.ID)
}

func TestLogin(t *testing.T) {
	stubContext := context.Background()
	stubInput := model.Login{
		Username: "TestUser",
		Password: "123",
	}
	stubToken := "randomaccesstoken"

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	result, err := resolver.Mutation().Login(stubContext, stubInput)

	assert.NoError(t, err)
	assert.Equal(t, stubToken, result)
}

func TestRefreshToken(t *testing.T) {
	stubContext := context.Background()
	stubInput := model.RefreshTokenInput{
		Token: "randomrefreshtoken",
	}
	//expectedToken := "access-token"

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	assert.Panics(t, func() {
		resolver.Mutation().RefreshToken(stubContext, stubInput)
	})
}

func TestExercises(t *testing.T) {
	stubContext := context.Background()

	resolver := &Resolver{}

	assert.Panics(t, func() {
		resolver.Query().Workouts(stubContext)
	})
}

func TestWorkouts(t *testing.T) {
	stubContext := context.Background()

	resolver := &Resolver{}

	assert.Panics(t, func() {
		resolver.Query().Workouts(stubContext)
	})
}
