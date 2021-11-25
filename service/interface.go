package service

import (
	"github.com/abdulmoeid7112/impact-analysis-api-svc/client"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/db"
)

// Service initializes our database instance.
type Service struct {
	db        db.DataStore
	apiClient *client.DatasetApiClient
}

// NewService creates a connection to our database.
func NewService(ds db.DataStore) *Service {
	return &Service{
		db:        ds,
		apiClient: client.NewDatasetApiClient(),
	}
}
