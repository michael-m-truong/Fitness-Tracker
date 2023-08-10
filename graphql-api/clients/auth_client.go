package client

import (
	pb "github.com/michael-m-truong/fitness-tracker/pb" // Import the generated gRPC code
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAuthClient pb.AuthServiceClient
)

func init() {
	// Create a gRPC connection to the server
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to connect to gRPC server")
	}

	// Create the gRPC client
	grpcAuthClient = pb.NewAuthServiceClient(conn)
}

// GetExerciseClient returns the gRPC client instance
func GetAuthClient() pb.AuthServiceClient {
	return grpcAuthClient
}
