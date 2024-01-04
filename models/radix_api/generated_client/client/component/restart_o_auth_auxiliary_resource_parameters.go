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

// NewRestartOAuthAuxiliaryResourceParams creates a new RestartOAuthAuxiliaryResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestartOAuthAuxiliaryResourceParams() *RestartOAuthAuxiliaryResourceParams {
	return &RestartOAuthAuxiliaryResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestartOAuthAuxiliaryResourceParamsWithTimeout creates a new RestartOAuthAuxiliaryResourceParams object
// with the ability to set a timeout on a request.
func NewRestartOAuthAuxiliaryResourceParamsWithTimeout(timeout time.Duration) *RestartOAuthAuxiliaryResourceParams {
	return &RestartOAuthAuxiliaryResourceParams{
		timeout: timeout,
	}
}

// NewRestartOAuthAuxiliaryResourceParamsWithContext creates a new RestartOAuthAuxiliaryResourceParams object
// with the ability to set a context for a request.
func NewRestartOAuthAuxiliaryResourceParamsWithContext(ctx context.Context) *RestartOAuthAuxiliaryResourceParams {
	return &RestartOAuthAuxiliaryResourceParams{
		Context: ctx,
	}
}

// NewRestartOAuthAuxiliaryResourceParamsWithHTTPClient creates a new RestartOAuthAuxiliaryResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestartOAuthAuxiliaryResourceParamsWithHTTPClient(client *http.Client) *RestartOAuthAuxiliaryResourceParams {
	return &RestartOAuthAuxiliaryResourceParams{
		HTTPClient: client,
	}
}

/*
RestartOAuthAuxiliaryResourceParams contains all the parameters to send to the API endpoint

	for the restart o auth auxiliary resource operation.

	Typically these are written to a http.Request.
*/
type RestartOAuthAuxiliaryResourceParams struct {

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

// WithDefaults hydrates default values in the restart o auth auxiliary resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartOAuthAuxiliaryResourceParams) WithDefaults() *RestartOAuthAuxiliaryResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restart o auth auxiliary resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartOAuthAuxiliaryResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) WithTimeout(timeout time.Duration) *RestartOAuthAuxiliaryResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) WithContext(ctx context.Context) *RestartOAuthAuxiliaryResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) WithHTTPClient(client *http.Client) *RestartOAuthAuxiliaryResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) WithImpersonateGroup(impersonateGroup *string) *RestartOAuthAuxiliaryResourceParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) WithImpersonateUser(impersonateUser *string) *RestartOAuthAuxiliaryResourceParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) WithAppName(appName string) *RestartOAuthAuxiliaryResourceParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithComponentName adds the componentName to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) WithComponentName(componentName string) *RestartOAuthAuxiliaryResourceParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithEnvName adds the envName to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) WithEnvName(envName string) *RestartOAuthAuxiliaryResourceParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the restart o auth auxiliary resource params
func (o *RestartOAuthAuxiliaryResourceParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *RestartOAuthAuxiliaryResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
