// Code generated by go-swagger; DO NOT EDIT.

package platform

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models"
)

// RegisterApplicationReader is a Reader for the RegisterApplication structure.
type RegisterApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRegisterApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRegisterApplicationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications] registerApplication", response, response.Code())
	}
}

// NewRegisterApplicationOK creates a RegisterApplicationOK with default headers values
func NewRegisterApplicationOK() *RegisterApplicationOK {
	return &RegisterApplicationOK{}
}

/*
RegisterApplicationOK describes a response with status code 200, with default header values.

Application registration operation details
*/
type RegisterApplicationOK struct {
	Payload *models.ApplicationRegistrationUpsertResponse
}

// IsSuccess returns true when this register application o k response has a 2xx status code
func (o *RegisterApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this register application o k response has a 3xx status code
func (o *RegisterApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register application o k response has a 4xx status code
func (o *RegisterApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this register application o k response has a 5xx status code
func (o *RegisterApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this register application o k response a status code equal to that given
func (o *RegisterApplicationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the register application o k response
func (o *RegisterApplicationOK) Code() int {
	return 200
}

func (o *RegisterApplicationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications][%d] registerApplicationOK %s", 200, payload)
}

func (o *RegisterApplicationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications][%d] registerApplicationOK %s", 200, payload)
}

func (o *RegisterApplicationOK) GetPayload() *models.ApplicationRegistrationUpsertResponse {
	return o.Payload
}

func (o *RegisterApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationRegistrationUpsertResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterApplicationBadRequest creates a RegisterApplicationBadRequest with default headers values
func NewRegisterApplicationBadRequest() *RegisterApplicationBadRequest {
	return &RegisterApplicationBadRequest{}
}

/*
RegisterApplicationBadRequest describes a response with status code 400, with default header values.

Invalid application registration
*/
type RegisterApplicationBadRequest struct {
}

// IsSuccess returns true when this register application bad request response has a 2xx status code
func (o *RegisterApplicationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this register application bad request response has a 3xx status code
func (o *RegisterApplicationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register application bad request response has a 4xx status code
func (o *RegisterApplicationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this register application bad request response has a 5xx status code
func (o *RegisterApplicationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this register application bad request response a status code equal to that given
func (o *RegisterApplicationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the register application bad request response
func (o *RegisterApplicationBadRequest) Code() int {
	return 400
}

func (o *RegisterApplicationBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationBadRequest", 400)
}

func (o *RegisterApplicationBadRequest) String() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationBadRequest", 400)
}

func (o *RegisterApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterApplicationUnauthorized creates a RegisterApplicationUnauthorized with default headers values
func NewRegisterApplicationUnauthorized() *RegisterApplicationUnauthorized {
	return &RegisterApplicationUnauthorized{}
}

/*
RegisterApplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RegisterApplicationUnauthorized struct {
}

// IsSuccess returns true when this register application unauthorized response has a 2xx status code
func (o *RegisterApplicationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this register application unauthorized response has a 3xx status code
func (o *RegisterApplicationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register application unauthorized response has a 4xx status code
func (o *RegisterApplicationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this register application unauthorized response has a 5xx status code
func (o *RegisterApplicationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this register application unauthorized response a status code equal to that given
func (o *RegisterApplicationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the register application unauthorized response
func (o *RegisterApplicationUnauthorized) Code() int {
	return 401
}

func (o *RegisterApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationUnauthorized", 401)
}

func (o *RegisterApplicationUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationUnauthorized", 401)
}

func (o *RegisterApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterApplicationConflict creates a RegisterApplicationConflict with default headers values
func NewRegisterApplicationConflict() *RegisterApplicationConflict {
	return &RegisterApplicationConflict{}
}

/*
RegisterApplicationConflict describes a response with status code 409, with default header values.

Conflict
*/
type RegisterApplicationConflict struct {
}

// IsSuccess returns true when this register application conflict response has a 2xx status code
func (o *RegisterApplicationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this register application conflict response has a 3xx status code
func (o *RegisterApplicationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register application conflict response has a 4xx status code
func (o *RegisterApplicationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this register application conflict response has a 5xx status code
func (o *RegisterApplicationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this register application conflict response a status code equal to that given
func (o *RegisterApplicationConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the register application conflict response
func (o *RegisterApplicationConflict) Code() int {
	return 409
}

func (o *RegisterApplicationConflict) Error() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationConflict", 409)
}

func (o *RegisterApplicationConflict) String() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationConflict", 409)
}

func (o *RegisterApplicationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
