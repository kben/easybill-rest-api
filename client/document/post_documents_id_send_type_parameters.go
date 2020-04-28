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

	"github.com/kben/easybill-rest-api/models"
)

// NewPostDocumentsIDSendTypeParams creates a new PostDocumentsIDSendTypeParams object
// with the default values initialized.
func NewPostDocumentsIDSendTypeParams() *PostDocumentsIDSendTypeParams {
	var ()
	return &PostDocumentsIDSendTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDocumentsIDSendTypeParamsWithTimeout creates a new PostDocumentsIDSendTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDocumentsIDSendTypeParamsWithTimeout(timeout time.Duration) *PostDocumentsIDSendTypeParams {
	var ()
	return &PostDocumentsIDSendTypeParams{

		timeout: timeout,
	}
}

// NewPostDocumentsIDSendTypeParamsWithContext creates a new PostDocumentsIDSendTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDocumentsIDSendTypeParamsWithContext(ctx context.Context) *PostDocumentsIDSendTypeParams {
	var ()
	return &PostDocumentsIDSendTypeParams{

		Context: ctx,
	}
}

// NewPostDocumentsIDSendTypeParamsWithHTTPClient creates a new PostDocumentsIDSendTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDocumentsIDSendTypeParamsWithHTTPClient(client *http.Client) *PostDocumentsIDSendTypeParams {
	var ()
	return &PostDocumentsIDSendTypeParams{
		HTTPClient: client,
	}
}

/*PostDocumentsIDSendTypeParams contains all the parameters to send to the API endpoint
for the post documents ID send type operation typically these are written to a http.Request
*/
type PostDocumentsIDSendTypeParams struct {

	/*Body*/
	Body *models.PostBoxRequest
	/*ID
	  ID of document

	*/
	ID int64
	/*Type*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) WithTimeout(timeout time.Duration) *PostDocumentsIDSendTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) WithContext(ctx context.Context) *PostDocumentsIDSendTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) WithHTTPClient(client *http.Client) *PostDocumentsIDSendTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) WithBody(body *models.PostBoxRequest) *PostDocumentsIDSendTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) SetBody(body *models.PostBoxRequest) {
	o.Body = body
}

// WithID adds the id to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) WithID(id int64) *PostDocumentsIDSendTypeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) SetID(id int64) {
	o.ID = id
}

// WithType adds the typeVar to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) WithType(typeVar string) *PostDocumentsIDSendTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the post documents ID send type params
func (o *PostDocumentsIDSendTypeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PostDocumentsIDSendTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
