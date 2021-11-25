// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTopCasesCountriesHandlerFunc turns a function with the right signature into a get top cases countries handler
type GetTopCasesCountriesHandlerFunc func(GetTopCasesCountriesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTopCasesCountriesHandlerFunc) Handle(params GetTopCasesCountriesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTopCasesCountriesHandler interface for that can handle valid get top cases countries params
type GetTopCasesCountriesHandler interface {
	Handle(GetTopCasesCountriesParams, interface{}) middleware.Responder
}

// NewGetTopCasesCountries creates a new http.Handler for the get top cases countries operation
func NewGetTopCasesCountries(ctx *middleware.Context, handler GetTopCasesCountriesHandler) *GetTopCasesCountries {
	return &GetTopCasesCountries{Context: ctx, Handler: handler}
}

/* GetTopCasesCountries swagger:route GET /top-cases-countries service getTopCasesCountries

retrieve top N countries with most reported cases today

*/
type GetTopCasesCountries struct {
	Context *middleware.Context
	Handler GetTopCasesCountriesHandler
}

func (o *GetTopCasesCountries) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTopCasesCountriesParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
