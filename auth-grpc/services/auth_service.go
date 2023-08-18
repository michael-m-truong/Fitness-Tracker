package service

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func Login(ctx context.Context, req *pb.User) (*pb.AccessToken, error) {
// 	fmt.Printf("Received Login: %+v\n", req)
// 	db := repository.GetDB() // Access the singleton database instance

// 	// Implement your login logic here
// 	username := req.Username
// 	password := req.Password

// 	// Prepare the SQL query to retrieve the hashed password by username
// 	query := "SELECT password_hash FROM users WHERE username = $1"
// 	row := db.QueryRow(query, username)

// 	var hashedPassword string
// 	err := row.Scan(&hashedPassword)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, status.Errorf(codes.NotFound, "Invalid username or password")
// 		}
// 		return nil, status.Errorf(codes.Internal, "Database error")
// 	}

// 	// Check if the provided password matches the hashed password
// 	if !helper.CheckPasswordHash(password, hashedPassword) {
// 		return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
// 	}

// 	// Generate and return an access token
// 	tokenString, err := jwt.GenerateToken(username)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to generate token")
// 	}

// 	return &pb.AccessToken{Token: tokenString}, nil
// }

// func CreateUser(ctx context.Context, req *pb.User) (*pb.NewUser, error) {
// 	fmt.Printf("Received Exercise: %+v\n", req)

// 	// Hash the password
// 	hashedPassword, err := helper.HashPassword(req.Password)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to hash password")
// 	}

// 	// Execute the INSERT query to create a new user
// 	db := repository.GetDB() // Access the singleton database instance
// 	insertQuery := `
// 		INSERT INTO users (username, password_hash)
// 		VALUES ($1, $2)
// 		RETURNING id
// 	`

// 	var userID int32
// 	err = db.QueryRowContext(ctx, insertQuery, req.Username, hashedPassword).Scan(&userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tokenString, err := jwt.GenerateToken(req.Username)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to generate token")
// 	}

// 	// Create a NewUser response
// 	newUser := &pb.NewUser{
// 		UserId:      userID,
// 		Username:    req.Username,
// 		AccessToken: tokenString,
// 	}

// 	return newUser, nil
// }

// func ParseToken(ctx context.Context, req *pb.AccessToken) (*pb.User, error) {
// 	tokenStr := req.Token

// 	username, err := jwt.ParseToken(tokenStr)
// 	if err != nil {
// 		return nil, status.Errorf(codes.InvalidArgument, "Invalid token")
// 	}

// 	// For demonstration, returning a user with the parsed username
// 	return &pb.User{Username: username}, nil
// }
