// Code generated by go-swagger; DO NOT EDIT.

package position

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

// NewPutPositionsIDParams creates a new PutPositionsIDParams object
// with the default values initialized.
func NewPutPositionsIDParams() *PutPositionsIDParams {
	var ()
	return &PutPositionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPositionsIDParamsWithTimeout creates a new PutPositionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPositionsIDParamsWithTimeout(timeout time.Duration) *PutPositionsIDParams {
	var ()
	return &PutPositionsIDParams{

		timeout: timeout,
	}
}

// NewPutPositionsIDParamsWithContext creates a new PutPositionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPositionsIDParamsWithContext(ctx context.Context) *PutPositionsIDParams {
	var ()
	return &PutPositionsIDParams{

		Context: ctx,
	}
}

// NewPutPositionsIDParamsWithHTTPClient creates a new PutPositionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPositionsIDParamsWithHTTPClient(client *http.Client) *PutPositionsIDParams {
	var ()
	return &PutPositionsIDParams{
		HTTPClient: client,
	}
}

/*PutPositionsIDParams contains all the parameters to send to the API endpoint
for the put positions ID operation typically these are written to a http.Request
*/
type PutPositionsIDParams struct {

	/*Body*/
	Body *models.Position
	/*ID
	  ID of position

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put positions ID params
func (o *PutPositionsIDParams) WithTimeout(timeout time.Duration) *PutPositionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put positions ID params
func (o *PutPositionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put positions ID params
func (o *PutPositionsIDParams) WithContext(ctx context.Context) *PutPositionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put positions ID params
func (o *PutPositionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put positions ID params
func (o *PutPositionsIDParams) WithHTTPClient(client *http.Client) *PutPositionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put positions ID params
func (o *PutPositionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put positions ID params
func (o *PutPositionsIDParams) WithBody(body *models.Position) *PutPositionsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put positions ID params
func (o *PutPositionsIDParams) SetBody(body *models.Position) {
	o.Body = body
}

// WithID adds the id to the put positions ID params
func (o *PutPositionsIDParams) WithID(id int64) *PutPositionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put positions ID params
func (o *PutPositionsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutPositionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
