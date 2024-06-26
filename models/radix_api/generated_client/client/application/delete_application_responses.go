// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteApplicationReader is a Reader for the DeleteApplication structure.
type DeleteApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteApplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /applications/{appName}] deleteApplication", response, response.Code())
	}
}

// NewDeleteApplicationOK creates a DeleteApplicationOK with default headers values
func NewDeleteApplicationOK() *DeleteApplicationOK {
	return &DeleteApplicationOK{}
}

/*
DeleteApplicationOK describes a response with status code 200, with default header values.

Application deleted ok
*/
type DeleteApplicationOK struct {
}

// IsSuccess returns true when this delete application o k response has a 2xx status code
func (o *DeleteApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete application o k response has a 3xx status code
func (o *DeleteApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete application o k response has a 4xx status code
func (o *DeleteApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete application o k response has a 5xx status code
func (o *DeleteApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete application o k response a status code equal to that given
func (o *DeleteApplicationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete application o k response
func (o *DeleteApplicationOK) Code() int {
	return 200
}

func (o *DeleteApplicationOK) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationOK", 200)
}

func (o *DeleteApplicationOK) String() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationOK", 200)
}

func (o *DeleteApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationUnauthorized creates a DeleteApplicationUnauthorized with default headers values
func NewDeleteApplicationUnauthorized() *DeleteApplicationUnauthorized {
	return &DeleteApplicationUnauthorized{}
}

/*
DeleteApplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteApplicationUnauthorized struct {
}

// IsSuccess returns true when this delete application unauthorized response has a 2xx status code
func (o *DeleteApplicationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete application unauthorized response has a 3xx status code
func (o *DeleteApplicationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete application unauthorized response has a 4xx status code
func (o *DeleteApplicationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete application unauthorized response has a 5xx status code
func (o *DeleteApplicationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete application unauthorized response a status code equal to that given
func (o *DeleteApplicationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete application unauthorized response
func (o *DeleteApplicationUnauthorized) Code() int {
	return 401
}

func (o *DeleteApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationUnauthorized", 401)
}

func (o *DeleteApplicationUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationUnauthorized", 401)
}

func (o *DeleteApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationForbidden creates a DeleteApplicationForbidden with default headers values
func NewDeleteApplicationForbidden() *DeleteApplicationForbidden {
	return &DeleteApplicationForbidden{}
}

/*
DeleteApplicationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteApplicationForbidden struct {
}

// IsSuccess returns true when this delete application forbidden response has a 2xx status code
func (o *DeleteApplicationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete application forbidden response has a 3xx status code
func (o *DeleteApplicationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete application forbidden response has a 4xx status code
func (o *DeleteApplicationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete application forbidden response has a 5xx status code
func (o *DeleteApplicationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete application forbidden response a status code equal to that given
func (o *DeleteApplicationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete application forbidden response
func (o *DeleteApplicationForbidden) Code() int {
	return 403
}

func (o *DeleteApplicationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationForbidden", 403)
}

func (o *DeleteApplicationForbidden) String() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationForbidden", 403)
}

func (o *DeleteApplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationNotFound creates a DeleteApplicationNotFound with default headers values
func NewDeleteApplicationNotFound() *DeleteApplicationNotFound {
	return &DeleteApplicationNotFound{}
}

/*
DeleteApplicationNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteApplicationNotFound struct {
}

// IsSuccess returns true when this delete application not found response has a 2xx status code
func (o *DeleteApplicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete application not found response has a 3xx status code
func (o *DeleteApplicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete application not found response has a 4xx status code
func (o *DeleteApplicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete application not found response has a 5xx status code
func (o *DeleteApplicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete application not found response a status code equal to that given
func (o *DeleteApplicationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete application not found response
func (o *DeleteApplicationNotFound) Code() int {
	return 404
}

func (o *DeleteApplicationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationNotFound", 404)
}

func (o *DeleteApplicationNotFound) String() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationNotFound", 404)
}

func (o *DeleteApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
