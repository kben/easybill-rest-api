// Code generated by go-swagger; DO NOT EDIT.

package customer_group

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

// NewPostCustomerGroupsParams creates a new PostCustomerGroupsParams object
// with the default values initialized.
func NewPostCustomerGroupsParams() *PostCustomerGroupsParams {
	var ()
	return &PostCustomerGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomerGroupsParamsWithTimeout creates a new PostCustomerGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomerGroupsParamsWithTimeout(timeout time.Duration) *PostCustomerGroupsParams {
	var ()
	return &PostCustomerGroupsParams{

		timeout: timeout,
	}
}

// NewPostCustomerGroupsParamsWithContext creates a new PostCustomerGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomerGroupsParamsWithContext(ctx context.Context) *PostCustomerGroupsParams {
	var ()
	return &PostCustomerGroupsParams{

		Context: ctx,
	}
}

// NewPostCustomerGroupsParamsWithHTTPClient creates a new PostCustomerGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomerGroupsParamsWithHTTPClient(client *http.Client) *PostCustomerGroupsParams {
	var ()
	return &PostCustomerGroupsParams{
		HTTPClient: client,
	}
}

/*PostCustomerGroupsParams contains all the parameters to send to the API endpoint
for the post customer groups operation typically these are written to a http.Request
*/
type PostCustomerGroupsParams struct {

	/*Body*/
	Body *models.CustomerGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customer groups params
func (o *PostCustomerGroupsParams) WithTimeout(timeout time.Duration) *PostCustomerGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customer groups params
func (o *PostCustomerGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customer groups params
func (o *PostCustomerGroupsParams) WithContext(ctx context.Context) *PostCustomerGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customer groups params
func (o *PostCustomerGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customer groups params
func (o *PostCustomerGroupsParams) WithHTTPClient(client *http.Client) *PostCustomerGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customer groups params
func (o *PostCustomerGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post customer groups params
func (o *PostCustomerGroupsParams) WithBody(body *models.CustomerGroup) *PostCustomerGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post customer groups params
func (o *PostCustomerGroupsParams) SetBody(body *models.CustomerGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomerGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
