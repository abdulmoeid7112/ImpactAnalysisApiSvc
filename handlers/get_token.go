package handlers

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/abdulmoeid7112/impact-analysis-api-svc"
	genModels "github.com/abdulmoeid7112/impact-analysis-api-svc/gen/models"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/gen/restapi/operations/service"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/models"
)

// NewServiceGetTokenHandler handles request for get token.
func NewServiceGetTokenHandler(ctx context.Context, rt *runtime.Runtime) service.GetTokenHandler {
	return &getToken{
		ctx: ctx,
		rt:  rt}
}

type getToken struct {
	ctx context.Context
	rt *runtime.Runtime
}

// Handle the get token request.
func (t *getToken) Handle(params service.GetTokenParams) middleware.Responder {
	token, err := t.rt.Service().GetToken(t.ctx, &models.User{
		Username: params.Username,
		TokenIssuesAt: time.Now(),
	})
	if err != nil {
		log().Errorf("failed to get token: %s", err)
		return service.NewGetTokenInternalServerError()
	}

	return service.NewGetTokenOK().WithPayload(&genModels.TokenResponse{
		Username: swag.String(params.Username),
		Token: swag.String(token),
	})
}
