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

	models "github.com/renderinc/stackpath-cdn-go/models"
)

// NewRequestCertificateParams creates a new RequestCertificateParams object
// with the default values initialized.
func NewRequestCertificateParams() *RequestCertificateParams {
	var ()
	return &RequestCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestCertificateParamsWithTimeout creates a new RequestCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestCertificateParamsWithTimeout(timeout time.Duration) *RequestCertificateParams {
	var ()
	return &RequestCertificateParams{

		timeout: timeout,
	}
}

// NewRequestCertificateParamsWithContext creates a new RequestCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestCertificateParamsWithContext(ctx context.Context) *RequestCertificateParams {
	var ()
	return &RequestCertificateParams{

		Context: ctx,
	}
}

// NewRequestCertificateParamsWithHTTPClient creates a new RequestCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestCertificateParamsWithHTTPClient(client *http.Client) *RequestCertificateParams {
	var ()
	return &RequestCertificateParams{
		HTTPClient: client,
	}
}

/*RequestCertificateParams contains all the parameters to send to the API endpoint
for the request certificate operation typically these are written to a http.Request
*/
type RequestCertificateParams struct {

	/*Body*/
	Body *models.CdnRequestCertificateRequest
	/*SiteID
	  The ID of the site to request an SSL certificate for

	*/
	SiteID string
	/*StackID
	  The ID of the stack containing the site to request an SSL certificate for

	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request certificate params
func (o *RequestCertificateParams) WithTimeout(timeout time.Duration) *RequestCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request certificate params
func (o *RequestCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request certificate params
func (o *RequestCertificateParams) WithContext(ctx context.Context) *RequestCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request certificate params
func (o *RequestCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request certificate params
func (o *RequestCertificateParams) WithHTTPClient(client *http.Client) *RequestCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request certificate params
func (o *RequestCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the request certificate params
func (o *RequestCertificateParams) WithBody(body *models.CdnRequestCertificateRequest) *RequestCertificateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the request certificate params
func (o *RequestCertificateParams) SetBody(body *models.CdnRequestCertificateRequest) {
	o.Body = body
}

// WithSiteID adds the siteID to the request certificate params
func (o *RequestCertificateParams) WithSiteID(siteID string) *RequestCertificateParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the request certificate params
func (o *RequestCertificateParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the request certificate params
func (o *RequestCertificateParams) WithStackID(stackID string) *RequestCertificateParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the request certificate params
func (o *RequestCertificateParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *RequestCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
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
