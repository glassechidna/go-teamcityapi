// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewGetBuildStatusTextParams creates a new GetBuildStatusTextParams object
// with the default values initialized.
func NewGetBuildStatusTextParams() *GetBuildStatusTextParams {
	var ()
	return &GetBuildStatusTextParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildStatusTextParamsWithTimeout creates a new GetBuildStatusTextParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBuildStatusTextParamsWithTimeout(timeout time.Duration) *GetBuildStatusTextParams {
	var ()
	return &GetBuildStatusTextParams{

		timeout: timeout,
	}
}

// NewGetBuildStatusTextParamsWithContext creates a new GetBuildStatusTextParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBuildStatusTextParamsWithContext(ctx context.Context) *GetBuildStatusTextParams {
	var ()
	return &GetBuildStatusTextParams{

		Context: ctx,
	}
}

// NewGetBuildStatusTextParamsWithHTTPClient creates a new GetBuildStatusTextParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBuildStatusTextParamsWithHTTPClient(client *http.Client) *GetBuildStatusTextParams {
	var ()
	return &GetBuildStatusTextParams{
		HTTPClient: client,
	}
}

/*GetBuildStatusTextParams contains all the parameters to send to the API endpoint
for the get build status text operation typically these are written to a http.Request
*/
type GetBuildStatusTextParams struct {

	/*BuildLocator*/
	BuildLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get build status text params
func (o *GetBuildStatusTextParams) WithTimeout(timeout time.Duration) *GetBuildStatusTextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build status text params
func (o *GetBuildStatusTextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build status text params
func (o *GetBuildStatusTextParams) WithContext(ctx context.Context) *GetBuildStatusTextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build status text params
func (o *GetBuildStatusTextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build status text params
func (o *GetBuildStatusTextParams) WithHTTPClient(client *http.Client) *GetBuildStatusTextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build status text params
func (o *GetBuildStatusTextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get build status text params
func (o *GetBuildStatusTextParams) WithBuildLocator(buildLocator string) *GetBuildStatusTextParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get build status text params
func (o *GetBuildStatusTextParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildStatusTextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
