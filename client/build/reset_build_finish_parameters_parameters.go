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

// NewResetBuildFinishParametersParams creates a new ResetBuildFinishParametersParams object
// with the default values initialized.
func NewResetBuildFinishParametersParams() *ResetBuildFinishParametersParams {
	var ()
	return &ResetBuildFinishParametersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResetBuildFinishParametersParamsWithTimeout creates a new ResetBuildFinishParametersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResetBuildFinishParametersParamsWithTimeout(timeout time.Duration) *ResetBuildFinishParametersParams {
	var ()
	return &ResetBuildFinishParametersParams{

		timeout: timeout,
	}
}

// NewResetBuildFinishParametersParamsWithContext creates a new ResetBuildFinishParametersParams object
// with the default values initialized, and the ability to set a context for a request
func NewResetBuildFinishParametersParamsWithContext(ctx context.Context) *ResetBuildFinishParametersParams {
	var ()
	return &ResetBuildFinishParametersParams{

		Context: ctx,
	}
}

// NewResetBuildFinishParametersParamsWithHTTPClient creates a new ResetBuildFinishParametersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResetBuildFinishParametersParamsWithHTTPClient(client *http.Client) *ResetBuildFinishParametersParams {
	var ()
	return &ResetBuildFinishParametersParams{
		HTTPClient: client,
	}
}

/*ResetBuildFinishParametersParams contains all the parameters to send to the API endpoint
for the reset build finish parameters operation typically these are written to a http.Request
*/
type ResetBuildFinishParametersParams struct {

	/*BuildLocator*/
	BuildLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reset build finish parameters params
func (o *ResetBuildFinishParametersParams) WithTimeout(timeout time.Duration) *ResetBuildFinishParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset build finish parameters params
func (o *ResetBuildFinishParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset build finish parameters params
func (o *ResetBuildFinishParametersParams) WithContext(ctx context.Context) *ResetBuildFinishParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset build finish parameters params
func (o *ResetBuildFinishParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset build finish parameters params
func (o *ResetBuildFinishParametersParams) WithHTTPClient(client *http.Client) *ResetBuildFinishParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset build finish parameters params
func (o *ResetBuildFinishParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the reset build finish parameters params
func (o *ResetBuildFinishParametersParams) WithBuildLocator(buildLocator string) *ResetBuildFinishParametersParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the reset build finish parameters params
func (o *ResetBuildFinishParametersParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ResetBuildFinishParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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