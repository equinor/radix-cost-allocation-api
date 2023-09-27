// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScheduledJobSummary ScheduledJobSummary holds general information about scheduled job
//
// swagger:model ScheduledJobSummary
type ScheduledJobSummary struct {

	// BackoffLimit Amount of retries due to a logical error in configuration etc.
	// Example: 1
	// Required: true
	BackoffLimit *int32 `json:"backoffLimit"`

	// BatchName Batch name, if any
	// Example: \"batch-abc\
	BatchName string `json:"batchName,omitempty"`

	// Created timestamp
	// Example: 2006-01-02T15:04:05Z
	Created string `json:"created,omitempty"`

	// DeploymentName name of RadixDeployment for the job
	DeploymentName string `json:"deploymentName,omitempty"`

	// Ended timestamp
	// Example: 2006-01-02T15:04:05Z
	Ended string `json:"ended,omitempty"`

	// FailedCount is the number of times the job has failed
	// Example: 1
	// Required: true
	FailedCount *int32 `json:"failedCount"`

	// JobId JobId, if any
	// Example: \"job1\
	JobID string `json:"jobId,omitempty"`

	// Message of a status, if any, of the job
	// Example: \"Error occurred\
	Message string `json:"message,omitempty"`

	// Name of the scheduled job
	// Example: job-component-20181029135644-algpv-6hznh
	Name string `json:"name,omitempty"`

	// Array of ReplicaSummary
	ReplicaList []*ReplicaSummary `json:"replicaList"`

	// Timestamp of the job restart, if applied.
	// +optional
	Restart string `json:"Restart,omitempty"`

	// Started timestamp
	// Example: 2006-01-02T15:04:05Z
	Started string `json:"started,omitempty"`

	// Status of the job
	// Example: Waiting
	// Required: true
	// Enum: [Waiting Running Succeeded Stopping Stopped Failed]
	Status *string `json:"status"`

	// TimeLimitSeconds How long the job supposed to run at maximum
	// Example: 3600
	TimeLimitSeconds int64 `json:"timeLimitSeconds,omitempty"`

	// node
	Node *Node `json:"node,omitempty"`

	// resources
	Resources *ResourceRequirements `json:"resources,omitempty"`
}

// Validate validates this scheduled job summary
func (m *ScheduledJobSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackoffLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
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

func (m *ScheduledJobSummary) validateBackoffLimit(formats strfmt.Registry) error {

	if err := validate.Required("backoffLimit", "body", m.BackoffLimit); err != nil {
		return err
	}

	return nil
}

func (m *ScheduledJobSummary) validateFailedCount(formats strfmt.Registry) error {

	if err := validate.Required("failedCount", "body", m.FailedCount); err != nil {
		return err
	}

	return nil
}

func (m *ScheduledJobSummary) validateReplicaList(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicaList) { // not required
		return nil
	}

	for i := 0; i < len(m.ReplicaList); i++ {
		if swag.IsZero(m.ReplicaList[i]) { // not required
			continue
		}

		if m.ReplicaList[i] != nil {
			if err := m.ReplicaList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicaList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicaList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var scheduledJobSummaryTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Waiting","Running","Succeeded","Stopping","Stopped","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduledJobSummaryTypeStatusPropEnum = append(scheduledJobSummaryTypeStatusPropEnum, v)
	}
}

const (

	// ScheduledJobSummaryStatusWaiting captures enum value "Waiting"
	ScheduledJobSummaryStatusWaiting string = "Waiting"

	// ScheduledJobSummaryStatusRunning captures enum value "Running"
	ScheduledJobSummaryStatusRunning string = "Running"

	// ScheduledJobSummaryStatusSucceeded captures enum value "Succeeded"
	ScheduledJobSummaryStatusSucceeded string = "Succeeded"

	// ScheduledJobSummaryStatusStopping captures enum value "Stopping"
	ScheduledJobSummaryStatusStopping string = "Stopping"

	// ScheduledJobSummaryStatusStopped captures enum value "Stopped"
	ScheduledJobSummaryStatusStopped string = "Stopped"

	// ScheduledJobSummaryStatusFailed captures enum value "Failed"
	ScheduledJobSummaryStatusFailed string = "Failed"
)

// prop value enum
func (m *ScheduledJobSummary) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduledJobSummaryTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScheduledJobSummary) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ScheduledJobSummary) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduledJobSummary) validateResources(formats strfmt.Registry) error {
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

// ContextValidate validate this scheduled job summary based on the context it is used
func (m *ScheduledJobSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReplicaList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
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

func (m *ScheduledJobSummary) contextValidateReplicaList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReplicaList); i++ {

		if m.ReplicaList[i] != nil {

			if swag.IsZero(m.ReplicaList[i]) { // not required
				return nil
			}

			if err := m.ReplicaList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicaList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicaList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScheduledJobSummary) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduledJobSummary) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ScheduledJobSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledJobSummary) UnmarshalBinary(b []byte) error {
	var res ScheduledJobSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}