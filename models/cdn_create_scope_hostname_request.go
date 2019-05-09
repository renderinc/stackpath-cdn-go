// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CdnCreateScopeHostnameRequest A request to add a hostname to a CDN site scope
// swagger:model cdnCreateScopeHostnameRequest
type CdnCreateScopeHostnameRequest struct {

	// Whether or not to add the hostname to a CDN site's CDN scope or its WAF scope
	//
	// When true, this call adds the hostname to a CDN site's scope instead of loading from a CDN site's WAF scope, if the site has WAF service.
	DisableTransparentMode bool `json:"disableTransparentMode,omitempty"`

	// The hostname to add to a scope
	Domain string `json:"domain,omitempty"`
}

// Validate validates this cdn create scope hostname request
func (m *CdnCreateScopeHostnameRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CdnCreateScopeHostnameRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnCreateScopeHostnameRequest) UnmarshalBinary(b []byte) error {
	var res CdnCreateScopeHostnameRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
