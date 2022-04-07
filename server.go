package main

import (
	"graphql-subscriptions/graph"
	"graphql-subscriptions/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		panic(err)
	}

	defer nc.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Nats: nc}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
