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

// CdnRequestCertificateResponse The response from a request to provision an SSL certificate for a CDN site
// swagger:model cdnRequestCertificateResponse
type CdnRequestCertificateResponse struct {

	// certificate
	Certificate *CdnCertificate `json:"certificate,omitempty"`

	// The certificate's verification requirements
	VerificationRequirements []*CdnVerificationRequirements `json:"verificationRequirements,omitempty"`
}

// Validate validates this cdn request certificate response
func (m *CdnRequestCertificateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerificationRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnRequestCertificateResponse) validateCertificate(formats strfmt.Registry) error {

	if swag.IsZero(m.Certificate) { // not required
		return nil
	}

	if m.Certificate != nil {
		if err := m.Certificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *CdnRequestCertificateResponse) validateVerificationRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.VerificationRequirements) { // not required
		return nil
	}

	for i := 0; i < len(m.VerificationRequirements); i++ {
		if swag.IsZero(m.VerificationRequirements[i]) { // not required
			continue
		}

		if m.VerificationRequirements[i] != nil {
			if err := m.VerificationRequirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("verificationRequirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnRequestCertificateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnRequestCertificateResponse) UnmarshalBinary(b []byte) error {
	var res CdnRequestCertificateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
