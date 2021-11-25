package handlers

import (
	"context"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	wraperrors "github.com/pkg/errors"

	runtime "github.com/abdulmoeid7112/impact-analysis-api-svc"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/gen/restapi/operations"
)

// Handler replaces swagger handler.
type Handler *operations.ImpactAnalysisAPI

// NewHandler overrides swagger api handlers.
func NewHandler(ctx context.Context, rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewImpactAnalysisAPI(spec)

	// student handlers.
	handler.ServiceGetTokenHandler = NewServiceGetTokenHandler(ctx, rt)
	handler.ServiceGetUsersHandler = NewServiceGetUsersHandler(ctx, rt)
	handler.ServiceGetAllCasesHandler = NewServiceGetAllCasesHandler(ctx, rt)
	handler.ServiceGetCountryCasesHandler = NewServiceGetCountryCasesHandler(ctx, rt)
	handler.ServiceGetTopCasesCountriesHandler = NewServiceGetTopCasesCountriesHandler(ctx, rt)

	handler.BearerTokenAuth = func(token string) (interface{}, error) {
		if err := validateToken(rt, ctx, token); err != nil {
			return nil, errors.New(401, "invalid bearer token")
		}

		return token, nil
	}

	return handler
}

func validateToken(rt *runtime.Runtime, ctx context.Context, tokenString string) error {
	token, _, err := new(jwt.Parser).ParseUnverified(strings.Fields(tokenString)[1], jwt.MapClaims{})
	if err != nil {
		return wraperrors.Wrap(err, "failed to parse token string")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expireTime, timeErr := time.Parse(time.RFC3339, claims["expired_at"].(string))
		if timeErr != nil {
			return wraperrors.Wrap(timeErr, "error while parsing time")
		}

		if okBool := time.Now().After(expireTime); okBool {
			return wraperrors.Wrap(timeErr, "token has expired")
		}

	}

	return nil
}
