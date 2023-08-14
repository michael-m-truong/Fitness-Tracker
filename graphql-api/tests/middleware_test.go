// resolver_test.go
package test

import (
	"context"
	"testing"

	// Import your models package
	"github.com/michael-m-truong/fitness-tracker/graph"
	"github.com/michael-m-truong/fitness-tracker/graph/model"
	stub "github.com/michael-m-truong/fitness-tracker/stubs"
	"github.com/stretchr/testify/assert" // Import the testify assertion library
)

func TestCreateUser(t *testing.T) {
	// Mocked data and context
	mockContext := context.Background() // Create a mocked context
	mockInput := model.NewUser{         // Mock resolver arguments
		Username: "TestUser",
		Password: "123",
	}
	expectedCharacter := "randomaccesstoken"

	// Create a new instance of your resolver
	resolver := &graph.Resolver{
		AuthService: stub.StubAuthService{},
	}

	// Invoke the resolver function
	result, err := resolver.Mutation().CreateUser(mockContext, mockInput)

	// Assertions
	assert.NoError(t, err)                     // No error should be returned
	assert.NotNil(t, result)                   // Result should not be nil
	assert.Equal(t, expectedCharacter, result) // Check if the result matches the expected character
}
