package repository

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckLoginInfo(ctx context.Context, username string, passwordHash string) (*string, error) {
	// Access the singleton database instance
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	query := "SELECT password_hash FROM users WHERE username = $1"
	row := db.QueryRow(query, username)

	var hashedPassword *string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Invalid username or password")
		}
		return nil, status.Errorf(codes.Internal, "Database error")
	}

	return hashedPassword, nil

}

func CreateUser(ctx context.Context, username string, passwordHash string) (*int32, error) {
	// // Access the singleton database instance
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	insertQuery := `
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
		RETURNING id
	`

	var userID *int32
	err = db.QueryRowContext(ctx, insertQuery, username, passwordHash).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func GetUserIdByUsername(username string) (*int32, error) {
	// Access the singleton database instance
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	// Prepare the SQL query to retrieve user ID by username
	query := "SELECT id FROM user WHERE username = $1"
	row := db.QueryRow(query, username)

	var userID int32
	err = row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, err
	}

	return &userID, nil
}
