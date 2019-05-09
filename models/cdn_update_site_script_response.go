// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CdnUpdateSiteScriptResponse The response from a request to update an EdgeEngine script
// swagger:model cdnUpdateSiteScriptResponse
type CdnUpdateSiteScriptResponse struct {

	// script
	Script *CdnSiteScript `json:"script,omitempty"`
}

// Validate validates this cdn update site script response
func (m *CdnUpdateSiteScriptResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnUpdateSiteScriptResponse) validateScript(formats strfmt.Registry) error {

	if swag.IsZero(m.Script) { // not required
		return nil
	}

	if m.Script != nil {
		if err := m.Script.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnUpdateSiteScriptResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnUpdateSiteScriptResponse) UnmarshalBinary(b []byte) error {
	var res CdnUpdateSiteScriptResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
