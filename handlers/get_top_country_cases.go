package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abdulmoeid7112/impact-analysis-api-svc"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/gen/restapi/operations/service"
)

// NewServiceGetTopCasesCountriesHandler handles request for get top Ncountry cases reported today.
func NewServiceGetTopCasesCountriesHandler(ctx context.Context, rt *runtime.Runtime) service.GetTopCasesCountriesHandler {
	return &getTopCountryCases{
		ctx: ctx,
		rt:  rt}
}

type getTopCountryCases struct {
	ctx context.Context
	rt *runtime.Runtime
}

// Handle the get country cases request.
func (c *getTopCountryCases) Handle(params service.GetTopCasesCountriesParams, _ interface{}) middleware.Responder {
	result, err := c.rt.Service().GetTopCasesCountries(params.Count)
	if err != nil {
		log().Errorf("failed to get top country cases: %s", err)
		return service.NewGetTopCasesCountriesInternalServerError()
	}


	return service.NewGetTopCasesCountriesOK().WithPayload(result)
}
