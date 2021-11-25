// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetCountryCasesReader is a Reader for the GetCountryCases structure.
type GetCountryCasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCountryCasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCountryCasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCountryCasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCountryCasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCountryCasesOK creates a GetCountryCasesOK with default headers values
func NewGetCountryCasesOK() *GetCountryCasesOK {
	return &GetCountryCasesOK{}
}

/* GetCountryCasesOK describes a response with status code 200, with default header values.

country covid cases retrieved
*/
type GetCountryCasesOK struct {
	Payload []interface{}
}

func (o *GetCountryCasesOK) Error() string {
	return fmt.Sprintf("[GET /country-cases][%d] getCountryCasesOK  %+v", 200, o.Payload)
}
func (o *GetCountryCasesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetCountryCasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCountryCasesUnauthorized creates a GetCountryCasesUnauthorized with default headers values
func NewGetCountryCasesUnauthorized() *GetCountryCasesUnauthorized {
	return &GetCountryCasesUnauthorized{}
}

/* GetCountryCasesUnauthorized describes a response with status code 401, with default header values.

Token is missing or invalid
*/
type GetCountryCasesUnauthorized struct {
	WWWAuthenticate string
}

func (o *GetCountryCasesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /country-cases][%d] getCountryCasesUnauthorized ", 401)
}

func (o *GetCountryCasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header WWW_Authenticate
	hdrWWWAuthenticate := response.GetHeader("WWW_Authenticate")

	if hdrWWWAuthenticate != "" {
		o.WWWAuthenticate = hdrWWWAuthenticate
	}

	return nil
}

// NewGetCountryCasesInternalServerError creates a GetCountryCasesInternalServerError with default headers values
func NewGetCountryCasesInternalServerError() *GetCountryCasesInternalServerError {
	return &GetCountryCasesInternalServerError{}
}

/* GetCountryCasesInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type GetCountryCasesInternalServerError struct {
}

func (o *GetCountryCasesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /country-cases][%d] getCountryCasesInternalServerError ", 500)
}

func (o *GetCountryCasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
