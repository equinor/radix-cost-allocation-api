// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models"
)

// GetDeployKeyAndSecretReader is a Reader for the GetDeployKeyAndSecret structure.
type GetDeployKeyAndSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeployKeyAndSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeployKeyAndSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeployKeyAndSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeployKeyAndSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/deploy-key-and-secret] getDeployKeyAndSecret", response, response.Code())
	}
}

// NewGetDeployKeyAndSecretOK creates a GetDeployKeyAndSecretOK with default headers values
func NewGetDeployKeyAndSecretOK() *GetDeployKeyAndSecretOK {
	return &GetDeployKeyAndSecretOK{}
}

/*
GetDeployKeyAndSecretOK describes a response with status code 200, with default header values.

Successful get deploy key and secret
*/
type GetDeployKeyAndSecretOK struct {
	Payload *models.DeployKeyAndSecret
}

// IsSuccess returns true when this get deploy key and secret o k response has a 2xx status code
func (o *GetDeployKeyAndSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deploy key and secret o k response has a 3xx status code
func (o *GetDeployKeyAndSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deploy key and secret o k response has a 4xx status code
func (o *GetDeployKeyAndSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deploy key and secret o k response has a 5xx status code
func (o *GetDeployKeyAndSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deploy key and secret o k response a status code equal to that given
func (o *GetDeployKeyAndSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deploy key and secret o k response
func (o *GetDeployKeyAndSecretOK) Code() int {
	return 200
}

func (o *GetDeployKeyAndSecretOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploy-key-and-secret][%d] getDeployKeyAndSecretOK  %+v", 200, o.Payload)
}

func (o *GetDeployKeyAndSecretOK) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploy-key-and-secret][%d] getDeployKeyAndSecretOK  %+v", 200, o.Payload)
}

func (o *GetDeployKeyAndSecretOK) GetPayload() *models.DeployKeyAndSecret {
	return o.Payload
}

func (o *GetDeployKeyAndSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeployKeyAndSecret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeployKeyAndSecretUnauthorized creates a GetDeployKeyAndSecretUnauthorized with default headers values
func NewGetDeployKeyAndSecretUnauthorized() *GetDeployKeyAndSecretUnauthorized {
	return &GetDeployKeyAndSecretUnauthorized{}
}

/*
GetDeployKeyAndSecretUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeployKeyAndSecretUnauthorized struct {
}

// IsSuccess returns true when this get deploy key and secret unauthorized response has a 2xx status code
func (o *GetDeployKeyAndSecretUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deploy key and secret unauthorized response has a 3xx status code
func (o *GetDeployKeyAndSecretUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deploy key and secret unauthorized response has a 4xx status code
func (o *GetDeployKeyAndSecretUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deploy key and secret unauthorized response has a 5xx status code
func (o *GetDeployKeyAndSecretUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deploy key and secret unauthorized response a status code equal to that given
func (o *GetDeployKeyAndSecretUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get deploy key and secret unauthorized response
func (o *GetDeployKeyAndSecretUnauthorized) Code() int {
	return 401
}

func (o *GetDeployKeyAndSecretUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploy-key-and-secret][%d] getDeployKeyAndSecretUnauthorized ", 401)
}

func (o *GetDeployKeyAndSecretUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploy-key-and-secret][%d] getDeployKeyAndSecretUnauthorized ", 401)
}

func (o *GetDeployKeyAndSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeployKeyAndSecretNotFound creates a GetDeployKeyAndSecretNotFound with default headers values
func NewGetDeployKeyAndSecretNotFound() *GetDeployKeyAndSecretNotFound {
	return &GetDeployKeyAndSecretNotFound{}
}

/*
GetDeployKeyAndSecretNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetDeployKeyAndSecretNotFound struct {
}

// IsSuccess returns true when this get deploy key and secret not found response has a 2xx status code
func (o *GetDeployKeyAndSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deploy key and secret not found response has a 3xx status code
func (o *GetDeployKeyAndSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deploy key and secret not found response has a 4xx status code
func (o *GetDeployKeyAndSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deploy key and secret not found response has a 5xx status code
func (o *GetDeployKeyAndSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deploy key and secret not found response a status code equal to that given
func (o *GetDeployKeyAndSecretNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deploy key and secret not found response
func (o *GetDeployKeyAndSecretNotFound) Code() int {
	return 404
}

func (o *GetDeployKeyAndSecretNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploy-key-and-secret][%d] getDeployKeyAndSecretNotFound ", 404)
}

func (o *GetDeployKeyAndSecretNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deploy-key-and-secret][%d] getDeployKeyAndSecretNotFound ", 404)
}

func (o *GetDeployKeyAndSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}