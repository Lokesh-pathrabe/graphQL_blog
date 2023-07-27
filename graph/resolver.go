package graph

import (
	"database/sql"
	"example/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct{}
type Resolver struct {
	db     *sql.DB
	subscriber chan *model.Person
}

func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{db: db,subscriber: make(chan *model.Person)}
}

func (m *Resolver) Subscribe() <-chan *model.Person {
	ch := make(chan *model.Person)
	m.subscriber = ch
	return ch
}

func (m *Resolver) Publish(person *model.Person) {
	ch:=m.subscriber
	ch <- person
}
