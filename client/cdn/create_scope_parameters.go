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

// NewCreateScopeParams creates a new CreateScopeParams object
// with the default values initialized.
func NewCreateScopeParams() *CreateScopeParams {
	var ()
	return &CreateScopeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScopeParamsWithTimeout creates a new CreateScopeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateScopeParamsWithTimeout(timeout time.Duration) *CreateScopeParams {
	var ()
	return &CreateScopeParams{

		timeout: timeout,
	}
}

// NewCreateScopeParamsWithContext creates a new CreateScopeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateScopeParamsWithContext(ctx context.Context) *CreateScopeParams {
	var ()
	return &CreateScopeParams{

		Context: ctx,
	}
}

// NewCreateScopeParamsWithHTTPClient creates a new CreateScopeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateScopeParamsWithHTTPClient(client *http.Client) *CreateScopeParams {
	var ()
	return &CreateScopeParams{
		HTTPClient: client,
	}
}

/*CreateScopeParams contains all the parameters to send to the API endpoint
for the create scope operation typically these are written to a http.Request
*/
type CreateScopeParams struct {

	/*Body*/
	Body *models.CdnCreateScopeRequest
	/*SiteID
	  The ID of the site to create a scope on

	*/
	SiteID string
	/*StackID
	  The ID of the stack containing the site to create a scope on

	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create scope params
func (o *CreateScopeParams) WithTimeout(timeout time.Duration) *CreateScopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create scope params
func (o *CreateScopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create scope params
func (o *CreateScopeParams) WithContext(ctx context.Context) *CreateScopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create scope params
func (o *CreateScopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create scope params
func (o *CreateScopeParams) WithHTTPClient(client *http.Client) *CreateScopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create scope params
func (o *CreateScopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create scope params
func (o *CreateScopeParams) WithBody(body *models.CdnCreateScopeRequest) *CreateScopeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create scope params
func (o *CreateScopeParams) SetBody(body *models.CdnCreateScopeRequest) {
	o.Body = body
}

// WithSiteID adds the siteID to the create scope params
func (o *CreateScopeParams) WithSiteID(siteID string) *CreateScopeParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the create scope params
func (o *CreateScopeParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the create scope params
func (o *CreateScopeParams) WithStackID(stackID string) *CreateScopeParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the create scope params
func (o *CreateScopeParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
