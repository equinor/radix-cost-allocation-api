// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

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

// NewRerunApplicationJobParams creates a new RerunApplicationJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRerunApplicationJobParams() *RerunApplicationJobParams {
	return &RerunApplicationJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRerunApplicationJobParamsWithTimeout creates a new RerunApplicationJobParams object
// with the ability to set a timeout on a request.
func NewRerunApplicationJobParamsWithTimeout(timeout time.Duration) *RerunApplicationJobParams {
	return &RerunApplicationJobParams{
		timeout: timeout,
	}
}

// NewRerunApplicationJobParamsWithContext creates a new RerunApplicationJobParams object
// with the ability to set a context for a request.
func NewRerunApplicationJobParamsWithContext(ctx context.Context) *RerunApplicationJobParams {
	return &RerunApplicationJobParams{
		Context: ctx,
	}
}

// NewRerunApplicationJobParamsWithHTTPClient creates a new RerunApplicationJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewRerunApplicationJobParamsWithHTTPClient(client *http.Client) *RerunApplicationJobParams {
	return &RerunApplicationJobParams{
		HTTPClient: client,
	}
}

/*
RerunApplicationJobParams contains all the parameters to send to the API endpoint

	for the rerun application job operation.

	Typically these are written to a http.Request.
*/
type RerunApplicationJobParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of application
	*/
	AppName string

	/* JobName.

	   name of job
	*/
	JobName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rerun application job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RerunApplicationJobParams) WithDefaults() *RerunApplicationJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rerun application job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RerunApplicationJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rerun application job params
func (o *RerunApplicationJobParams) WithTimeout(timeout time.Duration) *RerunApplicationJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rerun application job params
func (o *RerunApplicationJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rerun application job params
func (o *RerunApplicationJobParams) WithContext(ctx context.Context) *RerunApplicationJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rerun application job params
func (o *RerunApplicationJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rerun application job params
func (o *RerunApplicationJobParams) WithHTTPClient(client *http.Client) *RerunApplicationJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rerun application job params
func (o *RerunApplicationJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the rerun application job params
func (o *RerunApplicationJobParams) WithImpersonateGroup(impersonateGroup *string) *RerunApplicationJobParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the rerun application job params
func (o *RerunApplicationJobParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the rerun application job params
func (o *RerunApplicationJobParams) WithImpersonateUser(impersonateUser *string) *RerunApplicationJobParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the rerun application job params
func (o *RerunApplicationJobParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the rerun application job params
func (o *RerunApplicationJobParams) WithAppName(appName string) *RerunApplicationJobParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the rerun application job params
func (o *RerunApplicationJobParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithJobName adds the jobName to the rerun application job params
func (o *RerunApplicationJobParams) WithJobName(jobName string) *RerunApplicationJobParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the rerun application job params
func (o *RerunApplicationJobParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WriteToRequest writes these params to a swagger request
func (o *RerunApplicationJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param jobName
	if err := r.SetPathParam("jobName", o.JobName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
