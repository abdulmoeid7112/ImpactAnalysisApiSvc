package db

import (
	"context"
	"log"

	"github.com/abdulmoeid7112/impact-analysis-api-svc/models"
)

// DataStore is an interface for query ops.
type DataStore interface {
	AddUser(context context.Context, user *models.User) (*models.User, error)
	GetUserByName(context context.Context, name string) (*models.User, error)
	ListUser(context context.Context) ([]*models.User, error)
}

// Option holds configuration for data store clients.
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store.
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory.
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
}
