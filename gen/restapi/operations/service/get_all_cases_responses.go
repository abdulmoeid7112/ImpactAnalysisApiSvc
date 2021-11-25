// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAllCasesOKCode is the HTTP code returned for type GetAllCasesOK
const GetAllCasesOKCode int = 200

/*GetAllCasesOK covid cases retrieved

swagger:response getAllCasesOK
*/
type GetAllCasesOK struct {

	/*
	  In: Body
	*/
	Payload *GetAllCasesOKBody `json:"body,omitempty"`
}

// NewGetAllCasesOK creates GetAllCasesOK with default headers values
func NewGetAllCasesOK() *GetAllCasesOK {

	return &GetAllCasesOK{}
}

// WithPayload adds the payload to the get all cases o k response
func (o *GetAllCasesOK) WithPayload(payload *GetAllCasesOKBody) *GetAllCasesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all cases o k response
func (o *GetAllCasesOK) SetPayload(payload *GetAllCasesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllCasesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllCasesUnauthorizedCode is the HTTP code returned for type GetAllCasesUnauthorized
const GetAllCasesUnauthorizedCode int = 401

/*GetAllCasesUnauthorized Token is missing or invalid

swagger:response getAllCasesUnauthorized
*/
type GetAllCasesUnauthorized struct {
	/*

	 */
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewGetAllCasesUnauthorized creates GetAllCasesUnauthorized with default headers values
func NewGetAllCasesUnauthorized() *GetAllCasesUnauthorized {

	return &GetAllCasesUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the get all cases unauthorized response
func (o *GetAllCasesUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *GetAllCasesUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the get all cases unauthorized response
func (o *GetAllCasesUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *GetAllCasesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetAllCasesInternalServerErrorCode is the HTTP code returned for type GetAllCasesInternalServerError
const GetAllCasesInternalServerErrorCode int = 500

/*GetAllCasesInternalServerError internal server error

swagger:response getAllCasesInternalServerError
*/
type GetAllCasesInternalServerError struct {
}

// NewGetAllCasesInternalServerError creates GetAllCasesInternalServerError with default headers values
func NewGetAllCasesInternalServerError() *GetAllCasesInternalServerError {

	return &GetAllCasesInternalServerError{}
}

// WriteResponse to the client
func (o *GetAllCasesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
