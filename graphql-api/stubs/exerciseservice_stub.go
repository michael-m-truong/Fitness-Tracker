package stub

import "github.com/michael-m-truong/fitness-tracker/pb"

type StubExerciseService struct {
}

func (service StubExerciseService) CreateExercise(req *pb.NewExercise) (*pb.Exercise, error) {
	newExercise := &pb.NewExercise{
		Name:        "Push-up",
		Description: "Make sure to go to the floor or 90 degrees",
		MuscleGroup: "Chest",
		UserId:      1,
	}
	exercise := &pb.Exercise{
		Id:       123,
		Exercise: newExercise,
	}
	return exercise, nil
}
