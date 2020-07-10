// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewServeGroupParams creates a new ServeGroupParams object
// with the default values initialized.
func NewServeGroupParams() *ServeGroupParams {
	var ()
	return &ServeGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeGroupParamsWithTimeout creates a new ServeGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeGroupParamsWithTimeout(timeout time.Duration) *ServeGroupParams {
	var ()
	return &ServeGroupParams{

		timeout: timeout,
	}
}

// NewServeGroupParamsWithContext creates a new ServeGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeGroupParamsWithContext(ctx context.Context) *ServeGroupParams {
	var ()
	return &ServeGroupParams{

		Context: ctx,
	}
}

// NewServeGroupParamsWithHTTPClient creates a new ServeGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeGroupParamsWithHTTPClient(client *http.Client) *ServeGroupParams {
	var ()
	return &ServeGroupParams{
		HTTPClient: client,
	}
}

/*ServeGroupParams contains all the parameters to send to the API endpoint
for the serve group operation typically these are written to a http.Request
*/
type ServeGroupParams struct {

	/*Fields*/
	Fields *string
	/*GroupLocator*/
	GroupLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve group params
func (o *ServeGroupParams) WithTimeout(timeout time.Duration) *ServeGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve group params
func (o *ServeGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve group params
func (o *ServeGroupParams) WithContext(ctx context.Context) *ServeGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve group params
func (o *ServeGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve group params
func (o *ServeGroupParams) WithHTTPClient(client *http.Client) *ServeGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve group params
func (o *ServeGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve group params
func (o *ServeGroupParams) WithFields(fields *string) *ServeGroupParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve group params
func (o *ServeGroupParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGroupLocator adds the groupLocator to the serve group params
func (o *ServeGroupParams) WithGroupLocator(groupLocator string) *ServeGroupParams {
	o.SetGroupLocator(groupLocator)
	return o
}

// SetGroupLocator adds the groupLocator to the serve group params
func (o *ServeGroupParams) SetGroupLocator(groupLocator string) {
	o.GroupLocator = groupLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param groupLocator
	if err := r.SetPathParam("groupLocator", o.GroupLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
