// Code generated by go-swagger; DO NOT EDIT.

package document_payment

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

// NewGetDocumentPaymentsIDParams creates a new GetDocumentPaymentsIDParams object
// with the default values initialized.
func NewGetDocumentPaymentsIDParams() *GetDocumentPaymentsIDParams {
	var ()
	return &GetDocumentPaymentsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDocumentPaymentsIDParamsWithTimeout creates a new GetDocumentPaymentsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDocumentPaymentsIDParamsWithTimeout(timeout time.Duration) *GetDocumentPaymentsIDParams {
	var ()
	return &GetDocumentPaymentsIDParams{

		timeout: timeout,
	}
}

// NewGetDocumentPaymentsIDParamsWithContext creates a new GetDocumentPaymentsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDocumentPaymentsIDParamsWithContext(ctx context.Context) *GetDocumentPaymentsIDParams {
	var ()
	return &GetDocumentPaymentsIDParams{

		Context: ctx,
	}
}

// NewGetDocumentPaymentsIDParamsWithHTTPClient creates a new GetDocumentPaymentsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDocumentPaymentsIDParamsWithHTTPClient(client *http.Client) *GetDocumentPaymentsIDParams {
	var ()
	return &GetDocumentPaymentsIDParams{
		HTTPClient: client,
	}
}

/*GetDocumentPaymentsIDParams contains all the parameters to send to the API endpoint
for the get document payments ID operation typically these are written to a http.Request
*/
type GetDocumentPaymentsIDParams struct {

	/*ID
	  ID of document payment

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get document payments ID params
func (o *GetDocumentPaymentsIDParams) WithTimeout(timeout time.Duration) *GetDocumentPaymentsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get document payments ID params
func (o *GetDocumentPaymentsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get document payments ID params
func (o *GetDocumentPaymentsIDParams) WithContext(ctx context.Context) *GetDocumentPaymentsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get document payments ID params
func (o *GetDocumentPaymentsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get document payments ID params
func (o *GetDocumentPaymentsIDParams) WithHTTPClient(client *http.Client) *GetDocumentPaymentsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get document payments ID params
func (o *GetDocumentPaymentsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get document payments ID params
func (o *GetDocumentPaymentsIDParams) WithID(id int64) *GetDocumentPaymentsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get document payments ID params
func (o *GetDocumentPaymentsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDocumentPaymentsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
