// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// NewGetCachedBuildPromotionsStatsParams creates a new GetCachedBuildPromotionsStatsParams object
// with the default values initialized.
func NewGetCachedBuildPromotionsStatsParams() *GetCachedBuildPromotionsStatsParams {
	var ()
	return &GetCachedBuildPromotionsStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCachedBuildPromotionsStatsParamsWithTimeout creates a new GetCachedBuildPromotionsStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCachedBuildPromotionsStatsParamsWithTimeout(timeout time.Duration) *GetCachedBuildPromotionsStatsParams {
	var ()
	return &GetCachedBuildPromotionsStatsParams{

		timeout: timeout,
	}
}

// NewGetCachedBuildPromotionsStatsParamsWithContext creates a new GetCachedBuildPromotionsStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCachedBuildPromotionsStatsParamsWithContext(ctx context.Context) *GetCachedBuildPromotionsStatsParams {
	var ()
	return &GetCachedBuildPromotionsStatsParams{

		Context: ctx,
	}
}

// NewGetCachedBuildPromotionsStatsParamsWithHTTPClient creates a new GetCachedBuildPromotionsStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCachedBuildPromotionsStatsParamsWithHTTPClient(client *http.Client) *GetCachedBuildPromotionsStatsParams {
	var ()
	return &GetCachedBuildPromotionsStatsParams{
		HTTPClient: client,
	}
}

/*GetCachedBuildPromotionsStatsParams contains all the parameters to send to the API endpoint
for the get cached build promotions stats operation typically these are written to a http.Request
*/
type GetCachedBuildPromotionsStatsParams struct {

	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cached build promotions stats params
func (o *GetCachedBuildPromotionsStatsParams) WithTimeout(timeout time.Duration) *GetCachedBuildPromotionsStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cached build promotions stats params
func (o *GetCachedBuildPromotionsStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cached build promotions stats params
func (o *GetCachedBuildPromotionsStatsParams) WithContext(ctx context.Context) *GetCachedBuildPromotionsStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cached build promotions stats params
func (o *GetCachedBuildPromotionsStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cached build promotions stats params
func (o *GetCachedBuildPromotionsStatsParams) WithHTTPClient(client *http.Client) *GetCachedBuildPromotionsStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cached build promotions stats params
func (o *GetCachedBuildPromotionsStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get cached build promotions stats params
func (o *GetCachedBuildPromotionsStatsParams) WithFields(fields *string) *GetCachedBuildPromotionsStatsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get cached build promotions stats params
func (o *GetCachedBuildPromotionsStatsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetCachedBuildPromotionsStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
