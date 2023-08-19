package stub

import "github.com/michael-m-truong/fitness-tracker/pb"

type StubExerciseService struct {
}

func (service StubExerciseService) CreateExercise(req *pb.NewExercise) (*pb.Exercise, error) {
	newExercise := &pb.NewExercise{
		Name:        "Benchpress",
		Description: "Focus on form",
		MuscleGroup: "Chest",
		UserId:      1,
	}
	exercise := &pb.Exercise{
		Id:       123,
		Exercise: newExercise,
	}
	return exercise, nil
}

func (service StubExerciseService) CheckExerciseExists(*pb.ExerciseIds) (*pb.ExerciseExistence, error) {
	response := &pb.ExerciseExistence{
		Exists: true,
	}

	return response, nil
}
