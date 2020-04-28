// Code generated by go-swagger; DO NOT EDIT.

package sepa_payment

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

// NewPostSepaPaymentsParams creates a new PostSepaPaymentsParams object
// with the default values initialized.
func NewPostSepaPaymentsParams() *PostSepaPaymentsParams {
	var ()
	return &PostSepaPaymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSepaPaymentsParamsWithTimeout creates a new PostSepaPaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSepaPaymentsParamsWithTimeout(timeout time.Duration) *PostSepaPaymentsParams {
	var ()
	return &PostSepaPaymentsParams{

		timeout: timeout,
	}
}

// NewPostSepaPaymentsParamsWithContext creates a new PostSepaPaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSepaPaymentsParamsWithContext(ctx context.Context) *PostSepaPaymentsParams {
	var ()
	return &PostSepaPaymentsParams{

		Context: ctx,
	}
}

// NewPostSepaPaymentsParamsWithHTTPClient creates a new PostSepaPaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSepaPaymentsParamsWithHTTPClient(client *http.Client) *PostSepaPaymentsParams {
	var ()
	return &PostSepaPaymentsParams{
		HTTPClient: client,
	}
}

/*PostSepaPaymentsParams contains all the parameters to send to the API endpoint
for the post sepa payments operation typically these are written to a http.Request
*/
type PostSepaPaymentsParams struct {

	/*Body*/
	Body *models.SEPAPayment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sepa payments params
func (o *PostSepaPaymentsParams) WithTimeout(timeout time.Duration) *PostSepaPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sepa payments params
func (o *PostSepaPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sepa payments params
func (o *PostSepaPaymentsParams) WithContext(ctx context.Context) *PostSepaPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sepa payments params
func (o *PostSepaPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sepa payments params
func (o *PostSepaPaymentsParams) WithHTTPClient(client *http.Client) *PostSepaPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sepa payments params
func (o *PostSepaPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sepa payments params
func (o *PostSepaPaymentsParams) WithBody(body *models.SEPAPayment) *PostSepaPaymentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sepa payments params
func (o *PostSepaPaymentsParams) SetBody(body *models.SEPAPayment) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSepaPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
