// Code generated by go-swagger; DO NOT EDIT.

package agent

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

// NewGetAllowedRunConfigurationsParams creates a new GetAllowedRunConfigurationsParams object
// with the default values initialized.
func NewGetAllowedRunConfigurationsParams() *GetAllowedRunConfigurationsParams {
	var ()
	return &GetAllowedRunConfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllowedRunConfigurationsParamsWithTimeout creates a new GetAllowedRunConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllowedRunConfigurationsParamsWithTimeout(timeout time.Duration) *GetAllowedRunConfigurationsParams {
	var ()
	return &GetAllowedRunConfigurationsParams{

		timeout: timeout,
	}
}

// NewGetAllowedRunConfigurationsParamsWithContext creates a new GetAllowedRunConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllowedRunConfigurationsParamsWithContext(ctx context.Context) *GetAllowedRunConfigurationsParams {
	var ()
	return &GetAllowedRunConfigurationsParams{

		Context: ctx,
	}
}

// NewGetAllowedRunConfigurationsParamsWithHTTPClient creates a new GetAllowedRunConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllowedRunConfigurationsParamsWithHTTPClient(client *http.Client) *GetAllowedRunConfigurationsParams {
	var ()
	return &GetAllowedRunConfigurationsParams{
		HTTPClient: client,
	}
}

/*GetAllowedRunConfigurationsParams contains all the parameters to send to the API endpoint
for the get allowed run configurations operation typically these are written to a http.Request
*/
type GetAllowedRunConfigurationsParams struct {

	/*AgentLocator*/
	AgentLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) WithTimeout(timeout time.Duration) *GetAllowedRunConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) WithContext(ctx context.Context) *GetAllowedRunConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) WithHTTPClient(client *http.Client) *GetAllowedRunConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentLocator adds the agentLocator to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) WithAgentLocator(agentLocator string) *GetAllowedRunConfigurationsParams {
	o.SetAgentLocator(agentLocator)
	return o
}

// SetAgentLocator adds the agentLocator to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) SetAgentLocator(agentLocator string) {
	o.AgentLocator = agentLocator
}

// WithFields adds the fields to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) WithFields(fields *string) *GetAllowedRunConfigurationsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get allowed run configurations params
func (o *GetAllowedRunConfigurationsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllowedRunConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentLocator
	if err := r.SetPathParam("agentLocator", o.AgentLocator); err != nil {
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
