// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewGetSnapshotDepsParams creates a new GetSnapshotDepsParams object
// with the default values initialized.
func NewGetSnapshotDepsParams() *GetSnapshotDepsParams {
	var ()
	return &GetSnapshotDepsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotDepsParamsWithTimeout creates a new GetSnapshotDepsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnapshotDepsParamsWithTimeout(timeout time.Duration) *GetSnapshotDepsParams {
	var ()
	return &GetSnapshotDepsParams{

		timeout: timeout,
	}
}

// NewGetSnapshotDepsParamsWithContext creates a new GetSnapshotDepsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnapshotDepsParamsWithContext(ctx context.Context) *GetSnapshotDepsParams {
	var ()
	return &GetSnapshotDepsParams{

		Context: ctx,
	}
}

// NewGetSnapshotDepsParamsWithHTTPClient creates a new GetSnapshotDepsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnapshotDepsParamsWithHTTPClient(client *http.Client) *GetSnapshotDepsParams {
	var ()
	return &GetSnapshotDepsParams{
		HTTPClient: client,
	}
}

/*GetSnapshotDepsParams contains all the parameters to send to the API endpoint
for the get snapshot deps operation typically these are written to a http.Request
*/
type GetSnapshotDepsParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snapshot deps params
func (o *GetSnapshotDepsParams) WithTimeout(timeout time.Duration) *GetSnapshotDepsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot deps params
func (o *GetSnapshotDepsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot deps params
func (o *GetSnapshotDepsParams) WithContext(ctx context.Context) *GetSnapshotDepsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot deps params
func (o *GetSnapshotDepsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot deps params
func (o *GetSnapshotDepsParams) WithHTTPClient(client *http.Client) *GetSnapshotDepsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot deps params
func (o *GetSnapshotDepsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get snapshot deps params
func (o *GetSnapshotDepsParams) WithBtLocator(btLocator string) *GetSnapshotDepsParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get snapshot deps params
func (o *GetSnapshotDepsParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get snapshot deps params
func (o *GetSnapshotDepsParams) WithFields(fields *string) *GetSnapshotDepsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get snapshot deps params
func (o *GetSnapshotDepsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotDepsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
