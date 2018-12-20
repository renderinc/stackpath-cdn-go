// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CdnScopeOrigin cdn scope origin
// swagger:model cdnScopeOrigin
type CdnScopeOrigin struct {

	// origin
	Origin *SchemacdnOrigin `json:"origin,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`
}

// Validate validates this cdn scope origin
func (m *CdnScopeOrigin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnScopeOrigin) validateOrigin(formats strfmt.Registry) error {

	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnScopeOrigin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnScopeOrigin) UnmarshalBinary(b []byte) error {
	var res CdnScopeOrigin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}