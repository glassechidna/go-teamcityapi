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

// NewChangeStepSettingParams creates a new ChangeStepSettingParams object
// with the default values initialized.
func NewChangeStepSettingParams() *ChangeStepSettingParams {
	var ()
	return &ChangeStepSettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeStepSettingParamsWithTimeout creates a new ChangeStepSettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeStepSettingParamsWithTimeout(timeout time.Duration) *ChangeStepSettingParams {
	var ()
	return &ChangeStepSettingParams{

		timeout: timeout,
	}
}

// NewChangeStepSettingParamsWithContext creates a new ChangeStepSettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeStepSettingParamsWithContext(ctx context.Context) *ChangeStepSettingParams {
	var ()
	return &ChangeStepSettingParams{

		Context: ctx,
	}
}

// NewChangeStepSettingParamsWithHTTPClient creates a new ChangeStepSettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeStepSettingParamsWithHTTPClient(client *http.Client) *ChangeStepSettingParams {
	var ()
	return &ChangeStepSettingParams{
		HTTPClient: client,
	}
}

/*ChangeStepSettingParams contains all the parameters to send to the API endpoint
for the change step setting operation typically these are written to a http.Request
*/
type ChangeStepSettingParams struct {

	/*Body*/
	Body string
	/*BtLocator*/
	BtLocator string
	/*FieldName*/
	FieldName string
	/*StepID*/
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change step setting params
func (o *ChangeStepSettingParams) WithTimeout(timeout time.Duration) *ChangeStepSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change step setting params
func (o *ChangeStepSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change step setting params
func (o *ChangeStepSettingParams) WithContext(ctx context.Context) *ChangeStepSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change step setting params
func (o *ChangeStepSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change step setting params
func (o *ChangeStepSettingParams) WithHTTPClient(client *http.Client) *ChangeStepSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change step setting params
func (o *ChangeStepSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change step setting params
func (o *ChangeStepSettingParams) WithBody(body string) *ChangeStepSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change step setting params
func (o *ChangeStepSettingParams) SetBody(body string) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the change step setting params
func (o *ChangeStepSettingParams) WithBtLocator(btLocator string) *ChangeStepSettingParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the change step setting params
func (o *ChangeStepSettingParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFieldName adds the fieldName to the change step setting params
func (o *ChangeStepSettingParams) WithFieldName(fieldName string) *ChangeStepSettingParams {
	o.SetFieldName(fieldName)
	return o
}

// SetFieldName adds the fieldName to the change step setting params
func (o *ChangeStepSettingParams) SetFieldName(fieldName string) {
	o.FieldName = fieldName
}

// WithStepID adds the stepID to the change step setting params
func (o *ChangeStepSettingParams) WithStepID(stepID string) *ChangeStepSettingParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the change step setting params
func (o *ChangeStepSettingParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeStepSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param fieldName
	if err := r.SetPathParam("fieldName", o.FieldName); err != nil {
		return err
	}

	// path param stepId
	if err := r.SetPathParam("stepId", o.StepID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
