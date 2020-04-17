package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tsuki42/graphql-meetup/graph/dataloader"
	"github.com/tsuki42/graphql-meetup/graph/resolver"
	"github.com/tsuki42/graphql-meetup/middleware"
	"github.com/tsuki42/graphql-meetup/postgres"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tsuki42/graphql-meetup/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler)

	userRepo := postgres.UserRepo{DB: postgres.Connection}
	router.Use(middleware.AuthMiddleware(userRepo))

	config := generated.Config{Resolvers: &resolver.Resolver{
		MeetupRepo: postgres.MeetupRepo{DB: postgres.Connection},
		UserRepo:   userRepo,
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", dataloader.DataloaderMiddleware(postgres.Connection, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
