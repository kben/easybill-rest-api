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
)

// NewDeleteCustomersIDParams creates a new DeleteCustomersIDParams object
// with the default values initialized.
func NewDeleteCustomersIDParams() *DeleteCustomersIDParams {
	var ()
	return &DeleteCustomersIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCustomersIDParamsWithTimeout creates a new DeleteCustomersIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCustomersIDParamsWithTimeout(timeout time.Duration) *DeleteCustomersIDParams {
	var ()
	return &DeleteCustomersIDParams{

		timeout: timeout,
	}
}

// NewDeleteCustomersIDParamsWithContext creates a new DeleteCustomersIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCustomersIDParamsWithContext(ctx context.Context) *DeleteCustomersIDParams {
	var ()
	return &DeleteCustomersIDParams{

		Context: ctx,
	}
}

// NewDeleteCustomersIDParamsWithHTTPClient creates a new DeleteCustomersIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCustomersIDParamsWithHTTPClient(client *http.Client) *DeleteCustomersIDParams {
	var ()
	return &DeleteCustomersIDParams{
		HTTPClient: client,
	}
}

/*DeleteCustomersIDParams contains all the parameters to send to the API endpoint
for the delete customers ID operation typically these are written to a http.Request
*/
type DeleteCustomersIDParams struct {

	/*ID
	  ID of customer that needs to be deleted

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete customers ID params
func (o *DeleteCustomersIDParams) WithTimeout(timeout time.Duration) *DeleteCustomersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete customers ID params
func (o *DeleteCustomersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete customers ID params
func (o *DeleteCustomersIDParams) WithContext(ctx context.Context) *DeleteCustomersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete customers ID params
func (o *DeleteCustomersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete customers ID params
func (o *DeleteCustomersIDParams) WithHTTPClient(client *http.Client) *DeleteCustomersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete customers ID params
func (o *DeleteCustomersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete customers ID params
func (o *DeleteCustomersIDParams) WithID(id int64) *DeleteCustomersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete customers ID params
func (o *DeleteCustomersIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCustomersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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