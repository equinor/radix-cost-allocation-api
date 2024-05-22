// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IsDeployKeyValidReader is a Reader for the IsDeployKeyValid structure.
type IsDeployKeyValidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsDeployKeyValidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsDeployKeyValidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIsDeployKeyValidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIsDeployKeyValidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIsDeployKeyValidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIsDeployKeyValidConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIsDeployKeyValidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/deploykey-valid] isDeployKeyValid", response, response.Code())
	}
}

// NewIsDeployKeyValidOK creates a IsDeployKeyValidOK with default headers values
func NewIsDeployKeyValidOK() *IsDeployKeyValidOK {
	return &IsDeployKeyValidOK{}
}

/*
IsDeployKeyValidOK describes a response with status code 200, with default header values.

Deploy key is valid
*/
type IsDeployKeyValidOK struct {
}

// IsSuccess returns true when this is deploy key valid o k response has a 2xx status code
func (o *IsDeployKeyValidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this is deploy key valid o k response has a 3xx status code
func (o *IsDeployKeyValidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is deploy key valid o k response has a 4xx status code
func (o *IsDeployKeyValidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this is deploy key valid o k response has a 5xx status code
func (o *IsDeployKeyValidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this is deploy key valid o k response a status code equal to that given
func (o *IsDeployKeyValidOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the is deploy key valid o k response
func (o *IsDeployKeyValidOK) Code() int {
	return 200
}

func (o *IsDeployKeyValidOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidOK", 200)
}

func (o *IsDeployKeyValidOK) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidOK", 200)
}

func (o *IsDeployKeyValidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsDeployKeyValidUnauthorized creates a IsDeployKeyValidUnauthorized with default headers values
func NewIsDeployKeyValidUnauthorized() *IsDeployKeyValidUnauthorized {
	return &IsDeployKeyValidUnauthorized{}
}

/*
IsDeployKeyValidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type IsDeployKeyValidUnauthorized struct {
}

// IsSuccess returns true when this is deploy key valid unauthorized response has a 2xx status code
func (o *IsDeployKeyValidUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is deploy key valid unauthorized response has a 3xx status code
func (o *IsDeployKeyValidUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is deploy key valid unauthorized response has a 4xx status code
func (o *IsDeployKeyValidUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this is deploy key valid unauthorized response has a 5xx status code
func (o *IsDeployKeyValidUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this is deploy key valid unauthorized response a status code equal to that given
func (o *IsDeployKeyValidUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the is deploy key valid unauthorized response
func (o *IsDeployKeyValidUnauthorized) Code() int {
	return 401
}

func (o *IsDeployKeyValidUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidUnauthorized", 401)
}

func (o *IsDeployKeyValidUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidUnauthorized", 401)
}

func (o *IsDeployKeyValidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsDeployKeyValidForbidden creates a IsDeployKeyValidForbidden with default headers values
func NewIsDeployKeyValidForbidden() *IsDeployKeyValidForbidden {
	return &IsDeployKeyValidForbidden{}
}

/*
IsDeployKeyValidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type IsDeployKeyValidForbidden struct {
}

// IsSuccess returns true when this is deploy key valid forbidden response has a 2xx status code
func (o *IsDeployKeyValidForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is deploy key valid forbidden response has a 3xx status code
func (o *IsDeployKeyValidForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is deploy key valid forbidden response has a 4xx status code
func (o *IsDeployKeyValidForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this is deploy key valid forbidden response has a 5xx status code
func (o *IsDeployKeyValidForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this is deploy key valid forbidden response a status code equal to that given
func (o *IsDeployKeyValidForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the is deploy key valid forbidden response
func (o *IsDeployKeyValidForbidden) Code() int {
	return 403
}

func (o *IsDeployKeyValidForbidden) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidForbidden", 403)
}

func (o *IsDeployKeyValidForbidden) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidForbidden", 403)
}

func (o *IsDeployKeyValidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsDeployKeyValidNotFound creates a IsDeployKeyValidNotFound with default headers values
func NewIsDeployKeyValidNotFound() *IsDeployKeyValidNotFound {
	return &IsDeployKeyValidNotFound{}
}

/*
IsDeployKeyValidNotFound describes a response with status code 404, with default header values.

Not found
*/
type IsDeployKeyValidNotFound struct {
}

// IsSuccess returns true when this is deploy key valid not found response has a 2xx status code
func (o *IsDeployKeyValidNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is deploy key valid not found response has a 3xx status code
func (o *IsDeployKeyValidNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is deploy key valid not found response has a 4xx status code
func (o *IsDeployKeyValidNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this is deploy key valid not found response has a 5xx status code
func (o *IsDeployKeyValidNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this is deploy key valid not found response a status code equal to that given
func (o *IsDeployKeyValidNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the is deploy key valid not found response
func (o *IsDeployKeyValidNotFound) Code() int {
	return 404
}

func (o *IsDeployKeyValidNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidNotFound", 404)
}

func (o *IsDeployKeyValidNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidNotFound", 404)
}

func (o *IsDeployKeyValidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsDeployKeyValidConflict creates a IsDeployKeyValidConflict with default headers values
func NewIsDeployKeyValidConflict() *IsDeployKeyValidConflict {
	return &IsDeployKeyValidConflict{}
}

/*
IsDeployKeyValidConflict describes a response with status code 409, with default header values.

Conflict
*/
type IsDeployKeyValidConflict struct {
}

// IsSuccess returns true when this is deploy key valid conflict response has a 2xx status code
func (o *IsDeployKeyValidConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is deploy key valid conflict response has a 3xx status code
func (o *IsDeployKeyValidConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is deploy key valid conflict response has a 4xx status code
func (o *IsDeployKeyValidConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this is deploy key valid conflict response has a 5xx status code
func (o *IsDeployKeyValidConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this is deploy key valid conflict response a status code equal to that given
func (o *IsDeployKeyValidConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the is deploy key valid conflict response
func (o *IsDeployKeyValidConflict) Code() int {
	return 409
}

func (o *IsDeployKeyValidConflict) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidConflict", 409)
}

func (o *IsDeployKeyValidConflict) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidConflict", 409)
}

func (o *IsDeployKeyValidConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsDeployKeyValidInternalServerError creates a IsDeployKeyValidInternalServerError with default headers values
func NewIsDeployKeyValidInternalServerError() *IsDeployKeyValidInternalServerError {
	return &IsDeployKeyValidInternalServerError{}
}

/*
IsDeployKeyValidInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type IsDeployKeyValidInternalServerError struct {
}

// IsSuccess returns true when this is deploy key valid internal server error response has a 2xx status code
func (o *IsDeployKeyValidInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is deploy key valid internal server error response has a 3xx status code
func (o *IsDeployKeyValidInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is deploy key valid internal server error response has a 4xx status code
func (o *IsDeployKeyValidInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this is deploy key valid internal server error response has a 5xx status code
func (o *IsDeployKeyValidInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this is deploy key valid internal server error response a status code equal to that given
func (o *IsDeployKeyValidInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the is deploy key valid internal server error response
func (o *IsDeployKeyValidInternalServerError) Code() int {
	return 500
}

func (o *IsDeployKeyValidInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidInternalServerError", 500)
}

func (o *IsDeployKeyValidInternalServerError) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploykey-valid][%d] isDeployKeyValidInternalServerError", 500)
}

func (o *IsDeployKeyValidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
