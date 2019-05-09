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

// CdnGetCertificatesResponse The response from a request to retrieve a stack's SSL certificates
// swagger:model cdnGetCertificatesResponse
type CdnGetCertificatesResponse struct {

	// page info
	PageInfo *PaginationPageInfo `json:"pageInfo,omitempty"`

	// The requested SSL certificates
	Results []*CdnCertificate `json:"results"`
}

// Validate validates this cdn get certificates response
func (m *CdnGetCertificatesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnGetCertificatesResponse) validatePageInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PageInfo) { // not required
		return nil
	}

	if m.PageInfo != nil {
		if err := m.PageInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pageInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CdnGetCertificatesResponse) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnGetCertificatesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnGetCertificatesResponse) UnmarshalBinary(b []byte) error {
	var res CdnGetCertificatesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
