// Code generated by go-swagger; DO NOT EDIT.

package platform

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

	"github.com/equinor/radix-cost-allocation-api/models/radix_api/generated_client/models"
)

// NewRegisterApplicationParams creates a new RegisterApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterApplicationParams() *RegisterApplicationParams {
	return &RegisterApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterApplicationParamsWithTimeout creates a new RegisterApplicationParams object
// with the ability to set a timeout on a request.
func NewRegisterApplicationParamsWithTimeout(timeout time.Duration) *RegisterApplicationParams {
	return &RegisterApplicationParams{
		timeout: timeout,
	}
}

// NewRegisterApplicationParamsWithContext creates a new RegisterApplicationParams object
// with the ability to set a context for a request.
func NewRegisterApplicationParamsWithContext(ctx context.Context) *RegisterApplicationParams {
	return &RegisterApplicationParams{
		Context: ctx,
	}
}

// NewRegisterApplicationParamsWithHTTPClient creates a new RegisterApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterApplicationParamsWithHTTPClient(client *http.Client) *RegisterApplicationParams {
	return &RegisterApplicationParams{
		HTTPClient: client,
	}
}

/*
RegisterApplicationParams contains all the parameters to send to the API endpoint

	for the register application operation.

	Typically these are written to a http.Request.
*/
type RegisterApplicationParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* ApplicationRegistration.

	   Request for an Application to register
	*/
	ApplicationRegistration *models.ApplicationRegistrationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterApplicationParams) WithDefaults() *RegisterApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register application params
func (o *RegisterApplicationParams) WithTimeout(timeout time.Duration) *RegisterApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register application params
func (o *RegisterApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register application params
func (o *RegisterApplicationParams) WithContext(ctx context.Context) *RegisterApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register application params
func (o *RegisterApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register application params
func (o *RegisterApplicationParams) WithHTTPClient(client *http.Client) *RegisterApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register application params
func (o *RegisterApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the register application params
func (o *RegisterApplicationParams) WithImpersonateGroup(impersonateGroup []string) *RegisterApplicationParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the register application params
func (o *RegisterApplicationParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the register application params
func (o *RegisterApplicationParams) WithImpersonateUser(impersonateUser *string) *RegisterApplicationParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the register application params
func (o *RegisterApplicationParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithApplicationRegistration adds the applicationRegistration to the register application params
func (o *RegisterApplicationParams) WithApplicationRegistration(applicationRegistration *models.ApplicationRegistrationRequest) *RegisterApplicationParams {
	o.SetApplicationRegistration(applicationRegistration)
	return o
}

// SetApplicationRegistration adds the applicationRegistration to the register application params
func (o *RegisterApplicationParams) SetApplicationRegistration(applicationRegistration *models.ApplicationRegistrationRequest) {
	o.ApplicationRegistration = applicationRegistration
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// binding items for Impersonate-Group
		joinedImpersonateGroup := o.bindParamImpersonateGroup(reg)

		// header array param Impersonate-Group
		if len(joinedImpersonateGroup) > 0 {
			if err := r.SetHeaderParam("Impersonate-Group", joinedImpersonateGroup[0]); err != nil {
				return err
			}
		}
	}

	if o.ImpersonateUser != nil {

		// header param Impersonate-User
		if err := r.SetHeaderParam("Impersonate-User", *o.ImpersonateUser); err != nil {
			return err
		}
	}
	if o.ApplicationRegistration != nil {
		if err := r.SetBodyParam(o.ApplicationRegistration); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamRegisterApplication binds the parameter Impersonate-Group
func (o *RegisterApplicationParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
	impersonateGroupIR := o.ImpersonateGroup

	var impersonateGroupIC []string
	for _, impersonateGroupIIR := range impersonateGroupIR { // explode []string

		impersonateGroupIIV := impersonateGroupIIR // string as string
		impersonateGroupIC = append(impersonateGroupIC, impersonateGroupIIV)
	}

	// items.CollectionFormat: ""
	impersonateGroupIS := swag.JoinByFormat(impersonateGroupIC, "")

	return impersonateGroupIS
}