// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCertificateSitesParams creates a new GetCertificateSitesParams object
// with the default values initialized.
func NewGetCertificateSitesParams() *GetCertificateSitesParams {
	var ()
	return &GetCertificateSitesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificateSitesParamsWithTimeout creates a new GetCertificateSitesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCertificateSitesParamsWithTimeout(timeout time.Duration) *GetCertificateSitesParams {
	var ()
	return &GetCertificateSitesParams{

		timeout: timeout,
	}
}

// NewGetCertificateSitesParamsWithContext creates a new GetCertificateSitesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCertificateSitesParamsWithContext(ctx context.Context) *GetCertificateSitesParams {
	var ()
	return &GetCertificateSitesParams{

		Context: ctx,
	}
}

// NewGetCertificateSitesParamsWithHTTPClient creates a new GetCertificateSitesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCertificateSitesParamsWithHTTPClient(client *http.Client) *GetCertificateSitesParams {
	var ()
	return &GetCertificateSitesParams{
		HTTPClient: client,
	}
}

/*GetCertificateSitesParams contains all the parameters to send to the API endpoint
for the get certificate sites operation typically these are written to a http.Request
*/
type GetCertificateSitesParams struct {

	/*CertificateID*/
	CertificateID string
	/*PageRequestAfter
	  after is the cursor value after which data will be returned.

	*/
	PageRequestAfter *string
	/*PageRequestFilter
	  filter will accept sql style constraints.

	*/
	PageRequestFilter *string
	/*PageRequestFirst
	  first is the number of items desired.

	*/
	PageRequestFirst *string
	/*PageRequestSortBy
	  sort_by will sort the response by the given field.

	*/
	PageRequestSortBy *string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get certificate sites params
func (o *GetCertificateSitesParams) WithTimeout(timeout time.Duration) *GetCertificateSitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificate sites params
func (o *GetCertificateSitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificate sites params
func (o *GetCertificateSitesParams) WithContext(ctx context.Context) *GetCertificateSitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificate sites params
func (o *GetCertificateSitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificate sites params
func (o *GetCertificateSitesParams) WithHTTPClient(client *http.Client) *GetCertificateSitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificate sites params
func (o *GetCertificateSitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateID adds the certificateID to the get certificate sites params
func (o *GetCertificateSitesParams) WithCertificateID(certificateID string) *GetCertificateSitesParams {
	o.SetCertificateID(certificateID)
	return o
}

// SetCertificateID adds the certificateId to the get certificate sites params
func (o *GetCertificateSitesParams) SetCertificateID(certificateID string) {
	o.CertificateID = certificateID
}

// WithPageRequestAfter adds the pageRequestAfter to the get certificate sites params
func (o *GetCertificateSitesParams) WithPageRequestAfter(pageRequestAfter *string) *GetCertificateSitesParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get certificate sites params
func (o *GetCertificateSitesParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get certificate sites params
func (o *GetCertificateSitesParams) WithPageRequestFilter(pageRequestFilter *string) *GetCertificateSitesParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get certificate sites params
func (o *GetCertificateSitesParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get certificate sites params
func (o *GetCertificateSitesParams) WithPageRequestFirst(pageRequestFirst *string) *GetCertificateSitesParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get certificate sites params
func (o *GetCertificateSitesParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get certificate sites params
func (o *GetCertificateSitesParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetCertificateSitesParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get certificate sites params
func (o *GetCertificateSitesParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithStackID adds the stackID to the get certificate sites params
func (o *GetCertificateSitesParams) WithStackID(stackID string) *GetCertificateSitesParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get certificate sites params
func (o *GetCertificateSitesParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificateSitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param certificate_id
	if err := r.SetPathParam("certificate_id", o.CertificateID); err != nil {
		return err
	}

	if o.PageRequestAfter != nil {

		// query param page_request.after
		var qrPageRequestAfter string
		if o.PageRequestAfter != nil {
			qrPageRequestAfter = *o.PageRequestAfter
		}
		qPageRequestAfter := qrPageRequestAfter
		if qPageRequestAfter != "" {
			if err := r.SetQueryParam("page_request.after", qPageRequestAfter); err != nil {
				return err
			}
		}

	}

	if o.PageRequestFilter != nil {

		// query param page_request.filter
		var qrPageRequestFilter string
		if o.PageRequestFilter != nil {
			qrPageRequestFilter = *o.PageRequestFilter
		}
		qPageRequestFilter := qrPageRequestFilter
		if qPageRequestFilter != "" {
			if err := r.SetQueryParam("page_request.filter", qPageRequestFilter); err != nil {
				return err
			}
		}

	}

	if o.PageRequestFirst != nil {

		// query param page_request.first
		var qrPageRequestFirst string
		if o.PageRequestFirst != nil {
			qrPageRequestFirst = *o.PageRequestFirst
		}
		qPageRequestFirst := qrPageRequestFirst
		if qPageRequestFirst != "" {
			if err := r.SetQueryParam("page_request.first", qPageRequestFirst); err != nil {
				return err
			}
		}

	}

	if o.PageRequestSortBy != nil {

		// query param page_request.sort_by
		var qrPageRequestSortBy string
		if o.PageRequestSortBy != nil {
			qrPageRequestSortBy = *o.PageRequestSortBy
		}
		qPageRequestSortBy := qrPageRequestSortBy
		if qPageRequestSortBy != "" {
			if err := r.SetQueryParam("page_request.sort_by", qPageRequestSortBy); err != nil {
				return err
			}
		}

	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}