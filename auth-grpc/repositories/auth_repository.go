package respository

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckLoginInfo(ctx context.Context, username string, passwordHash string) (*string, error) {
	db := GetDB() // Access the singleton database instance

	query := "SELECT password_hash FROM users WHERE username = $1"
	row := db.QueryRow(query, username)

	var hashedPassword *string
	err := row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Invalid username or password")
		}
		return nil, status.Errorf(codes.Internal, "Database error")
	}

	return hashedPassword, nil

}

func CreateUser(ctx context.Context, username string, passwordHash string) (*int32, error) {
	// Execute the INSERT query to create a new user
	db := GetDB() // Access the singleton database instance
	insertQuery := `
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
		RETURNING id
	`

	var userID *int32
	err := db.QueryRowContext(ctx, insertQuery, username, passwordHash).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return userID, nil
}
