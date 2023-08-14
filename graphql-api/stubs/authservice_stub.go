package stub

import "github.com/michael-m-truong/fitness-tracker/pb"

type StubAuthService struct {
}

func (service StubAuthService) CreateUser(req *pb.User) (*pb.NewUser, error) {
	newUser := &pb.NewUser{
		Username:    "TestUser",
		UserId:      1,
		AccessToken: "randomaccesstoken",
	}
	return newUser, nil
}

func (service StubAuthService) Login(userReq *pb.User) (*pb.AccessToken, error) {
	// Stub implementation for Login
	accessToken := &pb.AccessToken{
		// Initialize fields as needed
	}
	return accessToken, nil
}

func (service StubAuthService) ParseToken(token string) (*pb.User, error) {
	// Stub implementation for ParseToken
	user := &pb.User{
		// Initialize fields as needed
	}
	return user, nil
}
