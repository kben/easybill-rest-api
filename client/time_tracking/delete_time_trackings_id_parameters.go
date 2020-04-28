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

// NewDeleteTimeTrackingsIDParams creates a new DeleteTimeTrackingsIDParams object
// with the default values initialized.
func NewDeleteTimeTrackingsIDParams() *DeleteTimeTrackingsIDParams {
	var ()
	return &DeleteTimeTrackingsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTimeTrackingsIDParamsWithTimeout creates a new DeleteTimeTrackingsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTimeTrackingsIDParamsWithTimeout(timeout time.Duration) *DeleteTimeTrackingsIDParams {
	var ()
	return &DeleteTimeTrackingsIDParams{

		timeout: timeout,
	}
}

// NewDeleteTimeTrackingsIDParamsWithContext creates a new DeleteTimeTrackingsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTimeTrackingsIDParamsWithContext(ctx context.Context) *DeleteTimeTrackingsIDParams {
	var ()
	return &DeleteTimeTrackingsIDParams{

		Context: ctx,
	}
}

// NewDeleteTimeTrackingsIDParamsWithHTTPClient creates a new DeleteTimeTrackingsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTimeTrackingsIDParamsWithHTTPClient(client *http.Client) *DeleteTimeTrackingsIDParams {
	var ()
	return &DeleteTimeTrackingsIDParams{
		HTTPClient: client,
	}
}

/*DeleteTimeTrackingsIDParams contains all the parameters to send to the API endpoint
for the delete time trackings ID operation typically these are written to a http.Request
*/
type DeleteTimeTrackingsIDParams struct {

	/*ID
	  ID of time tracking

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete time trackings ID params
func (o *DeleteTimeTrackingsIDParams) WithTimeout(timeout time.Duration) *DeleteTimeTrackingsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete time trackings ID params
func (o *DeleteTimeTrackingsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete time trackings ID params
func (o *DeleteTimeTrackingsIDParams) WithContext(ctx context.Context) *DeleteTimeTrackingsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete time trackings ID params
func (o *DeleteTimeTrackingsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete time trackings ID params
func (o *DeleteTimeTrackingsIDParams) WithHTTPClient(client *http.Client) *DeleteTimeTrackingsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete time trackings ID params
func (o *DeleteTimeTrackingsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete time trackings ID params
func (o *DeleteTimeTrackingsIDParams) WithID(id int64) *DeleteTimeTrackingsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete time trackings ID params
func (o *DeleteTimeTrackingsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTimeTrackingsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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