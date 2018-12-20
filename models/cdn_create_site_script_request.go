// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CdnCreateSiteScriptRequest cdn create site script request
// swagger:model cdnCreateSiteScriptRequest
type CdnCreateSiteScriptRequest struct {

	// code
	// Format: byte
	Code strfmt.Base64 `json:"code,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// paths
	Paths []string `json:"paths"`
}

// Validate validates this cdn create site script request
func (m *CdnCreateSiteScriptRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnCreateSiteScriptRequest) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *CdnCreateSiteScriptRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnCreateSiteScriptRequest) UnmarshalBinary(b []byte) error {
	var res CdnCreateSiteScriptRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}