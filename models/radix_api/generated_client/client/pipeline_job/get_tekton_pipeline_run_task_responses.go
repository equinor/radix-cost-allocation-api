// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models"
)

// GetTektonPipelineRunTaskReader is a Reader for the GetTektonPipelineRunTask structure.
type GetTektonPipelineRunTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTektonPipelineRunTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTektonPipelineRunTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTektonPipelineRunTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTektonPipelineRunTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}] getTektonPipelineRunTask", response, response.Code())
	}
}

// NewGetTektonPipelineRunTaskOK creates a GetTektonPipelineRunTaskOK with default headers values
func NewGetTektonPipelineRunTaskOK() *GetTektonPipelineRunTaskOK {
	return &GetTektonPipelineRunTaskOK{}
}

/*
GetTektonPipelineRunTaskOK describes a response with status code 200, with default header values.

Pipeline Run Task
*/
type GetTektonPipelineRunTaskOK struct {
	Payload *models.PipelineRunTask
}

// IsSuccess returns true when this get tekton pipeline run task o k response has a 2xx status code
func (o *GetTektonPipelineRunTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tekton pipeline run task o k response has a 3xx status code
func (o *GetTektonPipelineRunTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run task o k response has a 4xx status code
func (o *GetTektonPipelineRunTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tekton pipeline run task o k response has a 5xx status code
func (o *GetTektonPipelineRunTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run task o k response a status code equal to that given
func (o *GetTektonPipelineRunTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tekton pipeline run task o k response
func (o *GetTektonPipelineRunTaskOK) Code() int {
	return 200
}

func (o *GetTektonPipelineRunTaskOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}][%d] getTektonPipelineRunTaskOK  %+v", 200, o.Payload)
}

func (o *GetTektonPipelineRunTaskOK) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}][%d] getTektonPipelineRunTaskOK  %+v", 200, o.Payload)
}

func (o *GetTektonPipelineRunTaskOK) GetPayload() *models.PipelineRunTask {
	return o.Payload
}

func (o *GetTektonPipelineRunTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineRunTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTektonPipelineRunTaskUnauthorized creates a GetTektonPipelineRunTaskUnauthorized with default headers values
func NewGetTektonPipelineRunTaskUnauthorized() *GetTektonPipelineRunTaskUnauthorized {
	return &GetTektonPipelineRunTaskUnauthorized{}
}

/*
GetTektonPipelineRunTaskUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTektonPipelineRunTaskUnauthorized struct {
}

// IsSuccess returns true when this get tekton pipeline run task unauthorized response has a 2xx status code
func (o *GetTektonPipelineRunTaskUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline run task unauthorized response has a 3xx status code
func (o *GetTektonPipelineRunTaskUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run task unauthorized response has a 4xx status code
func (o *GetTektonPipelineRunTaskUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tekton pipeline run task unauthorized response has a 5xx status code
func (o *GetTektonPipelineRunTaskUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run task unauthorized response a status code equal to that given
func (o *GetTektonPipelineRunTaskUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get tekton pipeline run task unauthorized response
func (o *GetTektonPipelineRunTaskUnauthorized) Code() int {
	return 401
}

func (o *GetTektonPipelineRunTaskUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}][%d] getTektonPipelineRunTaskUnauthorized ", 401)
}

func (o *GetTektonPipelineRunTaskUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}][%d] getTektonPipelineRunTaskUnauthorized ", 401)
}

func (o *GetTektonPipelineRunTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTektonPipelineRunTaskNotFound creates a GetTektonPipelineRunTaskNotFound with default headers values
func NewGetTektonPipelineRunTaskNotFound() *GetTektonPipelineRunTaskNotFound {
	return &GetTektonPipelineRunTaskNotFound{}
}

/*
GetTektonPipelineRunTaskNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetTektonPipelineRunTaskNotFound struct {
}

// IsSuccess returns true when this get tekton pipeline run task not found response has a 2xx status code
func (o *GetTektonPipelineRunTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tekton pipeline run task not found response has a 3xx status code
func (o *GetTektonPipelineRunTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tekton pipeline run task not found response has a 4xx status code
func (o *GetTektonPipelineRunTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tekton pipeline run task not found response has a 5xx status code
func (o *GetTektonPipelineRunTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tekton pipeline run task not found response a status code equal to that given
func (o *GetTektonPipelineRunTaskNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get tekton pipeline run task not found response
func (o *GetTektonPipelineRunTaskNotFound) Code() int {
	return 404
}

func (o *GetTektonPipelineRunTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}][%d] getTektonPipelineRunTaskNotFound ", 404)
}

func (o *GetTektonPipelineRunTaskNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/pipelineruns/{pipelineRunName}/tasks/{taskName}][%d] getTektonPipelineRunTaskNotFound ", 404)
}

func (o *GetTektonPipelineRunTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
