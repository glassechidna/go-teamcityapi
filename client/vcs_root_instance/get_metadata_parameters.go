// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

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

// NewGetMetadataParams creates a new GetMetadataParams object
// with the default values initialized.
func NewGetMetadataParams() *GetMetadataParams {
	var ()
	return &GetMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetadataParamsWithTimeout creates a new GetMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMetadataParamsWithTimeout(timeout time.Duration) *GetMetadataParams {
	var ()
	return &GetMetadataParams{

		timeout: timeout,
	}
}

// NewGetMetadataParamsWithContext creates a new GetMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMetadataParamsWithContext(ctx context.Context) *GetMetadataParams {
	var ()
	return &GetMetadataParams{

		Context: ctx,
	}
}

// NewGetMetadataParamsWithHTTPClient creates a new GetMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMetadataParamsWithHTTPClient(client *http.Client) *GetMetadataParams {
	var ()
	return &GetMetadataParams{
		HTTPClient: client,
	}
}

/*GetMetadataParams contains all the parameters to send to the API endpoint
for the get metadata operation typically these are written to a http.Request
*/
type GetMetadataParams struct {

	/*Fields*/
	Fields *string
	/*Path*/
	Path string
	/*VcsRootInstanceLocator*/
	VcsRootInstanceLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get metadata params
func (o *GetMetadataParams) WithTimeout(timeout time.Duration) *GetMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metadata params
func (o *GetMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metadata params
func (o *GetMetadataParams) WithContext(ctx context.Context) *GetMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metadata params
func (o *GetMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metadata params
func (o *GetMetadataParams) WithHTTPClient(client *http.Client) *GetMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metadata params
func (o *GetMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get metadata params
func (o *GetMetadataParams) WithFields(fields *string) *GetMetadataParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get metadata params
func (o *GetMetadataParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithPath adds the path to the get metadata params
func (o *GetMetadataParams) WithPath(path string) *GetMetadataParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get metadata params
func (o *GetMetadataParams) SetPath(path string) {
	o.Path = path
}

// WithVcsRootInstanceLocator adds the vcsRootInstanceLocator to the get metadata params
func (o *GetMetadataParams) WithVcsRootInstanceLocator(vcsRootInstanceLocator string) *GetMetadataParams {
	o.SetVcsRootInstanceLocator(vcsRootInstanceLocator)
	return o
}

// SetVcsRootInstanceLocator adds the vcsRootInstanceLocator to the get metadata params
func (o *GetMetadataParams) SetVcsRootInstanceLocator(vcsRootInstanceLocator string) {
	o.VcsRootInstanceLocator = vcsRootInstanceLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	// path param vcsRootInstanceLocator
	if err := r.SetPathParam("vcsRootInstanceLocator", o.VcsRootInstanceLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
