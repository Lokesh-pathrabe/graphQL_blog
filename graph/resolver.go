package graph

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct{}
type Resolver struct {
	// MongoDB client and database.
	client   *mongo.Client
	database *mongo.Database
}

func NewResolver() (*Resolver, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://lokeshpathrabe:lokeshpathrabe@cluster0.mtmxz6b.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		return nil, err
	}
	database := client.Database("blogger")

	return &Resolver{
		client:   client,
		database: database,
	}, nil
}
