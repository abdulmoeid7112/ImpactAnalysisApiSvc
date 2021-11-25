package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/abdulmoeid7112/impact-analysis-api-svc"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/gen/restapi/operations/service"
)

// NewServiceGetAllCasesHandler handles request for get all cases reported today.
func NewServiceGetAllCasesHandler(ctx context.Context, rt *runtime.Runtime) service.GetAllCasesHandler {
	return &getAllCases{
		ctx: ctx,
		rt:  rt}
}

type getAllCases struct {
	ctx context.Context
	rt *runtime.Runtime
}

// Handle the get all cases request.
func (c *getAllCases) Handle(_ service.GetAllCasesParams, _ interface{}) middleware.Responder {
	todayDate, totalCases, err := c.rt.Service().GetAllCases()
	if err != nil {
		log().Errorf("failed to get all cases reported today: %s", err)
		return service.NewGetAllCasesInternalServerError()
	}

	return service.NewGetAllCasesOK().WithPayload(
		&service.GetAllCasesOKBody{
			TodayDate:swag.String(todayDate),
			TotalCases: swag.String(totalCases),
		})
}
