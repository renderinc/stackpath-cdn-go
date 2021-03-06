// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCDNIpsParams creates a new GetCDNIpsParams object
// with the default values initialized.
func NewGetCDNIpsParams() *GetCDNIpsParams {
	var (
		filterDefault       = string("ALL")
		responseTypeDefault = string("JSON")
	)
	return &GetCDNIpsParams{
		Filter:       &filterDefault,
		ResponseType: &responseTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCDNIpsParamsWithTimeout creates a new GetCDNIpsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCDNIpsParamsWithTimeout(timeout time.Duration) *GetCDNIpsParams {
	var (
		filterDefault       = string("ALL")
		responseTypeDefault = string("JSON")
	)
	return &GetCDNIpsParams{
		Filter:       &filterDefault,
		ResponseType: &responseTypeDefault,

		timeout: timeout,
	}
}

// NewGetCDNIpsParamsWithContext creates a new GetCDNIpsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCDNIpsParamsWithContext(ctx context.Context) *GetCDNIpsParams {
	var (
		filterDefault       = string("ALL")
		responseTypeDefault = string("JSON")
	)
	return &GetCDNIpsParams{
		Filter:       &filterDefault,
		ResponseType: &responseTypeDefault,

		Context: ctx,
	}
}

// NewGetCDNIpsParamsWithHTTPClient creates a new GetCDNIpsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCDNIpsParamsWithHTTPClient(client *http.Client) *GetCDNIpsParams {
	var (
		filterDefault       = string("ALL")
		responseTypeDefault = string("JSON")
	)
	return &GetCDNIpsParams{
		Filter:       &filterDefault,
		ResponseType: &responseTypeDefault,
		HTTPClient:   client,
	}
}

/*GetCDNIpsParams contains all the parameters to send to the API endpoint
for the get c d n ips operation typically these are written to a http.Request
*/
type GetCDNIpsParams struct {

	/*Filter*/
	Filter *string
	/*ResponseType*/
	ResponseType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get c d n ips params
func (o *GetCDNIpsParams) WithTimeout(timeout time.Duration) *GetCDNIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c d n ips params
func (o *GetCDNIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c d n ips params
func (o *GetCDNIpsParams) WithContext(ctx context.Context) *GetCDNIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c d n ips params
func (o *GetCDNIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c d n ips params
func (o *GetCDNIpsParams) WithHTTPClient(client *http.Client) *GetCDNIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c d n ips params
func (o *GetCDNIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get c d n ips params
func (o *GetCDNIpsParams) WithFilter(filter *string) *GetCDNIpsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get c d n ips params
func (o *GetCDNIpsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithResponseType adds the responseType to the get c d n ips params
func (o *GetCDNIpsParams) WithResponseType(responseType *string) *GetCDNIpsParams {
	o.SetResponseType(responseType)
	return o
}

// SetResponseType adds the responseType to the get c d n ips params
func (o *GetCDNIpsParams) SetResponseType(responseType *string) {
	o.ResponseType = responseType
}

// WriteToRequest writes these params to a swagger request
func (o *GetCDNIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.ResponseType != nil {

		// query param response_type
		var qrResponseType string
		if o.ResponseType != nil {
			qrResponseType = *o.ResponseType
		}
		qResponseType := qrResponseType
		if qResponseType != "" {
			if err := r.SetQueryParam("response_type", qResponseType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
