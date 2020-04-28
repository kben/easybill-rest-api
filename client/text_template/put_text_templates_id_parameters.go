// Code generated by go-swagger; DO NOT EDIT.

package text_template

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

// NewPutTextTemplatesIDParams creates a new PutTextTemplatesIDParams object
// with the default values initialized.
func NewPutTextTemplatesIDParams() *PutTextTemplatesIDParams {
	var ()
	return &PutTextTemplatesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTextTemplatesIDParamsWithTimeout creates a new PutTextTemplatesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTextTemplatesIDParamsWithTimeout(timeout time.Duration) *PutTextTemplatesIDParams {
	var ()
	return &PutTextTemplatesIDParams{

		timeout: timeout,
	}
}

// NewPutTextTemplatesIDParamsWithContext creates a new PutTextTemplatesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTextTemplatesIDParamsWithContext(ctx context.Context) *PutTextTemplatesIDParams {
	var ()
	return &PutTextTemplatesIDParams{

		Context: ctx,
	}
}

// NewPutTextTemplatesIDParamsWithHTTPClient creates a new PutTextTemplatesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTextTemplatesIDParamsWithHTTPClient(client *http.Client) *PutTextTemplatesIDParams {
	var ()
	return &PutTextTemplatesIDParams{
		HTTPClient: client,
	}
}

/*PutTextTemplatesIDParams contains all the parameters to send to the API endpoint
for the put text templates ID operation typically these are written to a http.Request
*/
type PutTextTemplatesIDParams struct {

	/*Body*/
	Body *models.TextTemplate
	/*ID
	  ID of text template

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put text templates ID params
func (o *PutTextTemplatesIDParams) WithTimeout(timeout time.Duration) *PutTextTemplatesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put text templates ID params
func (o *PutTextTemplatesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put text templates ID params
func (o *PutTextTemplatesIDParams) WithContext(ctx context.Context) *PutTextTemplatesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put text templates ID params
func (o *PutTextTemplatesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put text templates ID params
func (o *PutTextTemplatesIDParams) WithHTTPClient(client *http.Client) *PutTextTemplatesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put text templates ID params
func (o *PutTextTemplatesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put text templates ID params
func (o *PutTextTemplatesIDParams) WithBody(body *models.TextTemplate) *PutTextTemplatesIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put text templates ID params
func (o *PutTextTemplatesIDParams) SetBody(body *models.TextTemplate) {
	o.Body = body
}

// WithID adds the id to the put text templates ID params
func (o *PutTextTemplatesIDParams) WithID(id int64) *PutTextTemplatesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put text templates ID params
func (o *PutTextTemplatesIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutTextTemplatesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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