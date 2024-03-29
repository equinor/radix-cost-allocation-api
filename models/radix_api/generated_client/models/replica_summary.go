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

// ReplicaSummary ReplicaSummary describes condition of a pod
//
// swagger:model ReplicaSummary
type ReplicaSummary struct {

	// Container started timestamp
	// Example: 2006-01-02T15:04:05Z
	ContainerStarted string `json:"containerStarted,omitempty"`

	// Created timestamp
	// Example: 2006-01-02T15:04:05Z
	Created string `json:"created,omitempty"`

	// The image the container is running.
	// Example: radixdev.azurecr.io/app-server:cdgkg
	Image string `json:"image,omitempty"`

	// ImageID of the container's image.
	// Example: radixdev.azurecr.io/app-server@sha256:d40cda01916ef63da3607c03785efabc56eb2fc2e0dab0726b1a843e9ded093f
	ImageID string `json:"imageId,omitempty"`

	// Pod name
	// Example: server-78fc8857c4-hm76l
	// Required: true
	Name *string `json:"name"`

	// RestartCount count of restarts of a component container inside a pod
	RestartCount int32 `json:"restartCount,omitempty"`

	// StatusMessage provides message describing the status of a component container inside a pod
	StatusMessage string `json:"statusMessage,omitempty"`

	// replica status
	ReplicaStatus *ReplicaStatus `json:"replicaStatus,omitempty"`

	// resources
	Resources *ResourceRequirements `json:"resources,omitempty"`
}

// Validate validates this replica summary
func (m *ReplicaSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicaSummary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ReplicaSummary) validateReplicaStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicaStatus) { // not required
		return nil
	}

	if m.ReplicaStatus != nil {
		if err := m.ReplicaStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replicaStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replicaStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicaSummary) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this replica summary based on the context it is used
func (m *ReplicaSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReplicaStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicaSummary) contextValidateReplicaStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ReplicaStatus != nil {

		if swag.IsZero(m.ReplicaStatus) { // not required
			return nil
		}

		if err := m.ReplicaStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replicaStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replicaStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicaSummary) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {

		if swag.IsZero(m.Resources) { // not required
			return nil
		}

		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplicaSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicaSummary) UnmarshalBinary(b []byte) error {
	var res ReplicaSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
