package main

import (
	"log"
	"net/http"

	"go-graphql-mongodb/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	graph.ConnectMongo()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
