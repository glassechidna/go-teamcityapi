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

// NewServeBuildStatisticValueParams creates a new ServeBuildStatisticValueParams object
// with the default values initialized.
func NewServeBuildStatisticValueParams() *ServeBuildStatisticValueParams {
	var ()
	return &ServeBuildStatisticValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildStatisticValueParamsWithTimeout creates a new ServeBuildStatisticValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildStatisticValueParamsWithTimeout(timeout time.Duration) *ServeBuildStatisticValueParams {
	var ()
	return &ServeBuildStatisticValueParams{

		timeout: timeout,
	}
}

// NewServeBuildStatisticValueParamsWithContext creates a new ServeBuildStatisticValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildStatisticValueParamsWithContext(ctx context.Context) *ServeBuildStatisticValueParams {
	var ()
	return &ServeBuildStatisticValueParams{

		Context: ctx,
	}
}

// NewServeBuildStatisticValueParamsWithHTTPClient creates a new ServeBuildStatisticValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildStatisticValueParamsWithHTTPClient(client *http.Client) *ServeBuildStatisticValueParams {
	var ()
	return &ServeBuildStatisticValueParams{
		HTTPClient: client,
	}
}

/*ServeBuildStatisticValueParams contains all the parameters to send to the API endpoint
for the serve build statistic value operation typically these are written to a http.Request
*/
type ServeBuildStatisticValueParams struct {

	/*BuildLocator*/
	BuildLocator string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) WithTimeout(timeout time.Duration) *ServeBuildStatisticValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) WithContext(ctx context.Context) *ServeBuildStatisticValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) WithHTTPClient(client *http.Client) *ServeBuildStatisticValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) WithBuildLocator(buildLocator string) *ServeBuildStatisticValueParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithName adds the name to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) WithName(name string) *ServeBuildStatisticValueParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the serve build statistic value params
func (o *ServeBuildStatisticValueParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildStatisticValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}