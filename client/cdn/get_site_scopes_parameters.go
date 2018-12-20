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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSiteScopesParams creates a new GetSiteScopesParams object
// with the default values initialized.
func NewGetSiteScopesParams() *GetSiteScopesParams {
	var ()
	return &GetSiteScopesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSiteScopesParamsWithTimeout creates a new GetSiteScopesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSiteScopesParamsWithTimeout(timeout time.Duration) *GetSiteScopesParams {
	var ()
	return &GetSiteScopesParams{

		timeout: timeout,
	}
}

// NewGetSiteScopesParamsWithContext creates a new GetSiteScopesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSiteScopesParamsWithContext(ctx context.Context) *GetSiteScopesParams {
	var ()
	return &GetSiteScopesParams{

		Context: ctx,
	}
}

// NewGetSiteScopesParamsWithHTTPClient creates a new GetSiteScopesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSiteScopesParamsWithHTTPClient(client *http.Client) *GetSiteScopesParams {
	var ()
	return &GetSiteScopesParams{
		HTTPClient: client,
	}
}

/*GetSiteScopesParams contains all the parameters to send to the API endpoint
for the get site scopes operation typically these are written to a http.Request
*/
type GetSiteScopesParams struct {

	/*DisableTransparentMode*/
	DisableTransparentMode *bool
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
	/*SiteID*/
	SiteID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get site scopes params
func (o *GetSiteScopesParams) WithTimeout(timeout time.Duration) *GetSiteScopesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get site scopes params
func (o *GetSiteScopesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get site scopes params
func (o *GetSiteScopesParams) WithContext(ctx context.Context) *GetSiteScopesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get site scopes params
func (o *GetSiteScopesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get site scopes params
func (o *GetSiteScopesParams) WithHTTPClient(client *http.Client) *GetSiteScopesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get site scopes params
func (o *GetSiteScopesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisableTransparentMode adds the disableTransparentMode to the get site scopes params
func (o *GetSiteScopesParams) WithDisableTransparentMode(disableTransparentMode *bool) *GetSiteScopesParams {
	o.SetDisableTransparentMode(disableTransparentMode)
	return o
}

// SetDisableTransparentMode adds the disableTransparentMode to the get site scopes params
func (o *GetSiteScopesParams) SetDisableTransparentMode(disableTransparentMode *bool) {
	o.DisableTransparentMode = disableTransparentMode
}

// WithPageRequestAfter adds the pageRequestAfter to the get site scopes params
func (o *GetSiteScopesParams) WithPageRequestAfter(pageRequestAfter *string) *GetSiteScopesParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get site scopes params
func (o *GetSiteScopesParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get site scopes params
func (o *GetSiteScopesParams) WithPageRequestFilter(pageRequestFilter *string) *GetSiteScopesParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get site scopes params
func (o *GetSiteScopesParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get site scopes params
func (o *GetSiteScopesParams) WithPageRequestFirst(pageRequestFirst *string) *GetSiteScopesParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get site scopes params
func (o *GetSiteScopesParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get site scopes params
func (o *GetSiteScopesParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetSiteScopesParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get site scopes params
func (o *GetSiteScopesParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithSiteID adds the siteID to the get site scopes params
func (o *GetSiteScopesParams) WithSiteID(siteID string) *GetSiteScopesParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get site scopes params
func (o *GetSiteScopesParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the get site scopes params
func (o *GetSiteScopesParams) WithStackID(stackID string) *GetSiteScopesParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get site scopes params
func (o *GetSiteScopesParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteScopesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisableTransparentMode != nil {

		// query param disable_transparent_mode
		var qrDisableTransparentMode bool
		if o.DisableTransparentMode != nil {
			qrDisableTransparentMode = *o.DisableTransparentMode
		}
		qDisableTransparentMode := swag.FormatBool(qrDisableTransparentMode)
		if qDisableTransparentMode != "" {
			if err := r.SetQueryParam("disable_transparent_mode", qDisableTransparentMode); err != nil {
				return err
			}
		}

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