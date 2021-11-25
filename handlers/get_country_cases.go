package handlers

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abdulmoeid7112/impact-analysis-api-svc"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/gen/restapi/operations/service"
)

// NewServiceGetCountryCasesHandler handles request for get country cases reported today.
func NewServiceGetCountryCasesHandler(ctx context.Context, rt *runtime.Runtime) service.GetCountryCasesHandler {
	return &getCountryCases{
		ctx: ctx,
		rt:  rt}
}

type getCountryCases struct {
	ctx context.Context
	rt *runtime.Runtime
}

// Handle the get country cases request.
func (c *getCountryCases) Handle(params service.GetCountryCasesParams, _ interface{}) middleware.Responder {
	result, err := c.rt.Service().GetCountryCases(mkFilter(params))
	if err != nil {
		log().Errorf("failed to get country cases: %s", err)
		return service.NewGetCountryCasesInternalServerError()
	}


	return service.NewGetCountryCasesOK().WithPayload(result)
}

func mkFilter(params service.GetCountryCasesParams) map[string]interface{}  {
	filter := make(map[string]interface{})
	if params.CountryName != nil && params.DateFrom != nil{
		filter["country_name"] = params.CountryName
		filter["date_from"] = params.DateFrom
		filter["type"] = "date"

		return filter
	}

	if params.CountryName != nil {
		fmt.Println("idr gaya")
		filter["country_name"] = params.CountryName
		filter["type"] = "country"

		return filter
	}

	filter["type"] = "all"
	return filter
}
