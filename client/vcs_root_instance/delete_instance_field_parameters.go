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

// NewDeleteInstanceFieldParams creates a new DeleteInstanceFieldParams object
// with the default values initialized.
func NewDeleteInstanceFieldParams() *DeleteInstanceFieldParams {
	var ()
	return &DeleteInstanceFieldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceFieldParamsWithTimeout creates a new DeleteInstanceFieldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstanceFieldParamsWithTimeout(timeout time.Duration) *DeleteInstanceFieldParams {
	var ()
	return &DeleteInstanceFieldParams{

		timeout: timeout,
	}
}

// NewDeleteInstanceFieldParamsWithContext creates a new DeleteInstanceFieldParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstanceFieldParamsWithContext(ctx context.Context) *DeleteInstanceFieldParams {
	var ()
	return &DeleteInstanceFieldParams{

		Context: ctx,
	}
}

// NewDeleteInstanceFieldParamsWithHTTPClient creates a new DeleteInstanceFieldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstanceFieldParamsWithHTTPClient(client *http.Client) *DeleteInstanceFieldParams {
	var ()
	return &DeleteInstanceFieldParams{
		HTTPClient: client,
	}
}

/*DeleteInstanceFieldParams contains all the parameters to send to the API endpoint
for the delete instance field operation typically these are written to a http.Request
*/
type DeleteInstanceFieldParams struct {

	/*Field*/
	Field string
	/*VcsRootInstanceLocator*/
	VcsRootInstanceLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instance field params
func (o *DeleteInstanceFieldParams) WithTimeout(timeout time.Duration) *DeleteInstanceFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance field params
func (o *DeleteInstanceFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance field params
func (o *DeleteInstanceFieldParams) WithContext(ctx context.Context) *DeleteInstanceFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance field params
func (o *DeleteInstanceFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance field params
func (o *DeleteInstanceFieldParams) WithHTTPClient(client *http.Client) *DeleteInstanceFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance field params
func (o *DeleteInstanceFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithField adds the field to the delete instance field params
func (o *DeleteInstanceFieldParams) WithField(field string) *DeleteInstanceFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the delete instance field params
func (o *DeleteInstanceFieldParams) SetField(field string) {
	o.Field = field
}

// WithVcsRootInstanceLocator adds the vcsRootInstanceLocator to the delete instance field params
func (o *DeleteInstanceFieldParams) WithVcsRootInstanceLocator(vcsRootInstanceLocator string) *DeleteInstanceFieldParams {
	o.SetVcsRootInstanceLocator(vcsRootInstanceLocator)
	return o
}

// SetVcsRootInstanceLocator adds the vcsRootInstanceLocator to the delete instance field params
func (o *DeleteInstanceFieldParams) SetVcsRootInstanceLocator(vcsRootInstanceLocator string) {
	o.VcsRootInstanceLocator = vcsRootInstanceLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
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
