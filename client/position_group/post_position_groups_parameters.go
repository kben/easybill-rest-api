// Code generated by go-swagger; DO NOT EDIT.

package position_group

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

// NewPostPositionGroupsParams creates a new PostPositionGroupsParams object
// with the default values initialized.
func NewPostPositionGroupsParams() *PostPositionGroupsParams {
	var ()
	return &PostPositionGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPositionGroupsParamsWithTimeout creates a new PostPositionGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPositionGroupsParamsWithTimeout(timeout time.Duration) *PostPositionGroupsParams {
	var ()
	return &PostPositionGroupsParams{

		timeout: timeout,
	}
}

// NewPostPositionGroupsParamsWithContext creates a new PostPositionGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPositionGroupsParamsWithContext(ctx context.Context) *PostPositionGroupsParams {
	var ()
	return &PostPositionGroupsParams{

		Context: ctx,
	}
}

// NewPostPositionGroupsParamsWithHTTPClient creates a new PostPositionGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPositionGroupsParamsWithHTTPClient(client *http.Client) *PostPositionGroupsParams {
	var ()
	return &PostPositionGroupsParams{
		HTTPClient: client,
	}
}

/*PostPositionGroupsParams contains all the parameters to send to the API endpoint
for the post position groups operation typically these are written to a http.Request
*/
type PostPositionGroupsParams struct {

	/*Body*/
	Body *models.PositionGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post position groups params
func (o *PostPositionGroupsParams) WithTimeout(timeout time.Duration) *PostPositionGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post position groups params
func (o *PostPositionGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post position groups params
func (o *PostPositionGroupsParams) WithContext(ctx context.Context) *PostPositionGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post position groups params
func (o *PostPositionGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post position groups params
func (o *PostPositionGroupsParams) WithHTTPClient(client *http.Client) *PostPositionGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post position groups params
func (o *PostPositionGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post position groups params
func (o *PostPositionGroupsParams) WithBody(body *models.PositionGroup) *PostPositionGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post position groups params
func (o *PostPositionGroupsParams) SetBody(body *models.PositionGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPositionGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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