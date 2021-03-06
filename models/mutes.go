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

// Mutes mutes
// swagger:model mutes
type Mutes struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// default
	Default *bool `json:"default,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// mute
	Mute []*Mute `json:"mute"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref"`
}

// Validate validates this mutes
func (m *Mutes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mutes) validateMute(formats strfmt.Registry) error {

	if swag.IsZero(m.Mute) { // not required
		return nil
	}

	for i := 0; i < len(m.Mute); i++ {
		if swag.IsZero(m.Mute[i]) { // not required
			continue
		}

		if m.Mute[i] != nil {
			if err := m.Mute[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mute" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Mutes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mutes) UnmarshalBinary(b []byte) error {
	var res Mutes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
