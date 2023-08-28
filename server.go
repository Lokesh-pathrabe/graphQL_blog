package main

import (
	"example/graph"
	"example/graph/model"
	"log"
	"net/http"
	"os"
	// "strings"
	auth "example/internal/auth"
	// "github.com/99designs/gqlgen/graphql/handler/extension"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	db, err := model.InitDB("root:lokeshpathrabe@tcp(localhost:3306)/blog")
	if err != nil {
		log.Fatalf("failed to initialize the database: %v", err)
	}
	pubSub := graph.NewPubSubManager()
	resolver := graph.NewResolver(db, pubSub)

	c := graph.Config{ Resolvers: resolver}
	c.Directives.AllowQuery = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		operation := graphql.GetFieldContext(ctx).Field.Name
		fmt.Println("operation: ",operation)
		allowedOperations := map[string]bool{
			"allPersons":  true,
			"allPosts":    true,
			"personById":  true,
			"createPerson": true,
			"updatePerson": true,
			"deletePerson": true,
			"createPost":   true,
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
	// srv.Use(extension.FixedComplexityLimit(3))
	
	
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
