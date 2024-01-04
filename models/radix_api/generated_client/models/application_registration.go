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

// ApplicationRegistration ApplicationRegistration describe an application
//
// swagger:model ApplicationRegistration
type ApplicationRegistration struct {

	// AdGroups the groups that should be able to access the application
	// Required: true
	AdGroups []string `json:"adGroups"`

	// ConfigBranch information
	// Required: true
	ConfigBranch *string `json:"configBranch"`

	// ConfigurationItem is an identifier for an entity in a configuration management solution such as a CMDB.
	// ITIL defines a CI as any component that needs to be managed in order to deliver an IT Service
	// Ref: https://en.wikipedia.org/wiki/Configuration_item
	ConfigurationItem string `json:"configurationItem,omitempty"`

	// Owner of the application (email). Can be a single person or a shared group email
	// Required: true
	Creator *string `json:"creator"`

	// Name the unique name of the Radix application
	// Example: radix-canary-golang
	// Required: true
	Name *string `json:"name"`

	// Owner of the application (email). Can be a single person or a shared group email
	// Required: true
	Owner *string `json:"owner"`

	// radixconfig.yaml file name and path, starting from the GitHub repository root (without leading slash)
	RadixConfigFullName string `json:"radixConfigFullName,omitempty"`

	// ReaderAdGroups the groups that should be able to read the application
	ReaderAdGroups []string `json:"readerAdGroups"`

	// Repository the github repository
	// Example: https://github.com/equinor/radix-canary-golang
	// Required: true
	Repository *string `json:"repository"`

	// SharedSecret the shared secret of the webhook
	// Required: true
	SharedSecret *string `json:"sharedSecret"`

	// WBS information
	WBS string `json:"wbs,omitempty"`
}

// Validate validates this application registration
func (m *ApplicationRegistration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationRegistration) validateAdGroups(formats strfmt.Registry) error {

	if err := validate.Required("adGroups", "body", m.AdGroups); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRegistration) validateConfigBranch(formats strfmt.Registry) error {

	if err := validate.Required("configBranch", "body", m.ConfigBranch); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRegistration) validateCreator(formats strfmt.Registry) error {

	if err := validate.Required("creator", "body", m.Creator); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRegistration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRegistration) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRegistration) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationRegistration) validateSharedSecret(formats strfmt.Registry) error {

	if err := validate.Required("sharedSecret", "body", m.SharedSecret); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this application registration based on context it is used
func (m *ApplicationRegistration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationRegistration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationRegistration) UnmarshalBinary(b []byte) error {
	var res ApplicationRegistration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
