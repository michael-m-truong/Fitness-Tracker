package graph

import services "github.com/michael-m-truong/fitness-tracker/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthService     services.IAuthService
	ExerciseService services.IExerciseService
}
