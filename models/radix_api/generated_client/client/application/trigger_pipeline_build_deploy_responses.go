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

// TriggerPipelineBuildDeployReader is a Reader for the TriggerPipelineBuildDeploy structure.
type TriggerPipelineBuildDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerPipelineBuildDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriggerPipelineBuildDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTriggerPipelineBuildDeployForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTriggerPipelineBuildDeployNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/pipelines/build-deploy] triggerPipelineBuildDeploy", response, response.Code())
	}
}

// NewTriggerPipelineBuildDeployOK creates a TriggerPipelineBuildDeployOK with default headers values
func NewTriggerPipelineBuildDeployOK() *TriggerPipelineBuildDeployOK {
	return &TriggerPipelineBuildDeployOK{}
}

/*
TriggerPipelineBuildDeployOK describes a response with status code 200, with default header values.

Successful trigger pipeline
*/
type TriggerPipelineBuildDeployOK struct {
	Payload *models.JobSummary
}

// IsSuccess returns true when this trigger pipeline build deploy o k response has a 2xx status code
func (o *TriggerPipelineBuildDeployOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger pipeline build deploy o k response has a 3xx status code
func (o *TriggerPipelineBuildDeployOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline build deploy o k response has a 4xx status code
func (o *TriggerPipelineBuildDeployOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger pipeline build deploy o k response has a 5xx status code
func (o *TriggerPipelineBuildDeployOK) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline build deploy o k response a status code equal to that given
func (o *TriggerPipelineBuildDeployOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the trigger pipeline build deploy o k response
func (o *TriggerPipelineBuildDeployOK) Code() int {
	return 200
}

func (o *TriggerPipelineBuildDeployOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build-deploy][%d] triggerPipelineBuildDeployOK %s", 200, payload)
}

func (o *TriggerPipelineBuildDeployOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build-deploy][%d] triggerPipelineBuildDeployOK %s", 200, payload)
}

func (o *TriggerPipelineBuildDeployOK) GetPayload() *models.JobSummary {
	return o.Payload
}

func (o *TriggerPipelineBuildDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggerPipelineBuildDeployForbidden creates a TriggerPipelineBuildDeployForbidden with default headers values
func NewTriggerPipelineBuildDeployForbidden() *TriggerPipelineBuildDeployForbidden {
	return &TriggerPipelineBuildDeployForbidden{}
}

/*
TriggerPipelineBuildDeployForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TriggerPipelineBuildDeployForbidden struct {
}

// IsSuccess returns true when this trigger pipeline build deploy forbidden response has a 2xx status code
func (o *TriggerPipelineBuildDeployForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline build deploy forbidden response has a 3xx status code
func (o *TriggerPipelineBuildDeployForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline build deploy forbidden response has a 4xx status code
func (o *TriggerPipelineBuildDeployForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline build deploy forbidden response has a 5xx status code
func (o *TriggerPipelineBuildDeployForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline build deploy forbidden response a status code equal to that given
func (o *TriggerPipelineBuildDeployForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the trigger pipeline build deploy forbidden response
func (o *TriggerPipelineBuildDeployForbidden) Code() int {
	return 403
}

func (o *TriggerPipelineBuildDeployForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build-deploy][%d] triggerPipelineBuildDeployForbidden", 403)
}

func (o *TriggerPipelineBuildDeployForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build-deploy][%d] triggerPipelineBuildDeployForbidden", 403)
}

func (o *TriggerPipelineBuildDeployForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerPipelineBuildDeployNotFound creates a TriggerPipelineBuildDeployNotFound with default headers values
func NewTriggerPipelineBuildDeployNotFound() *TriggerPipelineBuildDeployNotFound {
	return &TriggerPipelineBuildDeployNotFound{}
}

/*
TriggerPipelineBuildDeployNotFound describes a response with status code 404, with default header values.

Not found
*/
type TriggerPipelineBuildDeployNotFound struct {
}

// IsSuccess returns true when this trigger pipeline build deploy not found response has a 2xx status code
func (o *TriggerPipelineBuildDeployNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline build deploy not found response has a 3xx status code
func (o *TriggerPipelineBuildDeployNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline build deploy not found response has a 4xx status code
func (o *TriggerPipelineBuildDeployNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline build deploy not found response has a 5xx status code
func (o *TriggerPipelineBuildDeployNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline build deploy not found response a status code equal to that given
func (o *TriggerPipelineBuildDeployNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the trigger pipeline build deploy not found response
func (o *TriggerPipelineBuildDeployNotFound) Code() int {
	return 404
}

func (o *TriggerPipelineBuildDeployNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build-deploy][%d] triggerPipelineBuildDeployNotFound", 404)
}

func (o *TriggerPipelineBuildDeployNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/build-deploy][%d] triggerPipelineBuildDeployNotFound", 404)
}

func (o *TriggerPipelineBuildDeployNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
