// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewServeVersionParams creates a new ServeVersionParams object
// with the default values initialized.
func NewServeVersionParams() *ServeVersionParams {

	return &ServeVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeVersionParamsWithTimeout creates a new ServeVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeVersionParamsWithTimeout(timeout time.Duration) *ServeVersionParams {

	return &ServeVersionParams{

		timeout: timeout,
	}
}

// NewServeVersionParamsWithContext creates a new ServeVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeVersionParamsWithContext(ctx context.Context) *ServeVersionParams {

	return &ServeVersionParams{

		Context: ctx,
	}
}

// NewServeVersionParamsWithHTTPClient creates a new ServeVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeVersionParamsWithHTTPClient(client *http.Client) *ServeVersionParams {

	return &ServeVersionParams{
		HTTPClient: client,
	}
}

/*ServeVersionParams contains all the parameters to send to the API endpoint
for the serve version operation typically these are written to a http.Request
*/
type ServeVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve version params
func (o *ServeVersionParams) WithTimeout(timeout time.Duration) *ServeVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve version params
func (o *ServeVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve version params
func (o *ServeVersionParams) WithContext(ctx context.Context) *ServeVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve version params
func (o *ServeVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve version params
func (o *ServeVersionParams) WithHTTPClient(client *http.Client) *ServeVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve version params
func (o *ServeVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ServeVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
