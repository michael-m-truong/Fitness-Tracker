package repository

import (
	"context"

	"github.com/michael-m-truong/workout-grpc/pb"
)

func CreateWorkout(ctx context.Context, workout *pb.NewWorkout) (*int, error) {
	// Get DB instance
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	// Prepare the INSERT query with placeholders
	// TODO: convert user_id to integer for workout table
	insertQuery := `
		INSERT INTO workout (name, description, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	// Execute the INSERT query with parameters and retrieve the inserted workout's ID
	var insertedID *int
	err = db.QueryRowContext(ctx, insertQuery, workout.Name, workout.Description, workout.UserId).Scan(&insertedID)
	if err != nil {
		return nil, err
	}

	// Update the request with the inserted ID
	return insertedID, nil
}

func AddExercisesToWorkout(ctx context.Context, workoutID int, exerciseIDs []int32) error {
	// Get DB instance
	db, err := GetDB()
	if err != nil {
		return err
	}

	// Prepare the INSERT query for adding exercises to the workout
	insertQuery := `
		INSERT INTO workout_exercise (workout_id, exercise_id)
		VALUES ($1, $2)
	`

	// Insert each exercise into the workout_exercise table
	for _, exerciseID := range exerciseIDs {
		_, err := db.ExecContext(ctx, insertQuery, workoutID, exerciseID)
		if err != nil {
			return err
		}
	}

	return nil
}
