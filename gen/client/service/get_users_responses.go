// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/* GetUsersOK describes a response with status code 200, with default header values.

user retrieved
*/
type GetUsersOK struct {
	Payload *GetUsersOKBody
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersOK  %+v", 200, o.Payload)
}
func (o *GetUsersOK) GetPayload() *GetUsersOKBody {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUsersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/* GetUsersUnauthorized describes a response with status code 401, with default header values.

Token is missing or invalid
*/
type GetUsersUnauthorized struct {
	WWWAuthenticate string
}

func (o *GetUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersUnauthorized ", 401)
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header WWW_Authenticate
	hdrWWWAuthenticate := response.GetHeader("WWW_Authenticate")

	if hdrWWWAuthenticate != "" {
		o.WWWAuthenticate = hdrWWWAuthenticate
	}

	return nil
}

// NewGetUsersInternalServerError creates a GetUsersInternalServerError with default headers values
func NewGetUsersInternalServerError() *GetUsersInternalServerError {
	return &GetUsersInternalServerError{}
}

/* GetUsersInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type GetUsersInternalServerError struct {
}

func (o *GetUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersInternalServerError ", 500)
}

func (o *GetUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetUsersOKBody get users o k body
swagger:model GetUsersOKBody
*/
type GetUsersOKBody struct {

	// users
	// Required: true
	Users []string `json:"users"`
}

// Validate validates this get users o k body
func (o *GetUsersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersOKBody) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("getUsersOK"+"."+"users", "body", o.Users); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get users o k body based on context it is used
func (o *GetUsersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersOKBody) UnmarshalBinary(b []byte) error {
	var res GetUsersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
