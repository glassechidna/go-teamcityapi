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

// Files files
// swagger:model files
type Files struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// file
	File []*File `json:"file"`

	// href
	Href string `json:"href,omitempty" xml:"href"`
}

// Validate validates this files
func (m *Files) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Files) validateFile(formats strfmt.Registry) error {

	if swag.IsZero(m.File) { // not required
		return nil
	}

	for i := 0; i < len(m.File); i++ {
		if swag.IsZero(m.File[i]) { // not required
			continue
		}

		if m.File[i] != nil {
			if err := m.File[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Files) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Files) UnmarshalBinary(b []byte) error {
	var res Files
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
