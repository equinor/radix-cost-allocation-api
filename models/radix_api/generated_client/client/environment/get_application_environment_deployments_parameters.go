// Code generated by go-swagger; DO NOT EDIT.

package environment

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
	"github.com/go-openapi/swag"
)

// NewGetApplicationEnvironmentDeploymentsParams creates a new GetApplicationEnvironmentDeploymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetApplicationEnvironmentDeploymentsParams() *GetApplicationEnvironmentDeploymentsParams {
	return &GetApplicationEnvironmentDeploymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationEnvironmentDeploymentsParamsWithTimeout creates a new GetApplicationEnvironmentDeploymentsParams object
// with the ability to set a timeout on a request.
func NewGetApplicationEnvironmentDeploymentsParamsWithTimeout(timeout time.Duration) *GetApplicationEnvironmentDeploymentsParams {
	return &GetApplicationEnvironmentDeploymentsParams{
		timeout: timeout,
	}
}

// NewGetApplicationEnvironmentDeploymentsParamsWithContext creates a new GetApplicationEnvironmentDeploymentsParams object
// with the ability to set a context for a request.
func NewGetApplicationEnvironmentDeploymentsParamsWithContext(ctx context.Context) *GetApplicationEnvironmentDeploymentsParams {
	return &GetApplicationEnvironmentDeploymentsParams{
		Context: ctx,
	}
}

// NewGetApplicationEnvironmentDeploymentsParamsWithHTTPClient creates a new GetApplicationEnvironmentDeploymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetApplicationEnvironmentDeploymentsParamsWithHTTPClient(client *http.Client) *GetApplicationEnvironmentDeploymentsParams {
	return &GetApplicationEnvironmentDeploymentsParams{
		HTTPClient: client,
	}
}

/*
GetApplicationEnvironmentDeploymentsParams contains all the parameters to send to the API endpoint

	for the get application environment deployments operation.

	Typically these are written to a http.Request.
*/
type GetApplicationEnvironmentDeploymentsParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of Radix application
	*/
	AppName string

	/* EnvName.

	   environment of Radix application
	*/
	EnvName string

	/* Latest.

	   indicator to allow only listing the latest
	*/
	Latest *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get application environment deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationEnvironmentDeploymentsParams) WithDefaults() *GetApplicationEnvironmentDeploymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get application environment deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationEnvironmentDeploymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) WithTimeout(timeout time.Duration) *GetApplicationEnvironmentDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) WithContext(ctx context.Context) *GetApplicationEnvironmentDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) WithHTTPClient(client *http.Client) *GetApplicationEnvironmentDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) WithImpersonateGroup(impersonateGroup *string) *GetApplicationEnvironmentDeploymentsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) WithImpersonateUser(impersonateUser *string) *GetApplicationEnvironmentDeploymentsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) WithAppName(appName string) *GetApplicationEnvironmentDeploymentsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) WithEnvName(envName string) *GetApplicationEnvironmentDeploymentsParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithLatest adds the latest to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) WithLatest(latest *bool) *GetApplicationEnvironmentDeploymentsParams {
	o.SetLatest(latest)
	return o
}

// SetLatest adds the latest to the get application environment deployments params
func (o *GetApplicationEnvironmentDeploymentsParams) SetLatest(latest *bool) {
	o.Latest = latest
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationEnvironmentDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Latest != nil {

		// query param latest
		var qrLatest bool

		if o.Latest != nil {
			qrLatest = *o.Latest
		}
		qLatest := swag.FormatBool(qrLatest)
		if qLatest != "" {

			if err := r.SetQueryParam("latest", qLatest); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
