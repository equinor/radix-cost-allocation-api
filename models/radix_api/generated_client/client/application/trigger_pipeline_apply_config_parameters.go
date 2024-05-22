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

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models"
)

// NewTriggerPipelineApplyConfigParams creates a new TriggerPipelineApplyConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriggerPipelineApplyConfigParams() *TriggerPipelineApplyConfigParams {
	return &TriggerPipelineApplyConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerPipelineApplyConfigParamsWithTimeout creates a new TriggerPipelineApplyConfigParams object
// with the ability to set a timeout on a request.
func NewTriggerPipelineApplyConfigParamsWithTimeout(timeout time.Duration) *TriggerPipelineApplyConfigParams {
	return &TriggerPipelineApplyConfigParams{
		timeout: timeout,
	}
}

// NewTriggerPipelineApplyConfigParamsWithContext creates a new TriggerPipelineApplyConfigParams object
// with the ability to set a context for a request.
func NewTriggerPipelineApplyConfigParamsWithContext(ctx context.Context) *TriggerPipelineApplyConfigParams {
	return &TriggerPipelineApplyConfigParams{
		Context: ctx,
	}
}

// NewTriggerPipelineApplyConfigParamsWithHTTPClient creates a new TriggerPipelineApplyConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriggerPipelineApplyConfigParamsWithHTTPClient(client *http.Client) *TriggerPipelineApplyConfigParams {
	return &TriggerPipelineApplyConfigParams{
		HTTPClient: client,
	}
}

/*
TriggerPipelineApplyConfigParams contains all the parameters to send to the API endpoint

	for the trigger pipeline apply config operation.

	Typically these are written to a http.Request.
*/
type TriggerPipelineApplyConfigParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* PipelineParametersApplyConfig.

	   Pipeline parameters
	*/
	PipelineParametersApplyConfig *models.PipelineParametersApplyConfig

	/* AppName.

	   Name of application
	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the trigger pipeline apply config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerPipelineApplyConfigParams) WithDefaults() *TriggerPipelineApplyConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trigger pipeline apply config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerPipelineApplyConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) WithTimeout(timeout time.Duration) *TriggerPipelineApplyConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) WithContext(ctx context.Context) *TriggerPipelineApplyConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) WithHTTPClient(client *http.Client) *TriggerPipelineApplyConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) WithImpersonateGroup(impersonateGroup *string) *TriggerPipelineApplyConfigParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) WithImpersonateUser(impersonateUser *string) *TriggerPipelineApplyConfigParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithPipelineParametersApplyConfig adds the pipelineParametersApplyConfig to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) WithPipelineParametersApplyConfig(pipelineParametersApplyConfig *models.PipelineParametersApplyConfig) *TriggerPipelineApplyConfigParams {
	o.SetPipelineParametersApplyConfig(pipelineParametersApplyConfig)
	return o
}

// SetPipelineParametersApplyConfig adds the pipelineParametersApplyConfig to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) SetPipelineParametersApplyConfig(pipelineParametersApplyConfig *models.PipelineParametersApplyConfig) {
	o.PipelineParametersApplyConfig = pipelineParametersApplyConfig
}

// WithAppName adds the appName to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) WithAppName(appName string) *TriggerPipelineApplyConfigParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the trigger pipeline apply config params
func (o *TriggerPipelineApplyConfigParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerPipelineApplyConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.PipelineParametersApplyConfig != nil {
		if err := r.SetBodyParam(o.PipelineParametersApplyConfig); err != nil {
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
