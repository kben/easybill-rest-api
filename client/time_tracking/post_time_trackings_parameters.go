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

	"github.com/kben/easybill-rest-api/models"
)

// NewPostTimeTrackingsParams creates a new PostTimeTrackingsParams object
// with the default values initialized.
func NewPostTimeTrackingsParams() *PostTimeTrackingsParams {
	var ()
	return &PostTimeTrackingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTimeTrackingsParamsWithTimeout creates a new PostTimeTrackingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTimeTrackingsParamsWithTimeout(timeout time.Duration) *PostTimeTrackingsParams {
	var ()
	return &PostTimeTrackingsParams{

		timeout: timeout,
	}
}

// NewPostTimeTrackingsParamsWithContext creates a new PostTimeTrackingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTimeTrackingsParamsWithContext(ctx context.Context) *PostTimeTrackingsParams {
	var ()
	return &PostTimeTrackingsParams{

		Context: ctx,
	}
}

// NewPostTimeTrackingsParamsWithHTTPClient creates a new PostTimeTrackingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTimeTrackingsParamsWithHTTPClient(client *http.Client) *PostTimeTrackingsParams {
	var ()
	return &PostTimeTrackingsParams{
		HTTPClient: client,
	}
}

/*PostTimeTrackingsParams contains all the parameters to send to the API endpoint
for the post time trackings operation typically these are written to a http.Request
*/
type PostTimeTrackingsParams struct {

	/*Body*/
	Body *models.TimeTracking

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post time trackings params
func (o *PostTimeTrackingsParams) WithTimeout(timeout time.Duration) *PostTimeTrackingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post time trackings params
func (o *PostTimeTrackingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post time trackings params
func (o *PostTimeTrackingsParams) WithContext(ctx context.Context) *PostTimeTrackingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post time trackings params
func (o *PostTimeTrackingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post time trackings params
func (o *PostTimeTrackingsParams) WithHTTPClient(client *http.Client) *PostTimeTrackingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post time trackings params
func (o *PostTimeTrackingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post time trackings params
func (o *PostTimeTrackingsParams) WithBody(body *models.TimeTracking) *PostTimeTrackingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post time trackings params
func (o *PostTimeTrackingsParams) SetBody(body *models.TimeTracking) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTimeTrackingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
