// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RelatedEntity related entity
// swagger:model relatedEntity
type RelatedEntity struct {

	// build
	Build *Build `json:"build,omitempty"`
}

// Validate validates this related entity
func (m *RelatedEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedEntity) validateBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelatedEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelatedEntity) UnmarshalBinary(b []byte) error {
	var res RelatedEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}