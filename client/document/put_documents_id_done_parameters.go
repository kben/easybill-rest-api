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

// NewPutDocumentsIDDoneParams creates a new PutDocumentsIDDoneParams object
// with the default values initialized.
func NewPutDocumentsIDDoneParams() *PutDocumentsIDDoneParams {
	var ()
	return &PutDocumentsIDDoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDocumentsIDDoneParamsWithTimeout creates a new PutDocumentsIDDoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDocumentsIDDoneParamsWithTimeout(timeout time.Duration) *PutDocumentsIDDoneParams {
	var ()
	return &PutDocumentsIDDoneParams{

		timeout: timeout,
	}
}

// NewPutDocumentsIDDoneParamsWithContext creates a new PutDocumentsIDDoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDocumentsIDDoneParamsWithContext(ctx context.Context) *PutDocumentsIDDoneParams {
	var ()
	return &PutDocumentsIDDoneParams{

		Context: ctx,
	}
}

// NewPutDocumentsIDDoneParamsWithHTTPClient creates a new PutDocumentsIDDoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDocumentsIDDoneParamsWithHTTPClient(client *http.Client) *PutDocumentsIDDoneParams {
	var ()
	return &PutDocumentsIDDoneParams{
		HTTPClient: client,
	}
}

/*PutDocumentsIDDoneParams contains all the parameters to send to the API endpoint
for the put documents ID done operation typically these are written to a http.Request
*/
type PutDocumentsIDDoneParams struct {

	/*ID
	  ID of document

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put documents ID done params
func (o *PutDocumentsIDDoneParams) WithTimeout(timeout time.Duration) *PutDocumentsIDDoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put documents ID done params
func (o *PutDocumentsIDDoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put documents ID done params
func (o *PutDocumentsIDDoneParams) WithContext(ctx context.Context) *PutDocumentsIDDoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put documents ID done params
func (o *PutDocumentsIDDoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put documents ID done params
func (o *PutDocumentsIDDoneParams) WithHTTPClient(client *http.Client) *PutDocumentsIDDoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put documents ID done params
func (o *PutDocumentsIDDoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put documents ID done params
func (o *PutDocumentsIDDoneParams) WithID(id int64) *PutDocumentsIDDoneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put documents ID done params
func (o *PutDocumentsIDDoneParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDocumentsIDDoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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