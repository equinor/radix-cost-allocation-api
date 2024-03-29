// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplicationsSearchRequest ApplicationsSearchRequest contains the list of application names to be queried
//
// swagger:model ApplicationsSearchRequest
type ApplicationsSearchRequest struct {

	// List of application names to be returned
	// Example: ["app1","app2"]
	// Required: true
	Names []string `json:"names"`

	// include fields
	IncludeFields *ApplicationSearchIncludeFields `json:"includeFields,omitempty"`
}

// Validate validates this applications search request
func (m *ApplicationsSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationsSearchRequest) validateNames(formats strfmt.Registry) error {

	if err := validate.Required("names", "body", m.Names); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationsSearchRequest) validateIncludeFields(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludeFields) { // not required
		return nil
	}

	if m.IncludeFields != nil {
		if err := m.IncludeFields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("includeFields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("includeFields")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this applications search request based on the context it is used
func (m *ApplicationsSearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIncludeFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationsSearchRequest) contextValidateIncludeFields(ctx context.Context, formats strfmt.Registry) error {

	if m.IncludeFields != nil {

		if swag.IsZero(m.IncludeFields) { // not required
			return nil
		}

		if err := m.IncludeFields.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("includeFields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("includeFields")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationsSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationsSearchRequest) UnmarshalBinary(b []byte) error {
	var res ApplicationsSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
