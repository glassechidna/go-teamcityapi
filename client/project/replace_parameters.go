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

// NewReplaceParams creates a new ReplaceParams object
// with the default values initialized.
func NewReplaceParams() *ReplaceParams {
	var ()
	return &ReplaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceParamsWithTimeout creates a new ReplaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceParamsWithTimeout(timeout time.Duration) *ReplaceParams {
	var ()
	return &ReplaceParams{

		timeout: timeout,
	}
}

// NewReplaceParamsWithContext creates a new ReplaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceParamsWithContext(ctx context.Context) *ReplaceParams {
	var ()
	return &ReplaceParams{

		Context: ctx,
	}
}

// NewReplaceParamsWithHTTPClient creates a new ReplaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceParamsWithHTTPClient(client *http.Client) *ReplaceParams {
	var ()
	return &ReplaceParams{
		HTTPClient: client,
	}
}

/*ReplaceParams contains all the parameters to send to the API endpoint
for the replace operation typically these are written to a http.Request
*/
type ReplaceParams struct {

	/*Body*/
	Body *models.ProjectFeature
	/*FeatureLocator*/
	FeatureLocator string
	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace params
func (o *ReplaceParams) WithTimeout(timeout time.Duration) *ReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace params
func (o *ReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace params
func (o *ReplaceParams) WithContext(ctx context.Context) *ReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace params
func (o *ReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace params
func (o *ReplaceParams) WithHTTPClient(client *http.Client) *ReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace params
func (o *ReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace params
func (o *ReplaceParams) WithBody(body *models.ProjectFeature) *ReplaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace params
func (o *ReplaceParams) SetBody(body *models.ProjectFeature) {
	o.Body = body
}

// WithFeatureLocator adds the featureLocator to the replace params
func (o *ReplaceParams) WithFeatureLocator(featureLocator string) *ReplaceParams {
	o.SetFeatureLocator(featureLocator)
	return o
}

// SetFeatureLocator adds the featureLocator to the replace params
func (o *ReplaceParams) SetFeatureLocator(featureLocator string) {
	o.FeatureLocator = featureLocator
}

// WithFields adds the fields to the replace params
func (o *ReplaceParams) WithFields(fields *string) *ReplaceParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace params
func (o *ReplaceParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the replace params
func (o *ReplaceParams) WithProjectLocator(projectLocator string) *ReplaceParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the replace params
func (o *ReplaceParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param featureLocator
	if err := r.SetPathParam("featureLocator", o.FeatureLocator); err != nil {
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

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
