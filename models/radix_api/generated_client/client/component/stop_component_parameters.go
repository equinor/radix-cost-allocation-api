// Code generated by go-swagger; DO NOT EDIT.

package component

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

// NewStopComponentParams creates a new StopComponentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopComponentParams() *StopComponentParams {
	return &StopComponentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopComponentParamsWithTimeout creates a new StopComponentParams object
// with the ability to set a timeout on a request.
func NewStopComponentParamsWithTimeout(timeout time.Duration) *StopComponentParams {
	return &StopComponentParams{
		timeout: timeout,
	}
}

// NewStopComponentParamsWithContext creates a new StopComponentParams object
// with the ability to set a context for a request.
func NewStopComponentParamsWithContext(ctx context.Context) *StopComponentParams {
	return &StopComponentParams{
		Context: ctx,
	}
}

// NewStopComponentParamsWithHTTPClient creates a new StopComponentParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopComponentParamsWithHTTPClient(client *http.Client) *StopComponentParams {
	return &StopComponentParams{
		HTTPClient: client,
	}
}

/*
StopComponentParams contains all the parameters to send to the API endpoint

	for the stop component operation.

	Typically these are written to a http.Request.
*/
type StopComponentParams struct {

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

	/* ComponentName.

	   Name of component
	*/
	ComponentName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopComponentParams) WithDefaults() *StopComponentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopComponentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop component params
func (o *StopComponentParams) WithTimeout(timeout time.Duration) *StopComponentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop component params
func (o *StopComponentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop component params
func (o *StopComponentParams) WithContext(ctx context.Context) *StopComponentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop component params
func (o *StopComponentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop component params
func (o *StopComponentParams) WithHTTPClient(client *http.Client) *StopComponentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop component params
func (o *StopComponentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the stop component params
func (o *StopComponentParams) WithImpersonateGroup(impersonateGroup *string) *StopComponentParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the stop component params
func (o *StopComponentParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the stop component params
func (o *StopComponentParams) WithImpersonateUser(impersonateUser *string) *StopComponentParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the stop component params
func (o *StopComponentParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the stop component params
func (o *StopComponentParams) WithAppName(appName string) *StopComponentParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the stop component params
func (o *StopComponentParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithComponentName adds the componentName to the stop component params
func (o *StopComponentParams) WithComponentName(componentName string) *StopComponentParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the stop component params
func (o *StopComponentParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithEnvName adds the envName to the stop component params
func (o *StopComponentParams) WithEnvName(envName string) *StopComponentParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the stop component params
func (o *StopComponentParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *StopComponentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param componentName
	if err := r.SetPathParam("componentName", o.ComponentName); err != nil {
		return err
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
