// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models"
)

// CopyBatchReader is a Reader for the CopyBatch structure.
type CopyBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCopyBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCopyBatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCopyBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCopyBatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCopyBatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy] copyBatch", response, response.Code())
	}
}

// NewCopyBatchOK creates a CopyBatchOK with default headers values
func NewCopyBatchOK() *CopyBatchOK {
	return &CopyBatchOK{}
}

/*
CopyBatchOK describes a response with status code 200, with default header values.

Success
*/
type CopyBatchOK struct {
	Payload *models.ScheduledBatchSummary
}

// IsSuccess returns true when this copy batch o k response has a 2xx status code
func (o *CopyBatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this copy batch o k response has a 3xx status code
func (o *CopyBatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy batch o k response has a 4xx status code
func (o *CopyBatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this copy batch o k response has a 5xx status code
func (o *CopyBatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this copy batch o k response a status code equal to that given
func (o *CopyBatchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the copy batch o k response
func (o *CopyBatchOK) Code() int {
	return 200
}

func (o *CopyBatchOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchOK  %+v", 200, o.Payload)
}

func (o *CopyBatchOK) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchOK  %+v", 200, o.Payload)
}

func (o *CopyBatchOK) GetPayload() *models.ScheduledBatchSummary {
	return o.Payload
}

func (o *CopyBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduledBatchSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyBatchBadRequest creates a CopyBatchBadRequest with default headers values
func NewCopyBatchBadRequest() *CopyBatchBadRequest {
	return &CopyBatchBadRequest{}
}

/*
CopyBatchBadRequest describes a response with status code 400, with default header values.

Invalid batch
*/
type CopyBatchBadRequest struct {
}

// IsSuccess returns true when this copy batch bad request response has a 2xx status code
func (o *CopyBatchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy batch bad request response has a 3xx status code
func (o *CopyBatchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy batch bad request response has a 4xx status code
func (o *CopyBatchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy batch bad request response has a 5xx status code
func (o *CopyBatchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this copy batch bad request response a status code equal to that given
func (o *CopyBatchBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the copy batch bad request response
func (o *CopyBatchBadRequest) Code() int {
	return 400
}

func (o *CopyBatchBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchBadRequest ", 400)
}

func (o *CopyBatchBadRequest) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchBadRequest ", 400)
}

func (o *CopyBatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyBatchUnauthorized creates a CopyBatchUnauthorized with default headers values
func NewCopyBatchUnauthorized() *CopyBatchUnauthorized {
	return &CopyBatchUnauthorized{}
}

/*
CopyBatchUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CopyBatchUnauthorized struct {
}

// IsSuccess returns true when this copy batch unauthorized response has a 2xx status code
func (o *CopyBatchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy batch unauthorized response has a 3xx status code
func (o *CopyBatchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy batch unauthorized response has a 4xx status code
func (o *CopyBatchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy batch unauthorized response has a 5xx status code
func (o *CopyBatchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this copy batch unauthorized response a status code equal to that given
func (o *CopyBatchUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the copy batch unauthorized response
func (o *CopyBatchUnauthorized) Code() int {
	return 401
}

func (o *CopyBatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchUnauthorized ", 401)
}

func (o *CopyBatchUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchUnauthorized ", 401)
}

func (o *CopyBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyBatchForbidden creates a CopyBatchForbidden with default headers values
func NewCopyBatchForbidden() *CopyBatchForbidden {
	return &CopyBatchForbidden{}
}

/*
CopyBatchForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CopyBatchForbidden struct {
}

// IsSuccess returns true when this copy batch forbidden response has a 2xx status code
func (o *CopyBatchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy batch forbidden response has a 3xx status code
func (o *CopyBatchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy batch forbidden response has a 4xx status code
func (o *CopyBatchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy batch forbidden response has a 5xx status code
func (o *CopyBatchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this copy batch forbidden response a status code equal to that given
func (o *CopyBatchForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the copy batch forbidden response
func (o *CopyBatchForbidden) Code() int {
	return 403
}

func (o *CopyBatchForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchForbidden ", 403)
}

func (o *CopyBatchForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchForbidden ", 403)
}

func (o *CopyBatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyBatchNotFound creates a CopyBatchNotFound with default headers values
func NewCopyBatchNotFound() *CopyBatchNotFound {
	return &CopyBatchNotFound{}
}

/*
CopyBatchNotFound describes a response with status code 404, with default header values.

Not found
*/
type CopyBatchNotFound struct {
}

// IsSuccess returns true when this copy batch not found response has a 2xx status code
func (o *CopyBatchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy batch not found response has a 3xx status code
func (o *CopyBatchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy batch not found response has a 4xx status code
func (o *CopyBatchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy batch not found response has a 5xx status code
func (o *CopyBatchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this copy batch not found response a status code equal to that given
func (o *CopyBatchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the copy batch not found response
func (o *CopyBatchNotFound) Code() int {
	return 404
}

func (o *CopyBatchNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchNotFound ", 404)
}

func (o *CopyBatchNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy][%d] copyBatchNotFound ", 404)
}

func (o *CopyBatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
