// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewGetBackupStatusParams creates a new GetBackupStatusParams object
// with the default values initialized.
func NewGetBackupStatusParams() *GetBackupStatusParams {
	var ()
	return &GetBackupStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupStatusParamsWithTimeout creates a new GetBackupStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBackupStatusParamsWithTimeout(timeout time.Duration) *GetBackupStatusParams {
	var ()
	return &GetBackupStatusParams{

		timeout: timeout,
	}
}

// NewGetBackupStatusParamsWithContext creates a new GetBackupStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBackupStatusParamsWithContext(ctx context.Context) *GetBackupStatusParams {
	var ()
	return &GetBackupStatusParams{

		Context: ctx,
	}
}

// NewGetBackupStatusParamsWithHTTPClient creates a new GetBackupStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBackupStatusParamsWithHTTPClient(client *http.Client) *GetBackupStatusParams {
	var ()
	return &GetBackupStatusParams{
		HTTPClient: client,
	}
}

/*GetBackupStatusParams contains all the parameters to send to the API endpoint
for the get backup status operation typically these are written to a http.Request
*/
type GetBackupStatusParams struct {

	/*Body*/
	Body *models.BackupProcessManager

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get backup status params
func (o *GetBackupStatusParams) WithTimeout(timeout time.Duration) *GetBackupStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup status params
func (o *GetBackupStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup status params
func (o *GetBackupStatusParams) WithContext(ctx context.Context) *GetBackupStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup status params
func (o *GetBackupStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup status params
func (o *GetBackupStatusParams) WithHTTPClient(client *http.Client) *GetBackupStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup status params
func (o *GetBackupStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get backup status params
func (o *GetBackupStatusParams) WithBody(body *models.BackupProcessManager) *GetBackupStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get backup status params
func (o *GetBackupStatusParams) SetBody(body *models.BackupProcessManager) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
