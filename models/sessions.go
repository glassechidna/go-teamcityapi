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

// Sessions sessions
// swagger:model sessions
type Sessions struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// max active
	MaxActive int32 `json:"maxActive,omitempty" xml:"maxActive"`

	// session
	Session []*Session `json:"session"`

	// session counter
	SessionCounter int32 `json:"sessionCounter,omitempty" xml:"sessionCounter"`

	// session create rate
	SessionCreateRate int32 `json:"sessionCreateRate,omitempty" xml:"sessionCreateRate"`

	// session expire rate
	SessionExpireRate int32 `json:"sessionExpireRate,omitempty" xml:"sessionExpireRate"`

	// session max alive time
	SessionMaxAliveTime int32 `json:"sessionMaxAliveTime,omitempty" xml:"sessionMaxAliveTime"`
}

// Validate validates this sessions
func (m *Sessions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSession(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Sessions) validateSession(formats strfmt.Registry) error {

	if swag.IsZero(m.Session) { // not required
		return nil
	}

	for i := 0; i < len(m.Session); i++ {
		if swag.IsZero(m.Session[i]) { // not required
			continue
		}

		if m.Session[i] != nil {
			if err := m.Session[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("session" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Sessions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sessions) UnmarshalBinary(b []byte) error {
	var res Sessions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}