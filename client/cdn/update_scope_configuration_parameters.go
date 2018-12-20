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

	models "github.com/renderinc/stackpath-cdn-go/models"
)

// NewUpdateScopeConfigurationParams creates a new UpdateScopeConfigurationParams object
// with the default values initialized.
func NewUpdateScopeConfigurationParams() *UpdateScopeConfigurationParams {
	var ()
	return &UpdateScopeConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScopeConfigurationParamsWithTimeout creates a new UpdateScopeConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateScopeConfigurationParamsWithTimeout(timeout time.Duration) *UpdateScopeConfigurationParams {
	var ()
	return &UpdateScopeConfigurationParams{

		timeout: timeout,
	}
}

// NewUpdateScopeConfigurationParamsWithContext creates a new UpdateScopeConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateScopeConfigurationParamsWithContext(ctx context.Context) *UpdateScopeConfigurationParams {
	var ()
	return &UpdateScopeConfigurationParams{

		Context: ctx,
	}
}

// NewUpdateScopeConfigurationParamsWithHTTPClient creates a new UpdateScopeConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateScopeConfigurationParamsWithHTTPClient(client *http.Client) *UpdateScopeConfigurationParams {
	var ()
	return &UpdateScopeConfigurationParams{
		HTTPClient: client,
	}
}

/*UpdateScopeConfigurationParams contains all the parameters to send to the API endpoint
for the update scope configuration operation typically these are written to a http.Request
*/
type UpdateScopeConfigurationParams struct {

	/*Body*/
	Body *models.CdnUpdateScopeConfigurationRequest
	/*ScopeID*/
	ScopeID string
	/*SiteID*/
	SiteID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update scope configuration params
func (o *UpdateScopeConfigurationParams) WithTimeout(timeout time.Duration) *UpdateScopeConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update scope configuration params
func (o *UpdateScopeConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update scope configuration params
func (o *UpdateScopeConfigurationParams) WithContext(ctx context.Context) *UpdateScopeConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update scope configuration params
func (o *UpdateScopeConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update scope configuration params
func (o *UpdateScopeConfigurationParams) WithHTTPClient(client *http.Client) *UpdateScopeConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update scope configuration params
func (o *UpdateScopeConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update scope configuration params
func (o *UpdateScopeConfigurationParams) WithBody(body *models.CdnUpdateScopeConfigurationRequest) *UpdateScopeConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update scope configuration params
func (o *UpdateScopeConfigurationParams) SetBody(body *models.CdnUpdateScopeConfigurationRequest) {
	o.Body = body
}

// WithScopeID adds the scopeID to the update scope configuration params
func (o *UpdateScopeConfigurationParams) WithScopeID(scopeID string) *UpdateScopeConfigurationParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the update scope configuration params
func (o *UpdateScopeConfigurationParams) SetScopeID(scopeID string) {
	o.ScopeID = scopeID
}

// WithSiteID adds the siteID to the update scope configuration params
func (o *UpdateScopeConfigurationParams) WithSiteID(siteID string) *UpdateScopeConfigurationParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the update scope configuration params
func (o *UpdateScopeConfigurationParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the update scope configuration params
func (o *UpdateScopeConfigurationParams) WithStackID(stackID string) *UpdateScopeConfigurationParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the update scope configuration params
func (o *UpdateScopeConfigurationParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScopeConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scope_id
	if err := r.SetPathParam("scope_id", o.ScopeID); err != nil {
		return err
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