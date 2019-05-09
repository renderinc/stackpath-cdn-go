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

// GetMetricsRequestMetricsGranularity The time intervals that metrics can be rounded to
//
// - AUTO: StackPath will choose the granularity based on start and end dates
//  - P1M: Retrieve metrics per minute
//  - PT5M: Retrieve metrics per five minute intervals
//  - PT1H: Retrieve metrics per hour
//  - P1D: Retrieve metrics per day
// swagger:model GetMetricsRequestMetricsGranularity
type GetMetricsRequestMetricsGranularity string

const (

	// GetMetricsRequestMetricsGranularityAUTO captures enum value "AUTO"
	GetMetricsRequestMetricsGranularityAUTO GetMetricsRequestMetricsGranularity = "AUTO"

	// GetMetricsRequestMetricsGranularityP1M captures enum value "P1M"
	GetMetricsRequestMetricsGranularityP1M GetMetricsRequestMetricsGranularity = "P1M"

	// GetMetricsRequestMetricsGranularityPT5M captures enum value "PT5M"
	GetMetricsRequestMetricsGranularityPT5M GetMetricsRequestMetricsGranularity = "PT5M"

	// GetMetricsRequestMetricsGranularityPT1H captures enum value "PT1H"
	GetMetricsRequestMetricsGranularityPT1H GetMetricsRequestMetricsGranularity = "PT1H"

	// GetMetricsRequestMetricsGranularityP1D captures enum value "P1D"
	GetMetricsRequestMetricsGranularityP1D GetMetricsRequestMetricsGranularity = "P1D"
)

// for schema
var getMetricsRequestMetricsGranularityEnum []interface{}

func init() {
	var res []GetMetricsRequestMetricsGranularity
	if err := json.Unmarshal([]byte(`["AUTO","P1M","PT5M","PT1H","P1D"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getMetricsRequestMetricsGranularityEnum = append(getMetricsRequestMetricsGranularityEnum, v)
	}
}

func (m GetMetricsRequestMetricsGranularity) validateGetMetricsRequestMetricsGranularityEnum(path, location string, value GetMetricsRequestMetricsGranularity) error {
	if err := validate.Enum(path, location, value, getMetricsRequestMetricsGranularityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this get metrics request metrics granularity
func (m GetMetricsRequestMetricsGranularity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGetMetricsRequestMetricsGranularityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
