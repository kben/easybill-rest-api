// Code generated by go-swagger; DO NOT EDIT.

package sepa_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetSepaPaymentsIDParams creates a new GetSepaPaymentsIDParams object
// with the default values initialized.
func NewGetSepaPaymentsIDParams() *GetSepaPaymentsIDParams {
	var ()
	return &GetSepaPaymentsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSepaPaymentsIDParamsWithTimeout creates a new GetSepaPaymentsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSepaPaymentsIDParamsWithTimeout(timeout time.Duration) *GetSepaPaymentsIDParams {
	var ()
	return &GetSepaPaymentsIDParams{

		timeout: timeout,
	}
}

// NewGetSepaPaymentsIDParamsWithContext creates a new GetSepaPaymentsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSepaPaymentsIDParamsWithContext(ctx context.Context) *GetSepaPaymentsIDParams {
	var ()
	return &GetSepaPaymentsIDParams{

		Context: ctx,
	}
}

// NewGetSepaPaymentsIDParamsWithHTTPClient creates a new GetSepaPaymentsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSepaPaymentsIDParamsWithHTTPClient(client *http.Client) *GetSepaPaymentsIDParams {
	var ()
	return &GetSepaPaymentsIDParams{
		HTTPClient: client,
	}
}

/*GetSepaPaymentsIDParams contains all the parameters to send to the API endpoint
for the get sepa payments ID operation typically these are written to a http.Request
*/
type GetSepaPaymentsIDParams struct {

	/*ID
	  ID of SEPA payment

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sepa payments ID params
func (o *GetSepaPaymentsIDParams) WithTimeout(timeout time.Duration) *GetSepaPaymentsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sepa payments ID params
func (o *GetSepaPaymentsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sepa payments ID params
func (o *GetSepaPaymentsIDParams) WithContext(ctx context.Context) *GetSepaPaymentsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sepa payments ID params
func (o *GetSepaPaymentsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sepa payments ID params
func (o *GetSepaPaymentsIDParams) WithHTTPClient(client *http.Client) *GetSepaPaymentsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sepa payments ID params
func (o *GetSepaPaymentsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get sepa payments ID params
func (o *GetSepaPaymentsIDParams) WithID(id int64) *GetSepaPaymentsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get sepa payments ID params
func (o *GetSepaPaymentsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSepaPaymentsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
