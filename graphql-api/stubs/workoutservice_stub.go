package stub

import "github.com/michael-m-truong/fitness-tracker/pb"

type StubWorkoutService struct {
}

func (service StubWorkoutService) CreateWorkout(req *pb.NewWorkout) (*pb.Workout, error) {
	newWorkout := &pb.NewWorkout{
		Name:        "Benchpress",
		Description: "Focus on form",
		UserId:      1,
	}
	workout := &pb.Workout{
		Id:      12345,
		Workout: newWorkout,
	}
	return workout, nil
}
