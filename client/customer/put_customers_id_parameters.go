// Code generated by go-swagger; DO NOT EDIT.

package customer

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

	"github.com/kben/easybill-rest-api/models"
)

// NewPutCustomersIDParams creates a new PutCustomersIDParams object
// with the default values initialized.
func NewPutCustomersIDParams() *PutCustomersIDParams {
	var ()
	return &PutCustomersIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersIDParamsWithTimeout creates a new PutCustomersIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersIDParamsWithTimeout(timeout time.Duration) *PutCustomersIDParams {
	var ()
	return &PutCustomersIDParams{

		timeout: timeout,
	}
}

// NewPutCustomersIDParamsWithContext creates a new PutCustomersIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersIDParamsWithContext(ctx context.Context) *PutCustomersIDParams {
	var ()
	return &PutCustomersIDParams{

		Context: ctx,
	}
}

// NewPutCustomersIDParamsWithHTTPClient creates a new PutCustomersIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersIDParamsWithHTTPClient(client *http.Client) *PutCustomersIDParams {
	var ()
	return &PutCustomersIDParams{
		HTTPClient: client,
	}
}

/*PutCustomersIDParams contains all the parameters to send to the API endpoint
for the put customers ID operation typically these are written to a http.Request
*/
type PutCustomersIDParams struct {

	/*Body*/
	Body *models.Customer
	/*ID
	  ID of customer that needs to be updated

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers ID params
func (o *PutCustomersIDParams) WithTimeout(timeout time.Duration) *PutCustomersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers ID params
func (o *PutCustomersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers ID params
func (o *PutCustomersIDParams) WithContext(ctx context.Context) *PutCustomersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers ID params
func (o *PutCustomersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers ID params
func (o *PutCustomersIDParams) WithHTTPClient(client *http.Client) *PutCustomersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers ID params
func (o *PutCustomersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put customers ID params
func (o *PutCustomersIDParams) WithBody(body *models.Customer) *PutCustomersIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put customers ID params
func (o *PutCustomersIDParams) SetBody(body *models.Customer) {
	o.Body = body
}

// WithID adds the id to the put customers ID params
func (o *PutCustomersIDParams) WithID(id int64) *PutCustomersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put customers ID params
func (o *PutCustomersIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}