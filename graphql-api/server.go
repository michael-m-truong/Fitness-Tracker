package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/michael-m-truong/fitness-tracker/graph"
	"github.com/michael-m-truong/fitness-tracker/middleware"
	"github.com/michael-m-truong/fitness-tracker/services"
)

const defaultPort = "8080"

func main() { //either manually go get all deps and generate or add -mod=mod -> go run -mod=mod github.com/99designs/gqlgen generate
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	//Init services
	authService := services.AuthResolverService{}
	exerciseService := services.ExerciseResolverService{}
	workoutService := services.WorkoutResolverService{}

	//Init middleware
	authMiddleware := middleware.AuthMiddleware{
		AuthService: authService,
	}

	router.Use(authMiddleware.Auth())

	// Pass your service to the resolver
	resolver := &graph.Resolver{
		AuthService:     authService,
		ExerciseService: exerciseService,
		WorkoutService:  workoutService,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
