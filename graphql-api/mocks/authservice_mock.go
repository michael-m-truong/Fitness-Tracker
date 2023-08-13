package mock

import "github.com/michael-m-truong/fitness-tracker/pb"

type MockAuthService struct {
}

func (service MockAuthService) CreateUser(req *pb.User) (*pb.NewUser, error) {
	newUser := &pb.NewUser{
		Username:    "TestUser",
		UserId:      1,
		AccessToken: "randomaccesstoken",
	}
	return newUser, nil
}

func (service MockAuthService) Login(userReq *pb.User) (*pb.AccessToken, error) {
	// Mock implementation for Login
	accessToken := &pb.AccessToken{
		// Initialize fields as needed
	}
	return accessToken, nil
}

func (service MockAuthService) ParseToken(token string) (*pb.User, error) {
	// Mock implementation for ParseToken
	user := &pb.User{
		// Initialize fields as needed
	}
	return user, nil
}
