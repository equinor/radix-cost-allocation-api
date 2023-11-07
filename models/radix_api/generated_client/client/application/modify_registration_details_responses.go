// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models"
)

// ModifyRegistrationDetailsReader is a Reader for the ModifyRegistrationDetails structure.
type ModifyRegistrationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyRegistrationDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyRegistrationDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewModifyRegistrationDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewModifyRegistrationDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewModifyRegistrationDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewModifyRegistrationDetailsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /applications/{appName}] modifyRegistrationDetails", response, response.Code())
	}
}

// NewModifyRegistrationDetailsOK creates a ModifyRegistrationDetailsOK with default headers values
func NewModifyRegistrationDetailsOK() *ModifyRegistrationDetailsOK {
	return &ModifyRegistrationDetailsOK{}
}

/*
ModifyRegistrationDetailsOK describes a response with status code 200, with default header values.

Modifying registration operation details
*/
type ModifyRegistrationDetailsOK struct {
	Payload *models.ApplicationRegistrationUpsertResponse
}

// IsSuccess returns true when this modify registration details o k response has a 2xx status code
func (o *ModifyRegistrationDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify registration details o k response has a 3xx status code
func (o *ModifyRegistrationDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify registration details o k response has a 4xx status code
func (o *ModifyRegistrationDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify registration details o k response has a 5xx status code
func (o *ModifyRegistrationDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify registration details o k response a status code equal to that given
func (o *ModifyRegistrationDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the modify registration details o k response
func (o *ModifyRegistrationDetailsOK) Code() int {
	return 200
}

func (o *ModifyRegistrationDetailsOK) Error() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsOK  %+v", 200, o.Payload)
}

func (o *ModifyRegistrationDetailsOK) String() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsOK  %+v", 200, o.Payload)
}

func (o *ModifyRegistrationDetailsOK) GetPayload() *models.ApplicationRegistrationUpsertResponse {
	return o.Payload
}

func (o *ModifyRegistrationDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationRegistrationUpsertResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyRegistrationDetailsBadRequest creates a ModifyRegistrationDetailsBadRequest with default headers values
func NewModifyRegistrationDetailsBadRequest() *ModifyRegistrationDetailsBadRequest {
	return &ModifyRegistrationDetailsBadRequest{}
}

/*
ModifyRegistrationDetailsBadRequest describes a response with status code 400, with default header values.

Invalid application
*/
type ModifyRegistrationDetailsBadRequest struct {
}

// IsSuccess returns true when this modify registration details bad request response has a 2xx status code
func (o *ModifyRegistrationDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify registration details bad request response has a 3xx status code
func (o *ModifyRegistrationDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify registration details bad request response has a 4xx status code
func (o *ModifyRegistrationDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify registration details bad request response has a 5xx status code
func (o *ModifyRegistrationDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this modify registration details bad request response a status code equal to that given
func (o *ModifyRegistrationDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the modify registration details bad request response
func (o *ModifyRegistrationDetailsBadRequest) Code() int {
	return 400
}

func (o *ModifyRegistrationDetailsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsBadRequest ", 400)
}

func (o *ModifyRegistrationDetailsBadRequest) String() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsBadRequest ", 400)
}

func (o *ModifyRegistrationDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyRegistrationDetailsUnauthorized creates a ModifyRegistrationDetailsUnauthorized with default headers values
func NewModifyRegistrationDetailsUnauthorized() *ModifyRegistrationDetailsUnauthorized {
	return &ModifyRegistrationDetailsUnauthorized{}
}

/*
ModifyRegistrationDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ModifyRegistrationDetailsUnauthorized struct {
}

// IsSuccess returns true when this modify registration details unauthorized response has a 2xx status code
func (o *ModifyRegistrationDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify registration details unauthorized response has a 3xx status code
func (o *ModifyRegistrationDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify registration details unauthorized response has a 4xx status code
func (o *ModifyRegistrationDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify registration details unauthorized response has a 5xx status code
func (o *ModifyRegistrationDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this modify registration details unauthorized response a status code equal to that given
func (o *ModifyRegistrationDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the modify registration details unauthorized response
func (o *ModifyRegistrationDetailsUnauthorized) Code() int {
	return 401
}

func (o *ModifyRegistrationDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsUnauthorized ", 401)
}

func (o *ModifyRegistrationDetailsUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsUnauthorized ", 401)
}

func (o *ModifyRegistrationDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyRegistrationDetailsNotFound creates a ModifyRegistrationDetailsNotFound with default headers values
func NewModifyRegistrationDetailsNotFound() *ModifyRegistrationDetailsNotFound {
	return &ModifyRegistrationDetailsNotFound{}
}

/*
ModifyRegistrationDetailsNotFound describes a response with status code 404, with default header values.

Not found
*/
type ModifyRegistrationDetailsNotFound struct {
}

// IsSuccess returns true when this modify registration details not found response has a 2xx status code
func (o *ModifyRegistrationDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify registration details not found response has a 3xx status code
func (o *ModifyRegistrationDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify registration details not found response has a 4xx status code
func (o *ModifyRegistrationDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify registration details not found response has a 5xx status code
func (o *ModifyRegistrationDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this modify registration details not found response a status code equal to that given
func (o *ModifyRegistrationDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the modify registration details not found response
func (o *ModifyRegistrationDetailsNotFound) Code() int {
	return 404
}

func (o *ModifyRegistrationDetailsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsNotFound ", 404)
}

func (o *ModifyRegistrationDetailsNotFound) String() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsNotFound ", 404)
}

func (o *ModifyRegistrationDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyRegistrationDetailsConflict creates a ModifyRegistrationDetailsConflict with default headers values
func NewModifyRegistrationDetailsConflict() *ModifyRegistrationDetailsConflict {
	return &ModifyRegistrationDetailsConflict{}
}

/*
ModifyRegistrationDetailsConflict describes a response with status code 409, with default header values.

Conflict
*/
type ModifyRegistrationDetailsConflict struct {
}

// IsSuccess returns true when this modify registration details conflict response has a 2xx status code
func (o *ModifyRegistrationDetailsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify registration details conflict response has a 3xx status code
func (o *ModifyRegistrationDetailsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify registration details conflict response has a 4xx status code
func (o *ModifyRegistrationDetailsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify registration details conflict response has a 5xx status code
func (o *ModifyRegistrationDetailsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this modify registration details conflict response a status code equal to that given
func (o *ModifyRegistrationDetailsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the modify registration details conflict response
func (o *ModifyRegistrationDetailsConflict) Code() int {
	return 409
}

func (o *ModifyRegistrationDetailsConflict) Error() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsConflict ", 409)
}

func (o *ModifyRegistrationDetailsConflict) String() string {
	return fmt.Sprintf("[PATCH /applications/{appName}][%d] modifyRegistrationDetailsConflict ", 409)
}

func (o *ModifyRegistrationDetailsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
