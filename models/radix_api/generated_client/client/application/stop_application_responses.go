// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopApplicationReader is a Reader for the StopApplication structure.
type StopApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/stop] stopApplication", response, response.Code())
	}
}

// NewStopApplicationOK creates a StopApplicationOK with default headers values
func NewStopApplicationOK() *StopApplicationOK {
	return &StopApplicationOK{}
}

/*
StopApplicationOK describes a response with status code 200, with default header values.

Application stopped ok
*/
type StopApplicationOK struct {
}

// IsSuccess returns true when this stop application o k response has a 2xx status code
func (o *StopApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop application o k response has a 3xx status code
func (o *StopApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop application o k response has a 4xx status code
func (o *StopApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop application o k response has a 5xx status code
func (o *StopApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop application o k response a status code equal to that given
func (o *StopApplicationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop application o k response
func (o *StopApplicationOK) Code() int {
	return 200
}

func (o *StopApplicationOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/stop][%d] stopApplicationOK ", 200)
}

func (o *StopApplicationOK) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/stop][%d] stopApplicationOK ", 200)
}

func (o *StopApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopApplicationUnauthorized creates a StopApplicationUnauthorized with default headers values
func NewStopApplicationUnauthorized() *StopApplicationUnauthorized {
	return &StopApplicationUnauthorized{}
}

/*
StopApplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopApplicationUnauthorized struct {
}

// IsSuccess returns true when this stop application unauthorized response has a 2xx status code
func (o *StopApplicationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop application unauthorized response has a 3xx status code
func (o *StopApplicationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop application unauthorized response has a 4xx status code
func (o *StopApplicationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop application unauthorized response has a 5xx status code
func (o *StopApplicationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stop application unauthorized response a status code equal to that given
func (o *StopApplicationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stop application unauthorized response
func (o *StopApplicationUnauthorized) Code() int {
	return 401
}

func (o *StopApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/stop][%d] stopApplicationUnauthorized ", 401)
}

func (o *StopApplicationUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/stop][%d] stopApplicationUnauthorized ", 401)
}

func (o *StopApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopApplicationNotFound creates a StopApplicationNotFound with default headers values
func NewStopApplicationNotFound() *StopApplicationNotFound {
	return &StopApplicationNotFound{}
}

/*
StopApplicationNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopApplicationNotFound struct {
}

// IsSuccess returns true when this stop application not found response has a 2xx status code
func (o *StopApplicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop application not found response has a 3xx status code
func (o *StopApplicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop application not found response has a 4xx status code
func (o *StopApplicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop application not found response has a 5xx status code
func (o *StopApplicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop application not found response a status code equal to that given
func (o *StopApplicationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop application not found response
func (o *StopApplicationNotFound) Code() int {
	return 404
}

func (o *StopApplicationNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/stop][%d] stopApplicationNotFound ", 404)
}

func (o *StopApplicationNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/stop][%d] stopApplicationNotFound ", 404)
}

func (o *StopApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
