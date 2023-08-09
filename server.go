package main

import (
	"example/graph"
	"example/graph/model"
	"log"
	"net/http"
	"os"
	// "strings"
	// "github.com/99designs/gqlgen/graphql/handler/extension"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql"
)

const defaultPort = "8080"

func main() {
	db, err := model.InitDB("root:lokeshpathrabe@tcp(localhost:3306)/blogger")
	if err != nil {
		log.Fatalf("failed to initialize the database: %v", err)
	}
	pubSub := graph.NewPubSubManager()
	resolver := graph.NewResolver(db, pubSub)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := graph.Config{ Resolvers: resolver}
	c.Directives.AllowQuery = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		// operation := strings.ToLower(graphql.GetOperationContext(ctx).Operation.Name)
		operation := graphql.GetFieldContext(ctx).Field.Name
		// operationContext := graphql.GetOperationContext(ctx)
		fmt.Println("lokesh",operation)
		allowedOperations := map[string]bool{
			// "allPersons":  true,
			"allPosts":    true,
			"personById":  true,
			"createPerson": true,
			"updatePerson": true,
			"deletePerson": true,
			// "createPost":   true,
			"updatePost":   true,
			"deletePost":   true,
		}
		if !allowedOperations[operation] && operation!="" {
			return nil, fmt.Errorf("operation not allowed: %s", operation)
		}
		return next(ctx)
	}

	// srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv := handler.New(graph.NewExecutableSchema(c))
	srv.AddTransport(transport.SSE{}) // <---- This is the important

	// default server
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	// srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
