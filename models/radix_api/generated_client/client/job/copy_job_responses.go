// Code generated by go-swagger; DO NOT EDIT.

package job

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

// CopyJobReader is a Reader for the CopyJob structure.
type CopyJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCopyJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCopyJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCopyJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCopyJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCopyJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy] copyJob", response, response.Code())
	}
}

// NewCopyJobOK creates a CopyJobOK with default headers values
func NewCopyJobOK() *CopyJobOK {
	return &CopyJobOK{}
}

/*
CopyJobOK describes a response with status code 200, with default header values.

Success
*/
type CopyJobOK struct {
	Payload *models.ScheduledJobSummary
}

// IsSuccess returns true when this copy job o k response has a 2xx status code
func (o *CopyJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this copy job o k response has a 3xx status code
func (o *CopyJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy job o k response has a 4xx status code
func (o *CopyJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this copy job o k response has a 5xx status code
func (o *CopyJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this copy job o k response a status code equal to that given
func (o *CopyJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the copy job o k response
func (o *CopyJobOK) Code() int {
	return 200
}

func (o *CopyJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobOK %s", 200, payload)
}

func (o *CopyJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobOK %s", 200, payload)
}

func (o *CopyJobOK) GetPayload() *models.ScheduledJobSummary {
	return o.Payload
}

func (o *CopyJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduledJobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyJobBadRequest creates a CopyJobBadRequest with default headers values
func NewCopyJobBadRequest() *CopyJobBadRequest {
	return &CopyJobBadRequest{}
}

/*
CopyJobBadRequest describes a response with status code 400, with default header values.

Invalid batch
*/
type CopyJobBadRequest struct {
}

// IsSuccess returns true when this copy job bad request response has a 2xx status code
func (o *CopyJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy job bad request response has a 3xx status code
func (o *CopyJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy job bad request response has a 4xx status code
func (o *CopyJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy job bad request response has a 5xx status code
func (o *CopyJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this copy job bad request response a status code equal to that given
func (o *CopyJobBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the copy job bad request response
func (o *CopyJobBadRequest) Code() int {
	return 400
}

func (o *CopyJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobBadRequest", 400)
}

func (o *CopyJobBadRequest) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobBadRequest", 400)
}

func (o *CopyJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyJobUnauthorized creates a CopyJobUnauthorized with default headers values
func NewCopyJobUnauthorized() *CopyJobUnauthorized {
	return &CopyJobUnauthorized{}
}

/*
CopyJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CopyJobUnauthorized struct {
}

// IsSuccess returns true when this copy job unauthorized response has a 2xx status code
func (o *CopyJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy job unauthorized response has a 3xx status code
func (o *CopyJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy job unauthorized response has a 4xx status code
func (o *CopyJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy job unauthorized response has a 5xx status code
func (o *CopyJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this copy job unauthorized response a status code equal to that given
func (o *CopyJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the copy job unauthorized response
func (o *CopyJobUnauthorized) Code() int {
	return 401
}

func (o *CopyJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobUnauthorized", 401)
}

func (o *CopyJobUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobUnauthorized", 401)
}

func (o *CopyJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyJobForbidden creates a CopyJobForbidden with default headers values
func NewCopyJobForbidden() *CopyJobForbidden {
	return &CopyJobForbidden{}
}

/*
CopyJobForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CopyJobForbidden struct {
}

// IsSuccess returns true when this copy job forbidden response has a 2xx status code
func (o *CopyJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy job forbidden response has a 3xx status code
func (o *CopyJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy job forbidden response has a 4xx status code
func (o *CopyJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy job forbidden response has a 5xx status code
func (o *CopyJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this copy job forbidden response a status code equal to that given
func (o *CopyJobForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the copy job forbidden response
func (o *CopyJobForbidden) Code() int {
	return 403
}

func (o *CopyJobForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobForbidden", 403)
}

func (o *CopyJobForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobForbidden", 403)
}

func (o *CopyJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyJobNotFound creates a CopyJobNotFound with default headers values
func NewCopyJobNotFound() *CopyJobNotFound {
	return &CopyJobNotFound{}
}

/*
CopyJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type CopyJobNotFound struct {
}

// IsSuccess returns true when this copy job not found response has a 2xx status code
func (o *CopyJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy job not found response has a 3xx status code
func (o *CopyJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy job not found response has a 4xx status code
func (o *CopyJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy job not found response has a 5xx status code
func (o *CopyJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this copy job not found response a status code equal to that given
func (o *CopyJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the copy job not found response
func (o *CopyJobNotFound) Code() int {
	return 404
}

func (o *CopyJobNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobNotFound", 404)
}

func (o *CopyJobNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy][%d] copyJobNotFound", 404)
}

func (o *CopyJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
