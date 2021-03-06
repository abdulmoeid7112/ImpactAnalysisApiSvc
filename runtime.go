package ImpactAnalysisApiSvc

import (
	"github.com/abdulmoeid7112/impact-analysis-api-svc/db"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/db/mongo"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/service"
)

// Runtime initializes values for entry point to our application.
type Runtime struct {
	svc *service.Service
}

// NewRuntime creates a new database service.
func NewRuntime() (*Runtime, error) {
	store, err := mongo.NewMongoClient(db.Option{})
	if err != nil {
		return nil, err
	}
	return &Runtime{svc: service.NewService(store)}, err
}

// Service returns pointer to our service variable.
func (r Runtime) Service() *service.Service {
	return r.svc
}
