package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Matterhorn-Apps/MatterhornApiService/auth"
	"github.com/Matterhorn-Apps/MatterhornApiService/database"
	"github.com/Matterhorn-Apps/MatterhornApiService/environment"
	"github.com/Matterhorn-Apps/MatterhornApiService/graph"
	"github.com/Matterhorn-Apps/MatterhornApiService/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Load environment variables
	environment.LoadEnv(".")

	// Connect to database
	db, err := database.DbConnect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Migrate database if necessary
	database.Migrate(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			DB: db,
		},
	}))

	jwtMiddleware := auth.BuildAuthenticationMiddleware()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", jwtMiddleware.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
