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

// CdnSiteFeature The features available to CDN sites
//
// Multiple products can be served on a single CDN site. Features control how those products are managed on the StackPath backend.
//
//  - UNKNOWN: StackPath is unable to determine a site's feature
//  - CDN: A site has CDN caching abilities
//  - WAF: A site is protected by the StackPath Web Application Firewall
// swagger:model cdnSiteFeature
type CdnSiteFeature string

const (

	// CdnSiteFeatureUNKNOWN captures enum value "UNKNOWN"
	CdnSiteFeatureUNKNOWN CdnSiteFeature = "UNKNOWN"

	// CdnSiteFeatureCDN captures enum value "CDN"
	CdnSiteFeatureCDN CdnSiteFeature = "CDN"

	// CdnSiteFeatureWAF captures enum value "WAF"
	CdnSiteFeatureWAF CdnSiteFeature = "WAF"
)

// for schema
var cdnSiteFeatureEnum []interface{}

func init() {
	var res []CdnSiteFeature
	if err := json.Unmarshal([]byte(`["UNKNOWN","CDN","WAF"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cdnSiteFeatureEnum = append(cdnSiteFeatureEnum, v)
	}
}

func (m CdnSiteFeature) validateCdnSiteFeatureEnum(path, location string, value CdnSiteFeature) error {
	if err := validate.Enum(path, location, value, cdnSiteFeatureEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this cdn site feature
func (m CdnSiteFeature) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCdnSiteFeatureEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
