// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models"
)

// GetEnvironmentEventsReader is a Reader for the GetEnvironmentEvents structure.
type GetEnvironmentEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnvironmentEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEnvironmentEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEnvironmentEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/events] getEnvironmentEvents", response, response.Code())
	}
}

// NewGetEnvironmentEventsOK creates a GetEnvironmentEventsOK with default headers values
func NewGetEnvironmentEventsOK() *GetEnvironmentEventsOK {
	return &GetEnvironmentEventsOK{}
}

/*
GetEnvironmentEventsOK describes a response with status code 200, with default header values.

Successful get environment events
*/
type GetEnvironmentEventsOK struct {
	Payload *models.Event
}

// IsSuccess returns true when this get environment events o k response has a 2xx status code
func (o *GetEnvironmentEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get environment events o k response has a 3xx status code
func (o *GetEnvironmentEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment events o k response has a 4xx status code
func (o *GetEnvironmentEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get environment events o k response has a 5xx status code
func (o *GetEnvironmentEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment events o k response a status code equal to that given
func (o *GetEnvironmentEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get environment events o k response
func (o *GetEnvironmentEventsOK) Code() int {
	return 200
}

func (o *GetEnvironmentEventsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/events][%d] getEnvironmentEventsOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmentEventsOK) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/events][%d] getEnvironmentEventsOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmentEventsOK) GetPayload() *models.Event {
	return o.Payload
}

func (o *GetEnvironmentEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvironmentEventsUnauthorized creates a GetEnvironmentEventsUnauthorized with default headers values
func NewGetEnvironmentEventsUnauthorized() *GetEnvironmentEventsUnauthorized {
	return &GetEnvironmentEventsUnauthorized{}
}

/*
GetEnvironmentEventsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetEnvironmentEventsUnauthorized struct {
}

// IsSuccess returns true when this get environment events unauthorized response has a 2xx status code
func (o *GetEnvironmentEventsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get environment events unauthorized response has a 3xx status code
func (o *GetEnvironmentEventsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment events unauthorized response has a 4xx status code
func (o *GetEnvironmentEventsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get environment events unauthorized response has a 5xx status code
func (o *GetEnvironmentEventsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment events unauthorized response a status code equal to that given
func (o *GetEnvironmentEventsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get environment events unauthorized response
func (o *GetEnvironmentEventsUnauthorized) Code() int {
	return 401
}

func (o *GetEnvironmentEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/events][%d] getEnvironmentEventsUnauthorized ", 401)
}

func (o *GetEnvironmentEventsUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/events][%d] getEnvironmentEventsUnauthorized ", 401)
}

func (o *GetEnvironmentEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmentEventsNotFound creates a GetEnvironmentEventsNotFound with default headers values
func NewGetEnvironmentEventsNotFound() *GetEnvironmentEventsNotFound {
	return &GetEnvironmentEventsNotFound{}
}

/*
GetEnvironmentEventsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetEnvironmentEventsNotFound struct {
}

// IsSuccess returns true when this get environment events not found response has a 2xx status code
func (o *GetEnvironmentEventsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get environment events not found response has a 3xx status code
func (o *GetEnvironmentEventsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment events not found response has a 4xx status code
func (o *GetEnvironmentEventsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get environment events not found response has a 5xx status code
func (o *GetEnvironmentEventsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment events not found response a status code equal to that given
func (o *GetEnvironmentEventsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get environment events not found response
func (o *GetEnvironmentEventsNotFound) Code() int {
	return 404
}

func (o *GetEnvironmentEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/events][%d] getEnvironmentEventsNotFound ", 404)
}

func (o *GetEnvironmentEventsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/events][%d] getEnvironmentEventsNotFound ", 404)
}

func (o *GetEnvironmentEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}