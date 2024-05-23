// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// DisableEnvironmentAlertingReader is a Reader for the DisableEnvironmentAlerting structure.
type DisableEnvironmentAlertingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableEnvironmentAlertingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableEnvironmentAlertingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDisableEnvironmentAlertingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDisableEnvironmentAlertingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDisableEnvironmentAlertingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDisableEnvironmentAlertingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDisableEnvironmentAlertingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/environments/{envName}/alerting/disable] disableEnvironmentAlerting", response, response.Code())
	}
}

// NewDisableEnvironmentAlertingOK creates a DisableEnvironmentAlertingOK with default headers values
func NewDisableEnvironmentAlertingOK() *DisableEnvironmentAlertingOK {
	return &DisableEnvironmentAlertingOK{}
}

/*
DisableEnvironmentAlertingOK describes a response with status code 200, with default header values.

Successful disable alerting
*/
type DisableEnvironmentAlertingOK struct {
	Payload *models.AlertingConfig
}

// IsSuccess returns true when this disable environment alerting o k response has a 2xx status code
func (o *DisableEnvironmentAlertingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable environment alerting o k response has a 3xx status code
func (o *DisableEnvironmentAlertingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable environment alerting o k response has a 4xx status code
func (o *DisableEnvironmentAlertingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable environment alerting o k response has a 5xx status code
func (o *DisableEnvironmentAlertingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disable environment alerting o k response a status code equal to that given
func (o *DisableEnvironmentAlertingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the disable environment alerting o k response
func (o *DisableEnvironmentAlertingOK) Code() int {
	return 200
}

func (o *DisableEnvironmentAlertingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingOK %s", 200, payload)
}

func (o *DisableEnvironmentAlertingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingOK %s", 200, payload)
}

func (o *DisableEnvironmentAlertingOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *DisableEnvironmentAlertingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableEnvironmentAlertingBadRequest creates a DisableEnvironmentAlertingBadRequest with default headers values
func NewDisableEnvironmentAlertingBadRequest() *DisableEnvironmentAlertingBadRequest {
	return &DisableEnvironmentAlertingBadRequest{}
}

/*
DisableEnvironmentAlertingBadRequest describes a response with status code 400, with default header values.

Alerting already enabled
*/
type DisableEnvironmentAlertingBadRequest struct {
}

// IsSuccess returns true when this disable environment alerting bad request response has a 2xx status code
func (o *DisableEnvironmentAlertingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable environment alerting bad request response has a 3xx status code
func (o *DisableEnvironmentAlertingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable environment alerting bad request response has a 4xx status code
func (o *DisableEnvironmentAlertingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable environment alerting bad request response has a 5xx status code
func (o *DisableEnvironmentAlertingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this disable environment alerting bad request response a status code equal to that given
func (o *DisableEnvironmentAlertingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the disable environment alerting bad request response
func (o *DisableEnvironmentAlertingBadRequest) Code() int {
	return 400
}

func (o *DisableEnvironmentAlertingBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingBadRequest", 400)
}

func (o *DisableEnvironmentAlertingBadRequest) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingBadRequest", 400)
}

func (o *DisableEnvironmentAlertingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableEnvironmentAlertingUnauthorized creates a DisableEnvironmentAlertingUnauthorized with default headers values
func NewDisableEnvironmentAlertingUnauthorized() *DisableEnvironmentAlertingUnauthorized {
	return &DisableEnvironmentAlertingUnauthorized{}
}

/*
DisableEnvironmentAlertingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DisableEnvironmentAlertingUnauthorized struct {
}

// IsSuccess returns true when this disable environment alerting unauthorized response has a 2xx status code
func (o *DisableEnvironmentAlertingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable environment alerting unauthorized response has a 3xx status code
func (o *DisableEnvironmentAlertingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable environment alerting unauthorized response has a 4xx status code
func (o *DisableEnvironmentAlertingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable environment alerting unauthorized response has a 5xx status code
func (o *DisableEnvironmentAlertingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this disable environment alerting unauthorized response a status code equal to that given
func (o *DisableEnvironmentAlertingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the disable environment alerting unauthorized response
func (o *DisableEnvironmentAlertingUnauthorized) Code() int {
	return 401
}

func (o *DisableEnvironmentAlertingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingUnauthorized", 401)
}

func (o *DisableEnvironmentAlertingUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingUnauthorized", 401)
}

func (o *DisableEnvironmentAlertingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableEnvironmentAlertingForbidden creates a DisableEnvironmentAlertingForbidden with default headers values
func NewDisableEnvironmentAlertingForbidden() *DisableEnvironmentAlertingForbidden {
	return &DisableEnvironmentAlertingForbidden{}
}

/*
DisableEnvironmentAlertingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DisableEnvironmentAlertingForbidden struct {
}

// IsSuccess returns true when this disable environment alerting forbidden response has a 2xx status code
func (o *DisableEnvironmentAlertingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable environment alerting forbidden response has a 3xx status code
func (o *DisableEnvironmentAlertingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable environment alerting forbidden response has a 4xx status code
func (o *DisableEnvironmentAlertingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable environment alerting forbidden response has a 5xx status code
func (o *DisableEnvironmentAlertingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this disable environment alerting forbidden response a status code equal to that given
func (o *DisableEnvironmentAlertingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the disable environment alerting forbidden response
func (o *DisableEnvironmentAlertingForbidden) Code() int {
	return 403
}

func (o *DisableEnvironmentAlertingForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingForbidden", 403)
}

func (o *DisableEnvironmentAlertingForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingForbidden", 403)
}

func (o *DisableEnvironmentAlertingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableEnvironmentAlertingNotFound creates a DisableEnvironmentAlertingNotFound with default headers values
func NewDisableEnvironmentAlertingNotFound() *DisableEnvironmentAlertingNotFound {
	return &DisableEnvironmentAlertingNotFound{}
}

/*
DisableEnvironmentAlertingNotFound describes a response with status code 404, with default header values.

Not found
*/
type DisableEnvironmentAlertingNotFound struct {
}

// IsSuccess returns true when this disable environment alerting not found response has a 2xx status code
func (o *DisableEnvironmentAlertingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable environment alerting not found response has a 3xx status code
func (o *DisableEnvironmentAlertingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable environment alerting not found response has a 4xx status code
func (o *DisableEnvironmentAlertingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this disable environment alerting not found response has a 5xx status code
func (o *DisableEnvironmentAlertingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this disable environment alerting not found response a status code equal to that given
func (o *DisableEnvironmentAlertingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the disable environment alerting not found response
func (o *DisableEnvironmentAlertingNotFound) Code() int {
	return 404
}

func (o *DisableEnvironmentAlertingNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingNotFound", 404)
}

func (o *DisableEnvironmentAlertingNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingNotFound", 404)
}

func (o *DisableEnvironmentAlertingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableEnvironmentAlertingInternalServerError creates a DisableEnvironmentAlertingInternalServerError with default headers values
func NewDisableEnvironmentAlertingInternalServerError() *DisableEnvironmentAlertingInternalServerError {
	return &DisableEnvironmentAlertingInternalServerError{}
}

/*
DisableEnvironmentAlertingInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DisableEnvironmentAlertingInternalServerError struct {
}

// IsSuccess returns true when this disable environment alerting internal server error response has a 2xx status code
func (o *DisableEnvironmentAlertingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this disable environment alerting internal server error response has a 3xx status code
func (o *DisableEnvironmentAlertingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable environment alerting internal server error response has a 4xx status code
func (o *DisableEnvironmentAlertingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable environment alerting internal server error response has a 5xx status code
func (o *DisableEnvironmentAlertingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this disable environment alerting internal server error response a status code equal to that given
func (o *DisableEnvironmentAlertingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the disable environment alerting internal server error response
func (o *DisableEnvironmentAlertingInternalServerError) Code() int {
	return 500
}

func (o *DisableEnvironmentAlertingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingInternalServerError", 500)
}

func (o *DisableEnvironmentAlertingInternalServerError) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingInternalServerError", 500)
}

func (o *DisableEnvironmentAlertingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
