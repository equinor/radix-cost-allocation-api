// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTektonPipelineRunTaskStepLogsReader is a Reader for the GetTektonPipelineRunTaskStepLogs structure.
type GetTektonPipelineRunTaskStepLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTektonPipelineRunTaskStepLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTektonPipelineRunTaskStepLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTektonPipelineRunTaskStepLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTektonPipelineRunTaskStepLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/logs/{stepName}] getTektonPipelineRunTaskStepLogs", response, response.Code())
	}
}

// NewGetTektonPipelineRunTaskStepLogsOK creates a GetTektonPipelineRunTaskStepLogsOK with default headers values
func NewGetTektonPipelineRunTaskStepLogsOK() *GetTektonPipelineRunTaskStepLogsOK {
	return &GetTektonPipelineRunTaskStepLogsOK{}
}

/*
GetTektonPipelineRunTaskStepLogsOK describes a response with status code 200, with default header values.

Task step log
*/
type GetTektonPipelineRunTaskStepLogsOK struct {
	Payload string
}

// IsSuccess returns true when this get tekton pipeline run task step logs o k response has a 2xx status code
func (o *GetTektonPipelineRunTaskStepLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tekton pipeline run task step logs o k response has a 3xx status code
func (o *GetTektonPipelineRunTaskStepLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run task step logs o k response has a 4xx status code
func (o *GetTektonPipelineRunTaskStepLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tekton pipeline run task step logs o k response has a 5xx status code
func (o *GetTektonPipelineRunTaskStepLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run task step logs o k response a status code equal to that given
func (o *GetTektonPipelineRunTaskStepLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tekton pipeline run task step logs o k response
func (o *GetTektonPipelineRunTaskStepLogsOK) Code() int {
	return 200
}

func (o *GetTektonPipelineRunTaskStepLogsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/logs/{stepName}][%d] getTektonPipelineRunTaskStepLogsOK %s", 200, payload)
}

func (o *GetTektonPipelineRunTaskStepLogsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/logs/{stepName}][%d] getTektonPipelineRunTaskStepLogsOK %s", 200, payload)
}

func (o *GetTektonPipelineRunTaskStepLogsOK) GetPayload() string {
	return o.Payload
}

func (o *GetTektonPipelineRunTaskStepLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTektonPipelineRunTaskStepLogsUnauthorized creates a GetTektonPipelineRunTaskStepLogsUnauthorized with default headers values
func NewGetTektonPipelineRunTaskStepLogsUnauthorized() *GetTektonPipelineRunTaskStepLogsUnauthorized {
	return &GetTektonPipelineRunTaskStepLogsUnauthorized{}
}

/*
GetTektonPipelineRunTaskStepLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTektonPipelineRunTaskStepLogsUnauthorized struct {
}

// IsSuccess returns true when this get tekton pipeline run task step logs unauthorized response has a 2xx status code
func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline run task step logs unauthorized response has a 3xx status code
func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run task step logs unauthorized response has a 4xx status code
func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tekton pipeline run task step logs unauthorized response has a 5xx status code
func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run task step logs unauthorized response a status code equal to that given
func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get tekton pipeline run task step logs unauthorized response
func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) Code() int {
	return 401
}

func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/logs/{stepName}][%d] getTektonPipelineRunTaskStepLogsUnauthorized", 401)
}

func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/logs/{stepName}][%d] getTektonPipelineRunTaskStepLogsUnauthorized", 401)
}

func (o *GetTektonPipelineRunTaskStepLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTektonPipelineRunTaskStepLogsNotFound creates a GetTektonPipelineRunTaskStepLogsNotFound with default headers values
func NewGetTektonPipelineRunTaskStepLogsNotFound() *GetTektonPipelineRunTaskStepLogsNotFound {
	return &GetTektonPipelineRunTaskStepLogsNotFound{}
}

/*
GetTektonPipelineRunTaskStepLogsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetTektonPipelineRunTaskStepLogsNotFound struct {
}

// IsSuccess returns true when this get tekton pipeline run task step logs not found response has a 2xx status code
func (o *GetTektonPipelineRunTaskStepLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline run task step logs not found response has a 3xx status code
func (o *GetTektonPipelineRunTaskStepLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run task step logs not found response has a 4xx status code
func (o *GetTektonPipelineRunTaskStepLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tekton pipeline run task step logs not found response has a 5xx status code
func (o *GetTektonPipelineRunTaskStepLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run task step logs not found response a status code equal to that given
func (o *GetTektonPipelineRunTaskStepLogsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get tekton pipeline run task step logs not found response
func (o *GetTektonPipelineRunTaskStepLogsNotFound) Code() int {
	return 404
}

func (o *GetTektonPipelineRunTaskStepLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/logs/{stepName}][%d] getTektonPipelineRunTaskStepLogsNotFound", 404)
}

func (o *GetTektonPipelineRunTaskStepLogsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}/logs/{stepName}][%d] getTektonPipelineRunTaskStepLogsNotFound", 404)
}

func (o *GetTektonPipelineRunTaskStepLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
