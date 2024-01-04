// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RerunApplicationJobReader is a Reader for the RerunApplicationJob structure.
type RerunApplicationJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RerunApplicationJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRerunApplicationJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRerunApplicationJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRerunApplicationJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/jobs/{jobName}/rerun] rerunApplicationJob", response, response.Code())
	}
}

// NewRerunApplicationJobNoContent creates a RerunApplicationJobNoContent with default headers values
func NewRerunApplicationJobNoContent() *RerunApplicationJobNoContent {
	return &RerunApplicationJobNoContent{}
}

/*
RerunApplicationJobNoContent describes a response with status code 204, with default header values.

Job rerun ok
*/
type RerunApplicationJobNoContent struct {
}

// IsSuccess returns true when this rerun application job no content response has a 2xx status code
func (o *RerunApplicationJobNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rerun application job no content response has a 3xx status code
func (o *RerunApplicationJobNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rerun application job no content response has a 4xx status code
func (o *RerunApplicationJobNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this rerun application job no content response has a 5xx status code
func (o *RerunApplicationJobNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this rerun application job no content response a status code equal to that given
func (o *RerunApplicationJobNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the rerun application job no content response
func (o *RerunApplicationJobNoContent) Code() int {
	return 204
}

func (o *RerunApplicationJobNoContent) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/rerun][%d] rerunApplicationJobNoContent ", 204)
}

func (o *RerunApplicationJobNoContent) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/rerun][%d] rerunApplicationJobNoContent ", 204)
}

func (o *RerunApplicationJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRerunApplicationJobUnauthorized creates a RerunApplicationJobUnauthorized with default headers values
func NewRerunApplicationJobUnauthorized() *RerunApplicationJobUnauthorized {
	return &RerunApplicationJobUnauthorized{}
}

/*
RerunApplicationJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RerunApplicationJobUnauthorized struct {
}

// IsSuccess returns true when this rerun application job unauthorized response has a 2xx status code
func (o *RerunApplicationJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rerun application job unauthorized response has a 3xx status code
func (o *RerunApplicationJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rerun application job unauthorized response has a 4xx status code
func (o *RerunApplicationJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this rerun application job unauthorized response has a 5xx status code
func (o *RerunApplicationJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this rerun application job unauthorized response a status code equal to that given
func (o *RerunApplicationJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the rerun application job unauthorized response
func (o *RerunApplicationJobUnauthorized) Code() int {
	return 401
}

func (o *RerunApplicationJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/rerun][%d] rerunApplicationJobUnauthorized ", 401)
}

func (o *RerunApplicationJobUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/rerun][%d] rerunApplicationJobUnauthorized ", 401)
}

func (o *RerunApplicationJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRerunApplicationJobNotFound creates a RerunApplicationJobNotFound with default headers values
func NewRerunApplicationJobNotFound() *RerunApplicationJobNotFound {
	return &RerunApplicationJobNotFound{}
}

/*
RerunApplicationJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type RerunApplicationJobNotFound struct {
}

// IsSuccess returns true when this rerun application job not found response has a 2xx status code
func (o *RerunApplicationJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rerun application job not found response has a 3xx status code
func (o *RerunApplicationJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rerun application job not found response has a 4xx status code
func (o *RerunApplicationJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this rerun application job not found response has a 5xx status code
func (o *RerunApplicationJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this rerun application job not found response a status code equal to that given
func (o *RerunApplicationJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the rerun application job not found response
func (o *RerunApplicationJobNotFound) Code() int {
	return 404
}

func (o *RerunApplicationJobNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/rerun][%d] rerunApplicationJobNotFound ", 404)
}

func (o *RerunApplicationJobNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/rerun][%d] rerunApplicationJobNotFound ", 404)
}

func (o *RerunApplicationJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
