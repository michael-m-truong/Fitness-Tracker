package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/michael-m-truong/auth-grpc/jwt"
	pb "github.com/michael-m-truong/auth-grpc/pb" // Import the generated gRPC code
	repository "github.com/michael-m-truong/auth-grpc/repositories"
	service "github.com/michael-m-truong/auth-grpc/services"

	// Import the dotenv package
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *authServer) Login(ctx context.Context, req *pb.User) (*pb.AccessToken, error) {
	// Call the service function to perform business logic
	passwordHash, err := service.HashPassword(req.Password)
	if err != nil {
		// Handle the error, perhaps logging or performing additional actions
		return nil, err
	}

	// Call the repository to save or fetch data (if needed)
	dbPasswordHash, err := repository.CheckLoginInfo(ctx, req.Username, passwordHash)
	if err != nil {
		// Handle the error, perhaps logging or performing additional actions
		return nil, err
	}

	if service.CheckPasswordHash(*dbPasswordHash, passwordHash) {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
	}
	// Generate and return an access token
	tokenString, err := jwt.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token")
	}
	return &pb.AccessToken{Token: tokenString}, nil
}

func (s *authServer) CreateUser(ctx context.Context, req *pb.User) (*pb.NewUser, error) {
	// Hash the password
	hashedPassword, err := service.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password")
	}

	userID, err := repository.CreateUser(ctx, req.Username, hashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user")
	}

	tokenString, err := jwt.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token")
	}

	// Create a NewUser response
	// TODO: figure out how to maybe pass nil values in proto
	newUser := &pb.NewUser{
		UserId:      *userID,
		Username:    req.Username,
		AccessToken: tokenString,
	}

	return newUser, nil
}

func (s *authServer) ParseToken(ctx context.Context, req *pb.AccessToken) (*pb.User, error) {
	tokenStr := req.Token

	username, err := jwt.ParseToken(tokenStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid token")
	}

	_, err = repository.GetUserIdByUsername(username)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "UserID not found")
	}

	// For demonstration, returning a user with the parsed username
	return &pb.User{Username: username}, nil
}

func main() {
	flag.Parse()

	// Initialize the database connection
	if _, err := repository.GetDB(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register the authServer with the gRPC server
	pb.RegisterAuthServiceServer(grpcServer, &authServer{})

	log.Printf("Starting gRPC server on :%d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
