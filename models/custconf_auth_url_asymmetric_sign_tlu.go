// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustconfAuthURLAsymmetricSignTlu custconf auth Url asymmetric sign tlu
// swagger:model custconfAuthUrlAsymmetricSignTlu
type CustconfAuthURLAsymmetricSignTlu struct {

	// algorithm Id map
	AlgorithmIDMap map[string]CustconfAuthURLAsymmetricSignTluAlgorithmIDMapEnumWrapperValue `json:"algorithmIdMap,omitempty"`

	// algorithm Id parameter name
	AlgorithmIDParameterName string `json:"algorithmIdParameterName,omitempty"`

	// digest parameter name
	DigestParameterName string `json:"digestParameterName,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// expire parameter name
	ExpireParameterName string `json:"expireParameterName,omitempty"`

	// String of values deliminated by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// key Id parameter name
	KeyIDParameterName string `json:"keyIdParameterName,omitempty"`

	// String of values deliminated by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// String of values deliminated by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// public key Id map
	PublicKeyIDMap map[string]string `json:"publicKeyIdMap,omitempty"`
}

// Validate validates this custconf auth Url asymmetric sign tlu
func (m *CustconfAuthURLAsymmetricSignTlu) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithmIDMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfAuthURLAsymmetricSignTlu) validateAlgorithmIDMap(formats strfmt.Registry) error {

	if swag.IsZero(m.AlgorithmIDMap) { // not required
		return nil
	}

	for k := range m.AlgorithmIDMap {

		if val, ok := m.AlgorithmIDMap[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAuthURLAsymmetricSignTlu) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAuthURLAsymmetricSignTlu) UnmarshalBinary(b []byte) error {
	var res CustconfAuthURLAsymmetricSignTlu
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}