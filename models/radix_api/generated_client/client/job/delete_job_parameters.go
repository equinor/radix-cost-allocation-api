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

// NewDeleteJobParams creates a new DeleteJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteJobParams() *DeleteJobParams {
	return &DeleteJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteJobParamsWithTimeout creates a new DeleteJobParams object
// with the ability to set a timeout on a request.
func NewDeleteJobParamsWithTimeout(timeout time.Duration) *DeleteJobParams {
	return &DeleteJobParams{
		timeout: timeout,
	}
}

// NewDeleteJobParamsWithContext creates a new DeleteJobParams object
// with the ability to set a context for a request.
func NewDeleteJobParamsWithContext(ctx context.Context) *DeleteJobParams {
	return &DeleteJobParams{
		Context: ctx,
	}
}

// NewDeleteJobParamsWithHTTPClient creates a new DeleteJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteJobParamsWithHTTPClient(client *http.Client) *DeleteJobParams {
	return &DeleteJobParams{
		HTTPClient: client,
	}
}

/*
DeleteJobParams contains all the parameters to send to the API endpoint

	for the delete job operation.

	Typically these are written to a http.Request.
*/
type DeleteJobParams struct {

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

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	/* JobComponentName.

	   Name of job-component
	*/
	JobComponentName string

	/* JobName.

	   Name of job
	*/
	JobName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteJobParams) WithDefaults() *DeleteJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete job params
func (o *DeleteJobParams) WithTimeout(timeout time.Duration) *DeleteJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete job params
func (o *DeleteJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete job params
func (o *DeleteJobParams) WithContext(ctx context.Context) *DeleteJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete job params
func (o *DeleteJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete job params
func (o *DeleteJobParams) WithHTTPClient(client *http.Client) *DeleteJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete job params
func (o *DeleteJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the delete job params
func (o *DeleteJobParams) WithImpersonateGroup(impersonateGroup *string) *DeleteJobParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the delete job params
func (o *DeleteJobParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the delete job params
func (o *DeleteJobParams) WithImpersonateUser(impersonateUser *string) *DeleteJobParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the delete job params
func (o *DeleteJobParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the delete job params
func (o *DeleteJobParams) WithAppName(appName string) *DeleteJobParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the delete job params
func (o *DeleteJobParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the delete job params
func (o *DeleteJobParams) WithEnvName(envName string) *DeleteJobParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the delete job params
func (o *DeleteJobParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithJobComponentName adds the jobComponentName to the delete job params
func (o *DeleteJobParams) WithJobComponentName(jobComponentName string) *DeleteJobParams {
	o.SetJobComponentName(jobComponentName)
	return o
}

// SetJobComponentName adds the jobComponentName to the delete job params
func (o *DeleteJobParams) SetJobComponentName(jobComponentName string) {
	o.JobComponentName = jobComponentName
}

// WithJobName adds the jobName to the delete job params
func (o *DeleteJobParams) WithJobName(jobName string) *DeleteJobParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the delete job params
func (o *DeleteJobParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	// path param jobComponentName
	if err := r.SetPathParam("jobComponentName", o.JobComponentName); err != nil {
		return err
	}

	// path param jobName
	if err := r.SetPathParam("jobName", o.JobName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
