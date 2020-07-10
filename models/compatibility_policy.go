// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CompatibilityPolicy compatibility policy
// swagger:model compatibilityPolicy
type CompatibilityPolicy struct {

	// build types
	BuildTypes *BuildTypes `json:"buildTypes,omitempty"`

	// policy
	Policy string `json:"policy,omitempty" xml:"policy"`
}

// Validate validates this compatibility policy
func (m *CompatibilityPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompatibilityPolicy) validateBuildTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildTypes) { // not required
		return nil
	}

	if m.BuildTypes != nil {
		if err := m.BuildTypes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildTypes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompatibilityPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompatibilityPolicy) UnmarshalBinary(b []byte) error {
	var res CompatibilityPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
