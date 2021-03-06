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

// NewDeleteAgentRequirementParams creates a new DeleteAgentRequirementParams object
// with the default values initialized.
func NewDeleteAgentRequirementParams() *DeleteAgentRequirementParams {
	var ()
	return &DeleteAgentRequirementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAgentRequirementParamsWithTimeout creates a new DeleteAgentRequirementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAgentRequirementParamsWithTimeout(timeout time.Duration) *DeleteAgentRequirementParams {
	var ()
	return &DeleteAgentRequirementParams{

		timeout: timeout,
	}
}

// NewDeleteAgentRequirementParamsWithContext creates a new DeleteAgentRequirementParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAgentRequirementParamsWithContext(ctx context.Context) *DeleteAgentRequirementParams {
	var ()
	return &DeleteAgentRequirementParams{

		Context: ctx,
	}
}

// NewDeleteAgentRequirementParamsWithHTTPClient creates a new DeleteAgentRequirementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAgentRequirementParamsWithHTTPClient(client *http.Client) *DeleteAgentRequirementParams {
	var ()
	return &DeleteAgentRequirementParams{
		HTTPClient: client,
	}
}

/*DeleteAgentRequirementParams contains all the parameters to send to the API endpoint
for the delete agent requirement operation typically these are written to a http.Request
*/
type DeleteAgentRequirementParams struct {

	/*AgentRequirementLocator*/
	AgentRequirementLocator string
	/*BtLocator*/
	BtLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete agent requirement params
func (o *DeleteAgentRequirementParams) WithTimeout(timeout time.Duration) *DeleteAgentRequirementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete agent requirement params
func (o *DeleteAgentRequirementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete agent requirement params
func (o *DeleteAgentRequirementParams) WithContext(ctx context.Context) *DeleteAgentRequirementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete agent requirement params
func (o *DeleteAgentRequirementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete agent requirement params
func (o *DeleteAgentRequirementParams) WithHTTPClient(client *http.Client) *DeleteAgentRequirementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete agent requirement params
func (o *DeleteAgentRequirementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentRequirementLocator adds the agentRequirementLocator to the delete agent requirement params
func (o *DeleteAgentRequirementParams) WithAgentRequirementLocator(agentRequirementLocator string) *DeleteAgentRequirementParams {
	o.SetAgentRequirementLocator(agentRequirementLocator)
	return o
}

// SetAgentRequirementLocator adds the agentRequirementLocator to the delete agent requirement params
func (o *DeleteAgentRequirementParams) SetAgentRequirementLocator(agentRequirementLocator string) {
	o.AgentRequirementLocator = agentRequirementLocator
}

// WithBtLocator adds the btLocator to the delete agent requirement params
func (o *DeleteAgentRequirementParams) WithBtLocator(btLocator string) *DeleteAgentRequirementParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the delete agent requirement params
func (o *DeleteAgentRequirementParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAgentRequirementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentRequirementLocator
	if err := r.SetPathParam("agentRequirementLocator", o.AgentRequirementLocator); err != nil {
		return err
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
