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
)

// NewGetPositionsIDParams creates a new GetPositionsIDParams object
// with the default values initialized.
func NewGetPositionsIDParams() *GetPositionsIDParams {
	var ()
	return &GetPositionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPositionsIDParamsWithTimeout creates a new GetPositionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPositionsIDParamsWithTimeout(timeout time.Duration) *GetPositionsIDParams {
	var ()
	return &GetPositionsIDParams{

		timeout: timeout,
	}
}

// NewGetPositionsIDParamsWithContext creates a new GetPositionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPositionsIDParamsWithContext(ctx context.Context) *GetPositionsIDParams {
	var ()
	return &GetPositionsIDParams{

		Context: ctx,
	}
}

// NewGetPositionsIDParamsWithHTTPClient creates a new GetPositionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPositionsIDParamsWithHTTPClient(client *http.Client) *GetPositionsIDParams {
	var ()
	return &GetPositionsIDParams{
		HTTPClient: client,
	}
}

/*GetPositionsIDParams contains all the parameters to send to the API endpoint
for the get positions ID operation typically these are written to a http.Request
*/
type GetPositionsIDParams struct {

	/*ID
	  ID of position

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get positions ID params
func (o *GetPositionsIDParams) WithTimeout(timeout time.Duration) *GetPositionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get positions ID params
func (o *GetPositionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get positions ID params
func (o *GetPositionsIDParams) WithContext(ctx context.Context) *GetPositionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get positions ID params
func (o *GetPositionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get positions ID params
func (o *GetPositionsIDParams) WithHTTPClient(client *http.Client) *GetPositionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get positions ID params
func (o *GetPositionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get positions ID params
func (o *GetPositionsIDParams) WithID(id int64) *GetPositionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get positions ID params
func (o *GetPositionsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPositionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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