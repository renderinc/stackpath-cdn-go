// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AuthACLAccessCodeEnumWrapperValue auth ACL access code enum wrapper value
// swagger:model AuthACLAccessCodeEnumWrapperValue
type AuthACLAccessCodeEnumWrapperValue string

const (

	// AuthACLAccessCodeEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	AuthACLAccessCodeEnumWrapperValueUNKNOWN AuthACLAccessCodeEnumWrapperValue = "UNKNOWN"

	// AuthACLAccessCodeEnumWrapperValueAllow captures enum value "allow"
	AuthACLAccessCodeEnumWrapperValueAllow AuthACLAccessCodeEnumWrapperValue = "allow"

	// AuthACLAccessCodeEnumWrapperValueDeny captures enum value "deny"
	AuthACLAccessCodeEnumWrapperValueDeny AuthACLAccessCodeEnumWrapperValue = "deny"
)

// for schema
var authAclAccessCodeEnumWrapperValueEnum []interface{}

func init() {
	var res []AuthACLAccessCodeEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authAclAccessCodeEnumWrapperValueEnum = append(authAclAccessCodeEnumWrapperValueEnum, v)
	}
}

func (m AuthACLAccessCodeEnumWrapperValue) validateAuthACLAccessCodeEnumWrapperValueEnum(path, location string, value AuthACLAccessCodeEnumWrapperValue) error {
	if err := validate.Enum(path, location, value, authAclAccessCodeEnumWrapperValueEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this auth ACL access code enum wrapper value
func (m AuthACLAccessCodeEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuthACLAccessCodeEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}