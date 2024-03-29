// Code generated by go-swagger; DO NOT EDIT.

package job

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

// NewGetBatchParams creates a new GetBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBatchParams() *GetBatchParams {
	return &GetBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchParamsWithTimeout creates a new GetBatchParams object
// with the ability to set a timeout on a request.
func NewGetBatchParamsWithTimeout(timeout time.Duration) *GetBatchParams {
	return &GetBatchParams{
		timeout: timeout,
	}
}

// NewGetBatchParamsWithContext creates a new GetBatchParams object
// with the ability to set a context for a request.
func NewGetBatchParamsWithContext(ctx context.Context) *GetBatchParams {
	return &GetBatchParams{
		Context: ctx,
	}
}

// NewGetBatchParamsWithHTTPClient creates a new GetBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBatchParamsWithHTTPClient(client *http.Client) *GetBatchParams {
	return &GetBatchParams{
		HTTPClient: client,
	}
}

/*
GetBatchParams contains all the parameters to send to the API endpoint

	for the get batch operation.

	Typically these are written to a http.Request.
*/
type GetBatchParams struct {

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

	/* BatchName.

	   Name of batch
	*/
	BatchName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	/* JobComponentName.

	   Name of job-component
	*/
	JobComponentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchParams) WithDefaults() *GetBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get batch params
func (o *GetBatchParams) WithTimeout(timeout time.Duration) *GetBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batch params
func (o *GetBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batch params
func (o *GetBatchParams) WithContext(ctx context.Context) *GetBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batch params
func (o *GetBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batch params
func (o *GetBatchParams) WithHTTPClient(client *http.Client) *GetBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batch params
func (o *GetBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get batch params
func (o *GetBatchParams) WithImpersonateGroup(impersonateGroup *string) *GetBatchParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get batch params
func (o *GetBatchParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get batch params
func (o *GetBatchParams) WithImpersonateUser(impersonateUser *string) *GetBatchParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get batch params
func (o *GetBatchParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get batch params
func (o *GetBatchParams) WithAppName(appName string) *GetBatchParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get batch params
func (o *GetBatchParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithBatchName adds the batchName to the get batch params
func (o *GetBatchParams) WithBatchName(batchName string) *GetBatchParams {
	o.SetBatchName(batchName)
	return o
}

// SetBatchName adds the batchName to the get batch params
func (o *GetBatchParams) SetBatchName(batchName string) {
	o.BatchName = batchName
}

// WithEnvName adds the envName to the get batch params
func (o *GetBatchParams) WithEnvName(envName string) *GetBatchParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the get batch params
func (o *GetBatchParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithJobComponentName adds the jobComponentName to the get batch params
func (o *GetBatchParams) WithJobComponentName(jobComponentName string) *GetBatchParams {
	o.SetJobComponentName(jobComponentName)
	return o
}

// SetJobComponentName adds the jobComponentName to the get batch params
func (o *GetBatchParams) SetJobComponentName(jobComponentName string) {
	o.JobComponentName = jobComponentName
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param batchName
	if err := r.SetPathParam("batchName", o.BatchName); err != nil {
		return err
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	// path param jobComponentName
	if err := r.SetPathParam("jobComponentName", o.JobComponentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
