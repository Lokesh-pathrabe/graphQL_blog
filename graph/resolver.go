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
	pubSub *PubSubManager
}


// type MyResolver struct {
// 	*Resolver
// }

// // Implement the graphql.Resolver interface for MyResolver
// func (r *MyResolver) AllPersons(ctx context.Context, last int) ([]*model.Person, error) {
// 	// Call the original resolver function on Resolver
// 	return r.Resolver.AllPersons(ctx, last)
// }

// func (r *MyResolver) CreatePerson(ctx context.Context, name string, age int) (*Person, error) {
// 	// Call the original resolver function on Resolver
// 	return r.Resolver.CreatePerson(ctx, name, age)
// }

func NewResolver(db *sql.DB, pubSub *PubSubManager) *Resolver {
	return &Resolver{db: db, pubSub: pubSub}
}

type PubSubManager struct {
	subscriberperson chan *model.Person
	subscriberpost chan *model.Post
}

func NewPubSubManager() *PubSubManager {
	return &PubSubManager{
		subscriberperson: make(chan *model.Person),
	}
}

func (m *PubSubManager) SubscribePerson() chan *model.Person {
	ch := make(chan *model.Person)
	m.subscriberperson = ch
	return ch
}

func (m *PubSubManager) UnsubscribePerson(ch chan *model.Person) {
	close(ch)
}

func (m *PubSubManager) PublishPerson(person *model.Person) {
	ch := m.subscriberperson
	ch <- person
}
func (m *PubSubManager) SubscribePost() chan *model.Post {
	ch := make(chan *model.Post)
	m.subscriberpost = ch
	return ch
}

func (m *PubSubManager) UnsubscribePost(ch chan *model.Post) {
	close(ch)
}

func (m *PubSubManager) PublishPost(post *model.Post) {
	ch := m.subscriberpost
	ch <- post
}

// type Resolver struct {
// 	db     *sql.DB
// 	subscriber chan *model.Person
// }

// func NewResolver(db *sql.DB) *Resolver {
// 	return &Resolver{db: db,subscriber: make(chan *model.Person)}
// }

// func (m *Resolver) Subscribe() <-chan *model.Person {
// 	ch := make(chan *model.Person)
// 	m.subscriber = ch
// 	return ch
// }

// func (m *Resolver) Publish(person *model.Person) {
// 	ch:=m.subscriber
// 	ch <- person
// }
