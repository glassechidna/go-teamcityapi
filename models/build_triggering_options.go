// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BuildTriggeringOptions build triggering options
// swagger:model buildTriggeringOptions
type BuildTriggeringOptions struct {

	// clean sources
	CleanSources *bool `json:"cleanSources,omitempty" xml:"cleanSources"`

	// clean sources in all dependencies
	CleanSourcesInAllDependencies *bool `json:"cleanSourcesInAllDependencies,omitempty" xml:"cleanSourcesInAllDependencies"`

	// freeze settings
	FreezeSettings *bool `json:"freezeSettings,omitempty" xml:"freezeSettings"`

	// queue at top
	QueueAtTop *bool `json:"queueAtTop,omitempty" xml:"queueAtTop"`

	// rebuild all dependencies
	RebuildAllDependencies *bool `json:"rebuildAllDependencies,omitempty" xml:"rebuildAllDependencies"`

	// rebuild dependencies
	RebuildDependencies *BuildTypes `json:"rebuildDependencies,omitempty"`

	// tag dependencies
	TagDependencies *bool `json:"tagDependencies,omitempty" xml:"tagDependencies"`
}

// Validate validates this build triggering options
func (m *BuildTriggeringOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRebuildDependencies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildTriggeringOptions) validateRebuildDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.RebuildDependencies) { // not required
		return nil
	}

	if m.RebuildDependencies != nil {
		if err := m.RebuildDependencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rebuildDependencies")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildTriggeringOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildTriggeringOptions) UnmarshalBinary(b []byte) error {
	var res BuildTriggeringOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
