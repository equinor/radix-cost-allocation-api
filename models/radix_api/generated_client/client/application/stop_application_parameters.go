// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStopApplicationParams creates a new StopApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopApplicationParams() *StopApplicationParams {
	return &StopApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopApplicationParamsWithTimeout creates a new StopApplicationParams object
// with the ability to set a timeout on a request.
func NewStopApplicationParamsWithTimeout(timeout time.Duration) *StopApplicationParams {
	return &StopApplicationParams{
		timeout: timeout,
	}
}

// NewStopApplicationParamsWithContext creates a new StopApplicationParams object
// with the ability to set a context for a request.
func NewStopApplicationParamsWithContext(ctx context.Context) *StopApplicationParams {
	return &StopApplicationParams{
		Context: ctx,
	}
}

// NewStopApplicationParamsWithHTTPClient creates a new StopApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopApplicationParamsWithHTTPClient(client *http.Client) *StopApplicationParams {
	return &StopApplicationParams{
		HTTPClient: client,
	}
}

/*
StopApplicationParams contains all the parameters to send to the API endpoint

	for the stop application operation.

	Typically these are written to a http.Request.
*/
type StopApplicationParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   Name of application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopApplicationParams) WithDefaults() *StopApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop application params
func (o *StopApplicationParams) WithTimeout(timeout time.Duration) *StopApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop application params
func (o *StopApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop application params
func (o *StopApplicationParams) WithContext(ctx context.Context) *StopApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop application params
func (o *StopApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop application params
func (o *StopApplicationParams) WithHTTPClient(client *http.Client) *StopApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop application params
func (o *StopApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the stop application params
func (o *StopApplicationParams) WithImpersonateGroup(impersonateGroup *string) *StopApplicationParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the stop application params
func (o *StopApplicationParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the stop application params
func (o *StopApplicationParams) WithImpersonateUser(impersonateUser *string) *StopApplicationParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the stop application params
func (o *StopApplicationParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the stop application params
func (o *StopApplicationParams) WithAppName(appName string) *StopApplicationParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the stop application params
func (o *StopApplicationParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *StopApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// header param Impersonate-Group
		if err := r.SetHeaderParam("Impersonate-Group", *o.ImpersonateGroup); err != nil {
			return err
		}
	}

	if o.ImpersonateUser != nil {

		// header param Impersonate-User
		if err := r.SetHeaderParam("Impersonate-User", *o.ImpersonateUser); err != nil {
			return err
		}
	}

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
