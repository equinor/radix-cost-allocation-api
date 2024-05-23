// Code generated by go-swagger; DO NOT EDIT.

package application

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

// UpdateApplicationAlertingConfigReader is a Reader for the UpdateApplicationAlertingConfig structure.
type UpdateApplicationAlertingConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationAlertingConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateApplicationAlertingConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationAlertingConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateApplicationAlertingConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateApplicationAlertingConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateApplicationAlertingConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateApplicationAlertingConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /applications/{appName}/alerting] updateApplicationAlertingConfig", response, response.Code())
	}
}

// NewUpdateApplicationAlertingConfigOK creates a UpdateApplicationAlertingConfigOK with default headers values
func NewUpdateApplicationAlertingConfigOK() *UpdateApplicationAlertingConfigOK {
	return &UpdateApplicationAlertingConfigOK{}
}

/*
UpdateApplicationAlertingConfigOK describes a response with status code 200, with default header values.

Successful alerts config update
*/
type UpdateApplicationAlertingConfigOK struct {
	Payload *models.AlertingConfig
}

// IsSuccess returns true when this update application alerting config o k response has a 2xx status code
func (o *UpdateApplicationAlertingConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update application alerting config o k response has a 3xx status code
func (o *UpdateApplicationAlertingConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application alerting config o k response has a 4xx status code
func (o *UpdateApplicationAlertingConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update application alerting config o k response has a 5xx status code
func (o *UpdateApplicationAlertingConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update application alerting config o k response a status code equal to that given
func (o *UpdateApplicationAlertingConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update application alerting config o k response
func (o *UpdateApplicationAlertingConfigOK) Code() int {
	return 200
}

func (o *UpdateApplicationAlertingConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigOK %s", 200, payload)
}

func (o *UpdateApplicationAlertingConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigOK %s", 200, payload)
}

func (o *UpdateApplicationAlertingConfigOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *UpdateApplicationAlertingConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationAlertingConfigBadRequest creates a UpdateApplicationAlertingConfigBadRequest with default headers values
func NewUpdateApplicationAlertingConfigBadRequest() *UpdateApplicationAlertingConfigBadRequest {
	return &UpdateApplicationAlertingConfigBadRequest{}
}

/*
UpdateApplicationAlertingConfigBadRequest describes a response with status code 400, with default header values.

Invalid configuration
*/
type UpdateApplicationAlertingConfigBadRequest struct {
}

// IsSuccess returns true when this update application alerting config bad request response has a 2xx status code
func (o *UpdateApplicationAlertingConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update application alerting config bad request response has a 3xx status code
func (o *UpdateApplicationAlertingConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application alerting config bad request response has a 4xx status code
func (o *UpdateApplicationAlertingConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update application alerting config bad request response has a 5xx status code
func (o *UpdateApplicationAlertingConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update application alerting config bad request response a status code equal to that given
func (o *UpdateApplicationAlertingConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update application alerting config bad request response
func (o *UpdateApplicationAlertingConfigBadRequest) Code() int {
	return 400
}

func (o *UpdateApplicationAlertingConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigBadRequest", 400)
}

func (o *UpdateApplicationAlertingConfigBadRequest) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigBadRequest", 400)
}

func (o *UpdateApplicationAlertingConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationAlertingConfigUnauthorized creates a UpdateApplicationAlertingConfigUnauthorized with default headers values
func NewUpdateApplicationAlertingConfigUnauthorized() *UpdateApplicationAlertingConfigUnauthorized {
	return &UpdateApplicationAlertingConfigUnauthorized{}
}

/*
UpdateApplicationAlertingConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateApplicationAlertingConfigUnauthorized struct {
}

// IsSuccess returns true when this update application alerting config unauthorized response has a 2xx status code
func (o *UpdateApplicationAlertingConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update application alerting config unauthorized response has a 3xx status code
func (o *UpdateApplicationAlertingConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application alerting config unauthorized response has a 4xx status code
func (o *UpdateApplicationAlertingConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update application alerting config unauthorized response has a 5xx status code
func (o *UpdateApplicationAlertingConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update application alerting config unauthorized response a status code equal to that given
func (o *UpdateApplicationAlertingConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update application alerting config unauthorized response
func (o *UpdateApplicationAlertingConfigUnauthorized) Code() int {
	return 401
}

func (o *UpdateApplicationAlertingConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigUnauthorized", 401)
}

func (o *UpdateApplicationAlertingConfigUnauthorized) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigUnauthorized", 401)
}

func (o *UpdateApplicationAlertingConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationAlertingConfigForbidden creates a UpdateApplicationAlertingConfigForbidden with default headers values
func NewUpdateApplicationAlertingConfigForbidden() *UpdateApplicationAlertingConfigForbidden {
	return &UpdateApplicationAlertingConfigForbidden{}
}

/*
UpdateApplicationAlertingConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateApplicationAlertingConfigForbidden struct {
}

// IsSuccess returns true when this update application alerting config forbidden response has a 2xx status code
func (o *UpdateApplicationAlertingConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update application alerting config forbidden response has a 3xx status code
func (o *UpdateApplicationAlertingConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application alerting config forbidden response has a 4xx status code
func (o *UpdateApplicationAlertingConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update application alerting config forbidden response has a 5xx status code
func (o *UpdateApplicationAlertingConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update application alerting config forbidden response a status code equal to that given
func (o *UpdateApplicationAlertingConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update application alerting config forbidden response
func (o *UpdateApplicationAlertingConfigForbidden) Code() int {
	return 403
}

func (o *UpdateApplicationAlertingConfigForbidden) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigForbidden", 403)
}

func (o *UpdateApplicationAlertingConfigForbidden) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigForbidden", 403)
}

func (o *UpdateApplicationAlertingConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationAlertingConfigNotFound creates a UpdateApplicationAlertingConfigNotFound with default headers values
func NewUpdateApplicationAlertingConfigNotFound() *UpdateApplicationAlertingConfigNotFound {
	return &UpdateApplicationAlertingConfigNotFound{}
}

/*
UpdateApplicationAlertingConfigNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateApplicationAlertingConfigNotFound struct {
}

// IsSuccess returns true when this update application alerting config not found response has a 2xx status code
func (o *UpdateApplicationAlertingConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update application alerting config not found response has a 3xx status code
func (o *UpdateApplicationAlertingConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application alerting config not found response has a 4xx status code
func (o *UpdateApplicationAlertingConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update application alerting config not found response has a 5xx status code
func (o *UpdateApplicationAlertingConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update application alerting config not found response a status code equal to that given
func (o *UpdateApplicationAlertingConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update application alerting config not found response
func (o *UpdateApplicationAlertingConfigNotFound) Code() int {
	return 404
}

func (o *UpdateApplicationAlertingConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigNotFound", 404)
}

func (o *UpdateApplicationAlertingConfigNotFound) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigNotFound", 404)
}

func (o *UpdateApplicationAlertingConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationAlertingConfigInternalServerError creates a UpdateApplicationAlertingConfigInternalServerError with default headers values
func NewUpdateApplicationAlertingConfigInternalServerError() *UpdateApplicationAlertingConfigInternalServerError {
	return &UpdateApplicationAlertingConfigInternalServerError{}
}

/*
UpdateApplicationAlertingConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateApplicationAlertingConfigInternalServerError struct {
}

// IsSuccess returns true when this update application alerting config internal server error response has a 2xx status code
func (o *UpdateApplicationAlertingConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update application alerting config internal server error response has a 3xx status code
func (o *UpdateApplicationAlertingConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update application alerting config internal server error response has a 4xx status code
func (o *UpdateApplicationAlertingConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update application alerting config internal server error response has a 5xx status code
func (o *UpdateApplicationAlertingConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update application alerting config internal server error response a status code equal to that given
func (o *UpdateApplicationAlertingConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update application alerting config internal server error response
func (o *UpdateApplicationAlertingConfigInternalServerError) Code() int {
	return 500
}

func (o *UpdateApplicationAlertingConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigInternalServerError", 500)
}

func (o *UpdateApplicationAlertingConfigInternalServerError) String() string {
	return fmt.Sprintf("[PUT /applications/{appName}/alerting][%d] updateApplicationAlertingConfigInternalServerError", 500)
}

func (o *UpdateApplicationAlertingConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
