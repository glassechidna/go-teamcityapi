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

	models "github.com/glassechidna/go-teamcityapi/models"
)

// NewAddFeatureParams creates a new AddFeatureParams object
// with the default values initialized.
func NewAddFeatureParams() *AddFeatureParams {
	var ()
	return &AddFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddFeatureParamsWithTimeout creates a new AddFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddFeatureParamsWithTimeout(timeout time.Duration) *AddFeatureParams {
	var ()
	return &AddFeatureParams{

		timeout: timeout,
	}
}

// NewAddFeatureParamsWithContext creates a new AddFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddFeatureParamsWithContext(ctx context.Context) *AddFeatureParams {
	var ()
	return &AddFeatureParams{

		Context: ctx,
	}
}

// NewAddFeatureParamsWithHTTPClient creates a new AddFeatureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddFeatureParamsWithHTTPClient(client *http.Client) *AddFeatureParams {
	var ()
	return &AddFeatureParams{
		HTTPClient: client,
	}
}

/*AddFeatureParams contains all the parameters to send to the API endpoint
for the add feature operation typically these are written to a http.Request
*/
type AddFeatureParams struct {

	/*Body*/
	Body *models.Feature
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add feature params
func (o *AddFeatureParams) WithTimeout(timeout time.Duration) *AddFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add feature params
func (o *AddFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add feature params
func (o *AddFeatureParams) WithContext(ctx context.Context) *AddFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add feature params
func (o *AddFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add feature params
func (o *AddFeatureParams) WithHTTPClient(client *http.Client) *AddFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add feature params
func (o *AddFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add feature params
func (o *AddFeatureParams) WithBody(body *models.Feature) *AddFeatureParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add feature params
func (o *AddFeatureParams) SetBody(body *models.Feature) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the add feature params
func (o *AddFeatureParams) WithBtLocator(btLocator string) *AddFeatureParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the add feature params
func (o *AddFeatureParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the add feature params
func (o *AddFeatureParams) WithFields(fields *string) *AddFeatureParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add feature params
func (o *AddFeatureParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
