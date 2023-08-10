package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/michael-m-truong/auth-grpc/jwt"
	pb "github.com/michael-m-truong/auth-grpc/pb" // Import the generated gRPC code

	// Import the dotenv package
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *authServer) CreateUser(ctx context.Context, req *pb.User) (*pb.NewUser, error) {
	// Create a dummy NewUser response
	newUser := &pb.NewUser{
		UserId:   123,          // Replace with actual user ID
		Username: req.Username, // Use the provided username from the request
	}

	// Return the dummy NewUser response
	return newUser, nil
}

func (s *authServer) GenerateToken(ctx context.Context, req *pb.User) (*pb.AccessToken, error) {
	username := req.Username

	tokenString, err := jwt.GenerateToken(username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token")
	}

	// For demonstration, returning a generated access token
	return &pb.AccessToken{Token: tokenString}, nil
}

func (s *authServer) ParseToken(ctx context.Context, req *pb.AccessToken) (*pb.User, error) {
	tokenStr := req.Token

	username, err := jwt.ParseToken(tokenStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid token")
	}

	// For demonstration, returning a user with the parsed username
	return &pb.User{Username: username}, nil
}

func GetUserIdByUsername(username string) (int32, error) {
	// Prepare the SQL query to retrieve user ID by username
	query := "SELECT id FROM user WHERE username = $1"
	row := db.QueryRow(query, username)

	var userID int32
	err := row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, status.Errorf(codes.NotFound, "User not found")
		}
		return 0, err
	}

	return userID, nil
}

func main() {
	flag.Parse()

	// Initialize the database connection
	// if err := initDB(); err != nil {
	// 	log.Fatal(err)
	// }

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
