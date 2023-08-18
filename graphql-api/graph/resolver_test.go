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

func TestCreateExercise(t *testing.T) {
	authUser := &pb.User{
		Username: "TestUser",
	}

	ctx := context.Background()

	// Put user information in the context
	stubContext := context.WithValue(ctx, auth.UserCtxKey, authUser)

	stubInput := model.NewExercise{
		Name:        "Push-Up",
		Description: ptrString("Make sure to go to the floor or 90 degrees"),
		MuscleGroup: "Chest",
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

/*func TestCreateEquipment(t *testing.T) {
	stubContext := context.Background()
	stubInput := model.NewEquipment{
		Name: "Dumbbells",
	}
	expectedID := "equipment-id-1"

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	result, err := resolver.Mutation().CreateEquipment(stubContext, stubInput)

	assert.Error(t, err)  // Function is not implemented, expecting error
	assert.Nil(t, result) // Result should be nil
	assert.Equal(t, expectedID, result.ID)
} */

/*func TestCreateWorkout(t *testing.T) {
	stubContext := context.Background()
	stubInput := model.NewWorkout{
		Name: "Upper Body Routine",
	}
	expectedID := "workout-id-1"

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	result, err := resolver.Mutation().CreateWorkout(stubContext, stubInput)

	assert.Error(t, err)  // Function is not implemented, expecting error
	assert.Nil(t, result) // Result should be nil
	assert.Equal(t, expectedID, result.ID)
} */

/*func TestLogin(t *testing.T) {
	stubContext := context.Background()
	stubInput := model.Login{
		Username: "TestUser",
		Password: "123",
	}
	expectedToken := "access-token"

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	result, err := resolver.Mutation().Login(stubContext, stubInput)

	assert.NoError(t, err)
	assert.Equal(t, expectedToken, result)
}

func TestRefreshToken(t *testing.T) {
	stubContext := context.Background()
	stubInput := model.RefreshTokenInput{
		Token: "refresh-token",
	}
	expectedToken := "new-access-token"

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	result, err := resolver.Mutation().RefreshToken(stubContext, stubInput)

	assert.Error(t, err) // Function is not implemented, expecting error
	assert.Equal(t, expectedToken, result)
} */

/*func TestExercises(t *testing.T) {
	stubContext := context.Background()

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	result, err := resolver.Query().Exercises(stubContext)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1) // Assuming the Exercises resolver returns a single dummy exercise
} */

/*func TestWorkouts(t *testing.T) {
	stubContext := context.Background()

	resolver := &Resolver{
		AuthService: stub.StubAuthService{},
	}

	result, err := resolver.Query().Workouts(stubContext)

	assert.Error(t, err)  // Function is not implemented, expecting error
	assert.Nil(t, result) // Result should be nil
} */
