// Code generated by go-swagger; DO NOT EDIT.

package change

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

// NewGetChangeIssueParams creates a new GetChangeIssueParams object
// with the default values initialized.
func NewGetChangeIssueParams() *GetChangeIssueParams {
	var ()
	return &GetChangeIssueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChangeIssueParamsWithTimeout creates a new GetChangeIssueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChangeIssueParamsWithTimeout(timeout time.Duration) *GetChangeIssueParams {
	var ()
	return &GetChangeIssueParams{

		timeout: timeout,
	}
}

// NewGetChangeIssueParamsWithContext creates a new GetChangeIssueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChangeIssueParamsWithContext(ctx context.Context) *GetChangeIssueParams {
	var ()
	return &GetChangeIssueParams{

		Context: ctx,
	}
}

// NewGetChangeIssueParamsWithHTTPClient creates a new GetChangeIssueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChangeIssueParamsWithHTTPClient(client *http.Client) *GetChangeIssueParams {
	var ()
	return &GetChangeIssueParams{
		HTTPClient: client,
	}
}

/*GetChangeIssueParams contains all the parameters to send to the API endpoint
for the get change issue operation typically these are written to a http.Request
*/
type GetChangeIssueParams struct {

	/*ChangeLocator*/
	ChangeLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get change issue params
func (o *GetChangeIssueParams) WithTimeout(timeout time.Duration) *GetChangeIssueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get change issue params
func (o *GetChangeIssueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get change issue params
func (o *GetChangeIssueParams) WithContext(ctx context.Context) *GetChangeIssueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get change issue params
func (o *GetChangeIssueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get change issue params
func (o *GetChangeIssueParams) WithHTTPClient(client *http.Client) *GetChangeIssueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get change issue params
func (o *GetChangeIssueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeLocator adds the changeLocator to the get change issue params
func (o *GetChangeIssueParams) WithChangeLocator(changeLocator string) *GetChangeIssueParams {
	o.SetChangeLocator(changeLocator)
	return o
}

// SetChangeLocator adds the changeLocator to the get change issue params
func (o *GetChangeIssueParams) SetChangeLocator(changeLocator string) {
	o.ChangeLocator = changeLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetChangeIssueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param changeLocator
	if err := r.SetPathParam("changeLocator", o.ChangeLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
