package services

import (
	// ...
	"context"
	"log"

	client "github.com/michael-m-truong/fitness-tracker/clients"
	pb "github.com/michael-m-truong/fitness-tracker/pb" // Import the generated gRPC code
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// ...
)

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

func CreateUser(req *pb.User) (*pb.NewUser, error) {
	// Get the gRPC auth client instance
	grpcClient := client.GetAuthClient()

	// Send the RPC and get the response
	resp, err := grpcClient.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call gRPC: %v", err)
		return nil, err
	}

	// Process the response
	log.Printf("Created User: %s with ID: %d", resp.Username, resp.UserId)

	return resp, nil
}

// Login is a function that authenticates a user and generates an access token.
func Login(userReq *pb.User) (*pb.AccessToken, error) {
	// Get the gRPC auth client instance
	grpcClient := client.GetAuthClient()

	// Call the Login service
	resp, err := grpcClient.Login(context.Background(), userReq)
	if err != nil {
		// Handle gRPC errors
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, &WrongUsernameOrPasswordError{}
		}
		return nil, err
	}

	// Prepare the response
	accessToken := &pb.AccessToken{
		Token: resp.Token,
	}

	return accessToken, nil
}

// ParseToken is a function that parses a token and returns the user associated with it.
func ParseToken(token string) (*pb.User, error) {
	// Get the gRPC auth client instance
	grpcClient := client.GetAuthClient()

	// Call the ParseToken service
	resp, err := grpcClient.ParseToken(context.Background(), &pb.AccessToken{Token: token})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/* Token will be generated on login anyway */
// GenerateToken is a function that generates an access token for a user.
// func GenerateToken(user *pb.User) (*pb.AccessToken, error) {
// 	// Get the gRPC auth client instance
// 	grpcClient := client.GetAuthClient()

// 	// Call the GenerateToken service
// 	resp, err := grpcClient.GenerateToken(context.Background(), user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }
