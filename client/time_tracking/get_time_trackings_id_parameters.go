// Code generated by go-swagger; DO NOT EDIT.

package time_tracking

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

// NewGetTimeTrackingsIDParams creates a new GetTimeTrackingsIDParams object
// with the default values initialized.
func NewGetTimeTrackingsIDParams() *GetTimeTrackingsIDParams {
	var ()
	return &GetTimeTrackingsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTimeTrackingsIDParamsWithTimeout creates a new GetTimeTrackingsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTimeTrackingsIDParamsWithTimeout(timeout time.Duration) *GetTimeTrackingsIDParams {
	var ()
	return &GetTimeTrackingsIDParams{

		timeout: timeout,
	}
}

// NewGetTimeTrackingsIDParamsWithContext creates a new GetTimeTrackingsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTimeTrackingsIDParamsWithContext(ctx context.Context) *GetTimeTrackingsIDParams {
	var ()
	return &GetTimeTrackingsIDParams{

		Context: ctx,
	}
}

// NewGetTimeTrackingsIDParamsWithHTTPClient creates a new GetTimeTrackingsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTimeTrackingsIDParamsWithHTTPClient(client *http.Client) *GetTimeTrackingsIDParams {
	var ()
	return &GetTimeTrackingsIDParams{
		HTTPClient: client,
	}
}

/*GetTimeTrackingsIDParams contains all the parameters to send to the API endpoint
for the get time trackings ID operation typically these are written to a http.Request
*/
type GetTimeTrackingsIDParams struct {

	/*ID
	  ID of time tracking

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get time trackings ID params
func (o *GetTimeTrackingsIDParams) WithTimeout(timeout time.Duration) *GetTimeTrackingsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get time trackings ID params
func (o *GetTimeTrackingsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get time trackings ID params
func (o *GetTimeTrackingsIDParams) WithContext(ctx context.Context) *GetTimeTrackingsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get time trackings ID params
func (o *GetTimeTrackingsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get time trackings ID params
func (o *GetTimeTrackingsIDParams) WithHTTPClient(client *http.Client) *GetTimeTrackingsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get time trackings ID params
func (o *GetTimeTrackingsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get time trackings ID params
func (o *GetTimeTrackingsIDParams) WithID(id int64) *GetTimeTrackingsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get time trackings ID params
func (o *GetTimeTrackingsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTimeTrackingsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
