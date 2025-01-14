// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReplicationPolicy replication policy
//
// swagger:model ReplicationPolicy
type ReplicationPolicy struct {

	// The create time of the policy.
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty" js:"creationTime"`

	// Deprecated, use "replicate_deletion" instead. Whether to replicate the deletion operation.
	Deletion bool `json:"deletion,omitempty" js:"deletion"`

	// The description of the policy.
	Description string `json:"description,omitempty" js:"description"`

	// The destination namespace.
	DestNamespace string `json:"dest_namespace,omitempty" js:"destNamespace"`

	// Specify how many path components will be replaced by the provided destination namespace.
	// The default value is -1 in which case the legacy mode will be applied.
	DestNamespaceReplaceCount *int8 `json:"dest_namespace_replace_count,omitempty" js:"destNamespaceReplaceCount"`

	// The destination registry.
	DestRegistry *Registry `json:"dest_registry,omitempty" js:"destRegistry"`

	// Whether the policy is enabled or not.
	Enabled bool `json:"enabled,omitempty" js:"enabled"`

	// The replication policy filter array.
	Filters []*ReplicationFilter `json:"filters" js:"filters"`

	// The policy ID.
	ID int64 `json:"id,omitempty" js:"id"`

	// The policy name.
	Name string `json:"name,omitempty" js:"name"`

	// Whether to override the resources on the destination registry.
	Override bool `json:"override,omitempty" js:"override"`

	// Whether to replicate the deletion operation.
	ReplicateDeletion bool `json:"replicate_deletion,omitempty" js:"replicateDeletion"`

	// The source registry.
	SrcRegistry *Registry `json:"src_registry,omitempty" js:"srcRegistry"`

	// trigger
	Trigger *ReplicationTrigger `json:"trigger,omitempty" js:"trigger"`

	// The update time of the policy.
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty" js:"updateTime"`
}

// Validate validates this replication policy
func (m *ReplicationPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicationPolicy) validateCreationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReplicationPolicy) validateDestRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.DestRegistry) { // not required
		return nil
	}

	if m.DestRegistry != nil {
		if err := m.DestRegistry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dest_registry")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationPolicy) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReplicationPolicy) validateSrcRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.SrcRegistry) { // not required
		return nil
	}

	if m.SrcRegistry != nil {
		if err := m.SrcRegistry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("src_registry")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationPolicy) validateTrigger(formats strfmt.Registry) error {
	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if m.Trigger != nil {
		if err := m.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationPolicy) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this replication policy based on the context it is used
func (m *ReplicationPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSrcRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrigger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicationPolicy) contextValidateDestRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.DestRegistry != nil {
		if err := m.DestRegistry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dest_registry")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationPolicy) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReplicationPolicy) contextValidateSrcRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.SrcRegistry != nil {
		if err := m.SrcRegistry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("src_registry")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationPolicy) contextValidateTrigger(ctx context.Context, formats strfmt.Registry) error {

	if m.Trigger != nil {
		if err := m.Trigger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationPolicy) UnmarshalBinary(b []byte) error {
	var res ReplicationPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
