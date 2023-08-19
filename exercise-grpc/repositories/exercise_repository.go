package repository

import (
	"context"

	"github.com/michael-m-truong/exercise-grpc/pb"
)

func CreateExercise(ctx context.Context, exercise *pb.NewExercise) (*int, error) {
	// Get DB instance
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	// Prepare the INSERT query with placeholders
	// TODO: convert user_id to integer for exercise table
	insertQuery := `
		INSERT INTO exercise (name, description, muscle_group, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	// Execute the INSERT query with parameters and retrieve the inserted exercise's ID
	var insertedID *int
	err = db.QueryRowContext(ctx, insertQuery, exercise.Name, exercise.Description, exercise.MuscleGroup, exercise.UserId).Scan(&insertedID)
	if err != nil {
		return nil, err
	}

	// Update the request with the inserted ID
	return insertedID, nil
}
