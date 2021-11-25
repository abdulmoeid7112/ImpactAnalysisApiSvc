package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abdulmoeid7112/impact-analysis-api-svc"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/gen/restapi/operations/service"
)

// NewServiceGetUsersHandler handles request for get users.
func NewServiceGetUsersHandler(ctx context.Context, rt *runtime.Runtime) service.GetUsersHandler {
	return &getUsers{
		ctx: ctx,
		rt:  rt}
}

type getUsers struct {
	ctx context.Context
	rt *runtime.Runtime
}

// Handle the add student request.
func (u *getUsers) Handle(_ service.GetUsersParams, _ interface{}) middleware.Responder {
	usersList, err := u.rt.Service().GetUsers(u.ctx)
	if err != nil {
		log().Errorf("failed to get users list: %s", err)
		return service.NewGetUsersInternalServerError()
	}

	return service.NewGetUsersOK().WithPayload(
		&service.GetUsersOKBody{
			Users: usersList,
	})
}
