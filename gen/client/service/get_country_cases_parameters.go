// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetCountryCasesParams creates a new GetCountryCasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCountryCasesParams() *GetCountryCasesParams {
	return &GetCountryCasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCountryCasesParamsWithTimeout creates a new GetCountryCasesParams object
// with the ability to set a timeout on a request.
func NewGetCountryCasesParamsWithTimeout(timeout time.Duration) *GetCountryCasesParams {
	return &GetCountryCasesParams{
		timeout: timeout,
	}
}

// NewGetCountryCasesParamsWithContext creates a new GetCountryCasesParams object
// with the ability to set a context for a request.
func NewGetCountryCasesParamsWithContext(ctx context.Context) *GetCountryCasesParams {
	return &GetCountryCasesParams{
		Context: ctx,
	}
}

// NewGetCountryCasesParamsWithHTTPClient creates a new GetCountryCasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCountryCasesParamsWithHTTPClient(client *http.Client) *GetCountryCasesParams {
	return &GetCountryCasesParams{
		HTTPClient: client,
	}
}

/* GetCountryCasesParams contains all the parameters to send to the API endpoint
   for the get country cases operation.

   Typically these are written to a http.Request.
*/
type GetCountryCasesParams struct {

	// CountryName.
	CountryName *string

	// DateFrom.
	DateFrom *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get country cases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCountryCasesParams) WithDefaults() *GetCountryCasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get country cases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCountryCasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get country cases params
func (o *GetCountryCasesParams) WithTimeout(timeout time.Duration) *GetCountryCasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get country cases params
func (o *GetCountryCasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get country cases params
func (o *GetCountryCasesParams) WithContext(ctx context.Context) *GetCountryCasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get country cases params
func (o *GetCountryCasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get country cases params
func (o *GetCountryCasesParams) WithHTTPClient(client *http.Client) *GetCountryCasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get country cases params
func (o *GetCountryCasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCountryName adds the countryName to the get country cases params
func (o *GetCountryCasesParams) WithCountryName(countryName *string) *GetCountryCasesParams {
	o.SetCountryName(countryName)
	return o
}

// SetCountryName adds the countryName to the get country cases params
func (o *GetCountryCasesParams) SetCountryName(countryName *string) {
	o.CountryName = countryName
}

// WithDateFrom adds the dateFrom to the get country cases params
func (o *GetCountryCasesParams) WithDateFrom(dateFrom *string) *GetCountryCasesParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the get country cases params
func (o *GetCountryCasesParams) SetDateFrom(dateFrom *string) {
	o.DateFrom = dateFrom
}

// WriteToRequest writes these params to a swagger request
func (o *GetCountryCasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CountryName != nil {

		// query param country_name
		var qrCountryName string

		if o.CountryName != nil {
			qrCountryName = *o.CountryName
		}
		qCountryName := qrCountryName
		if qCountryName != "" {

			if err := r.SetQueryParam("country_name", qCountryName); err != nil {
				return err
			}
		}
	}

	if o.DateFrom != nil {

		// query param date_from
		var qrDateFrom string

		if o.DateFrom != nil {
			qrDateFrom = *o.DateFrom
		}
		qDateFrom := qrDateFrom
		if qDateFrom != "" {

			if err := r.SetQueryParam("date_from", qDateFrom); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
