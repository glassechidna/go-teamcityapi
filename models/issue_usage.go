// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IssueUsage issue usage
// swagger:model IssueUsage
type IssueUsage struct {

	// changes
	Changes *Changes `json:"changes,omitempty"`

	// issue
	Issue *Issue `json:"issue,omitempty"`
}

// Validate validates this issue usage
func (m *IssueUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueUsage) validateChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	if m.Changes != nil {
		if err := m.Changes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changes")
			}
			return err
		}
	}

	return nil
}

func (m *IssueUsage) validateIssue(formats strfmt.Registry) error {

	if swag.IsZero(m.Issue) { // not required
		return nil
	}

	if m.Issue != nil {
		if err := m.Issue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueUsage) UnmarshalBinary(b []byte) error {
	var res IssueUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
