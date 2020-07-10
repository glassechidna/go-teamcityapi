// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewSetParameterTypeParams creates a new SetParameterTypeParams object
// with the default values initialized.
func NewSetParameterTypeParams() *SetParameterTypeParams {
	var ()
	return &SetParameterTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetParameterTypeParamsWithTimeout creates a new SetParameterTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetParameterTypeParamsWithTimeout(timeout time.Duration) *SetParameterTypeParams {
	var ()
	return &SetParameterTypeParams{

		timeout: timeout,
	}
}

// NewSetParameterTypeParamsWithContext creates a new SetParameterTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetParameterTypeParamsWithContext(ctx context.Context) *SetParameterTypeParams {
	var ()
	return &SetParameterTypeParams{

		Context: ctx,
	}
}

// NewSetParameterTypeParamsWithHTTPClient creates a new SetParameterTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetParameterTypeParamsWithHTTPClient(client *http.Client) *SetParameterTypeParams {
	var ()
	return &SetParameterTypeParams{
		HTTPClient: client,
	}
}

/*SetParameterTypeParams contains all the parameters to send to the API endpoint
for the set parameter type operation typically these are written to a http.Request
*/
type SetParameterTypeParams struct {

	/*Body*/
	Body *models.Type
	/*Name*/
	Name string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set parameter type params
func (o *SetParameterTypeParams) WithTimeout(timeout time.Duration) *SetParameterTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set parameter type params
func (o *SetParameterTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set parameter type params
func (o *SetParameterTypeParams) WithContext(ctx context.Context) *SetParameterTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set parameter type params
func (o *SetParameterTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set parameter type params
func (o *SetParameterTypeParams) WithHTTPClient(client *http.Client) *SetParameterTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set parameter type params
func (o *SetParameterTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set parameter type params
func (o *SetParameterTypeParams) WithBody(body *models.Type) *SetParameterTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set parameter type params
func (o *SetParameterTypeParams) SetBody(body *models.Type) {
	o.Body = body
}

// WithName adds the name to the set parameter type params
func (o *SetParameterTypeParams) WithName(name string) *SetParameterTypeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the set parameter type params
func (o *SetParameterTypeParams) SetName(name string) {
	o.Name = name
}

// WithProjectLocator adds the projectLocator to the set parameter type params
func (o *SetParameterTypeParams) WithProjectLocator(projectLocator string) *SetParameterTypeParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the set parameter type params
func (o *SetParameterTypeParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetParameterTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}