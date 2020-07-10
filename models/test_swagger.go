// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Test test
// swagger:model test
type Test struct {

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// investigations
	Investigations *Investigations `json:"investigations,omitempty"`

	// locator
	Locator string `json:"locator,omitempty" xml:"locator"`

	// mutes
	Mutes *Mutes `json:"mutes,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// test occurrences
	TestOccurrences *TestOccurrences `json:"testOccurrences,omitempty"`
}

// Validate validates this test
func (m *Test) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvestigations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestOccurrences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Test) validateInvestigations(formats strfmt.Registry) error {

	if swag.IsZero(m.Investigations) { // not required
		return nil
	}

	if m.Investigations != nil {
		if err := m.Investigations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("investigations")
			}
			return err
		}
	}

	return nil
}

func (m *Test) validateMutes(formats strfmt.Registry) error {

	if swag.IsZero(m.Mutes) { // not required
		return nil
	}

	if m.Mutes != nil {
		if err := m.Mutes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mutes")
			}
			return err
		}
	}

	return nil
}

func (m *Test) validateTestOccurrences(formats strfmt.Registry) error {

	if swag.IsZero(m.TestOccurrences) { // not required
		return nil
	}

	if m.TestOccurrences != nil {
		if err := m.TestOccurrences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testOccurrences")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Test) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Test) UnmarshalBinary(b []byte) error {
	var res Test
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
