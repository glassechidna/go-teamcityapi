// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OperationResult operation result
// swagger:model operationResult
type OperationResult struct {

	// message
	Message string `json:"message,omitempty"`

	// related
	Related *RelatedEntity `json:"related,omitempty"`
}

// Validate validates this operation result
func (m *OperationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationResult) validateRelated(formats strfmt.Registry) error {

	if swag.IsZero(m.Related) { // not required
		return nil
	}

	if m.Related != nil {
		if err := m.Related.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("related")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationResult) UnmarshalBinary(b []byte) error {
	var res OperationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
