// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Changes changes
// swagger:model changes
type Changes struct {

	// change
	Change []*Change `json:"change"`

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref"`
}

// Validate validates this changes
func (m *Changes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Changes) validateChange(formats strfmt.Registry) error {

	if swag.IsZero(m.Change) { // not required
		return nil
	}

	for i := 0; i < len(m.Change); i++ {
		if swag.IsZero(m.Change[i]) { // not required
			continue
		}

		if m.Change[i] != nil {
			if err := m.Change[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("change" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Changes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Changes) UnmarshalBinary(b []byte) error {
	var res Changes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
