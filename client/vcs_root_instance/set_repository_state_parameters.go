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

	models "github.com/glassechidna/go-teamcityapi/models"
)

// NewSetRepositoryStateParams creates a new SetRepositoryStateParams object
// with the default values initialized.
func NewSetRepositoryStateParams() *SetRepositoryStateParams {
	var ()
	return &SetRepositoryStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetRepositoryStateParamsWithTimeout creates a new SetRepositoryStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetRepositoryStateParamsWithTimeout(timeout time.Duration) *SetRepositoryStateParams {
	var ()
	return &SetRepositoryStateParams{

		timeout: timeout,
	}
}

// NewSetRepositoryStateParamsWithContext creates a new SetRepositoryStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetRepositoryStateParamsWithContext(ctx context.Context) *SetRepositoryStateParams {
	var ()
	return &SetRepositoryStateParams{

		Context: ctx,
	}
}

// NewSetRepositoryStateParamsWithHTTPClient creates a new SetRepositoryStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetRepositoryStateParamsWithHTTPClient(client *http.Client) *SetRepositoryStateParams {
	var ()
	return &SetRepositoryStateParams{
		HTTPClient: client,
	}
}

/*SetRepositoryStateParams contains all the parameters to send to the API endpoint
for the set repository state operation typically these are written to a http.Request
*/
type SetRepositoryStateParams struct {

	/*Body*/
	Body *models.Entries
	/*Fields*/
	Fields *string
	/*VcsRootInstanceLocator*/
	VcsRootInstanceLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set repository state params
func (o *SetRepositoryStateParams) WithTimeout(timeout time.Duration) *SetRepositoryStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set repository state params
func (o *SetRepositoryStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set repository state params
func (o *SetRepositoryStateParams) WithContext(ctx context.Context) *SetRepositoryStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set repository state params
func (o *SetRepositoryStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set repository state params
func (o *SetRepositoryStateParams) WithHTTPClient(client *http.Client) *SetRepositoryStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set repository state params
func (o *SetRepositoryStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set repository state params
func (o *SetRepositoryStateParams) WithBody(body *models.Entries) *SetRepositoryStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set repository state params
func (o *SetRepositoryStateParams) SetBody(body *models.Entries) {
	o.Body = body
}

// WithFields adds the fields to the set repository state params
func (o *SetRepositoryStateParams) WithFields(fields *string) *SetRepositoryStateParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the set repository state params
func (o *SetRepositoryStateParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithVcsRootInstanceLocator adds the vcsRootInstanceLocator to the set repository state params
func (o *SetRepositoryStateParams) WithVcsRootInstanceLocator(vcsRootInstanceLocator string) *SetRepositoryStateParams {
	o.SetVcsRootInstanceLocator(vcsRootInstanceLocator)
	return o
}

// SetVcsRootInstanceLocator adds the vcsRootInstanceLocator to the set repository state params
func (o *SetRepositoryStateParams) SetVcsRootInstanceLocator(vcsRootInstanceLocator string) {
	o.VcsRootInstanceLocator = vcsRootInstanceLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetRepositoryStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	// path param vcsRootInstanceLocator
	if err := r.SetPathParam("vcsRootInstanceLocator", o.VcsRootInstanceLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
