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

// CdnGetCertificateVerificationDetailsResponse cdn get certificate verification details response
// swagger:model cdnGetCertificateVerificationDetailsResponse
type CdnGetCertificateVerificationDetailsResponse struct {

	// manual verification required
	ManualVerificationRequired bool `json:"manualVerificationRequired,omitempty"`

	// verification requirements
	VerificationRequirements []*CdnVerificationRequirements `json:"verificationRequirements"`
}

// Validate validates this cdn get certificate verification details response
func (m *CdnGetCertificateVerificationDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerificationRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnGetCertificateVerificationDetailsResponse) validateVerificationRequirements(formats strfmt.Registry) error {

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
func (m *CdnGetCertificateVerificationDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnGetCertificateVerificationDetailsResponse) UnmarshalBinary(b []byte) error {
	var res CdnGetCertificateVerificationDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}