// Code generated by go-swagger; DO NOT EDIT.

package document

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

// NewGetDocumentsIDPdfParams creates a new GetDocumentsIDPdfParams object
// with the default values initialized.
func NewGetDocumentsIDPdfParams() *GetDocumentsIDPdfParams {
	var ()
	return &GetDocumentsIDPdfParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDocumentsIDPdfParamsWithTimeout creates a new GetDocumentsIDPdfParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDocumentsIDPdfParamsWithTimeout(timeout time.Duration) *GetDocumentsIDPdfParams {
	var ()
	return &GetDocumentsIDPdfParams{

		timeout: timeout,
	}
}

// NewGetDocumentsIDPdfParamsWithContext creates a new GetDocumentsIDPdfParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDocumentsIDPdfParamsWithContext(ctx context.Context) *GetDocumentsIDPdfParams {
	var ()
	return &GetDocumentsIDPdfParams{

		Context: ctx,
	}
}

// NewGetDocumentsIDPdfParamsWithHTTPClient creates a new GetDocumentsIDPdfParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDocumentsIDPdfParamsWithHTTPClient(client *http.Client) *GetDocumentsIDPdfParams {
	var ()
	return &GetDocumentsIDPdfParams{
		HTTPClient: client,
	}
}

/*GetDocumentsIDPdfParams contains all the parameters to send to the API endpoint
for the get documents ID pdf operation typically these are written to a http.Request
*/
type GetDocumentsIDPdfParams struct {

	/*ID
	  ID of document

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get documents ID pdf params
func (o *GetDocumentsIDPdfParams) WithTimeout(timeout time.Duration) *GetDocumentsIDPdfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get documents ID pdf params
func (o *GetDocumentsIDPdfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get documents ID pdf params
func (o *GetDocumentsIDPdfParams) WithContext(ctx context.Context) *GetDocumentsIDPdfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get documents ID pdf params
func (o *GetDocumentsIDPdfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get documents ID pdf params
func (o *GetDocumentsIDPdfParams) WithHTTPClient(client *http.Client) *GetDocumentsIDPdfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get documents ID pdf params
func (o *GetDocumentsIDPdfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get documents ID pdf params
func (o *GetDocumentsIDPdfParams) WithID(id int64) *GetDocumentsIDPdfParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get documents ID pdf params
func (o *GetDocumentsIDPdfParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDocumentsIDPdfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
