package stub

import (
	"github.com/michael-m-truong/fitness-tracker/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StubFailAuthService struct {
}

func (service StubFailAuthService) CreateUser(req *pb.User) (*pb.NewUser, error) {
	newUser := &pb.NewUser{
		Username:    "TestUser",
		UserId:      0,
		AccessToken: "randomaccesstoken",
	}
	return newUser, status.Errorf(codes.Internal, "Failed to create user")
}

func (service StubFailAuthService) Login(userReq *pb.User) (*pb.AccessToken, error) {
	// Stub implementation for Login
	accessToken := &pb.AccessToken{
		Token: "",
	}
	return accessToken, nil
}

func (service StubFailAuthService) ParseToken(token string) (*pb.User, error) {
	// Stub implementation for ParseToken
	user := &pb.User{
		// Initialize fields as needed
	}
	return user, nil
}
