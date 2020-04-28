// Code generated by go-swagger; DO NOT EDIT.

package webhook

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

// NewPostWebhooksParams creates a new PostWebhooksParams object
// with the default values initialized.
func NewPostWebhooksParams() *PostWebhooksParams {
	var ()
	return &PostWebhooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWebhooksParamsWithTimeout creates a new PostWebhooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWebhooksParamsWithTimeout(timeout time.Duration) *PostWebhooksParams {
	var ()
	return &PostWebhooksParams{

		timeout: timeout,
	}
}

// NewPostWebhooksParamsWithContext creates a new PostWebhooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWebhooksParamsWithContext(ctx context.Context) *PostWebhooksParams {
	var ()
	return &PostWebhooksParams{

		Context: ctx,
	}
}

// NewPostWebhooksParamsWithHTTPClient creates a new PostWebhooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWebhooksParamsWithHTTPClient(client *http.Client) *PostWebhooksParams {
	var ()
	return &PostWebhooksParams{
		HTTPClient: client,
	}
}

/*PostWebhooksParams contains all the parameters to send to the API endpoint
for the post webhooks operation typically these are written to a http.Request
*/
type PostWebhooksParams struct {

	/*Body*/
	Body *models.WebHook

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post webhooks params
func (o *PostWebhooksParams) WithTimeout(timeout time.Duration) *PostWebhooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post webhooks params
func (o *PostWebhooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post webhooks params
func (o *PostWebhooksParams) WithContext(ctx context.Context) *PostWebhooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post webhooks params
func (o *PostWebhooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post webhooks params
func (o *PostWebhooksParams) WithHTTPClient(client *http.Client) *PostWebhooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post webhooks params
func (o *PostWebhooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post webhooks params
func (o *PostWebhooksParams) WithBody(body *models.WebHook) *PostWebhooksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post webhooks params
func (o *PostWebhooksParams) SetBody(body *models.WebHook) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWebhooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
