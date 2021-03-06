// Code generated by go-swagger; DO NOT EDIT.

package logins

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

// NewGetLoginsParams creates a new GetLoginsParams object
// with the default values initialized.
func NewGetLoginsParams() *GetLoginsParams {
	var ()
	return &GetLoginsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginsParamsWithTimeout creates a new GetLoginsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoginsParamsWithTimeout(timeout time.Duration) *GetLoginsParams {
	var ()
	return &GetLoginsParams{

		timeout: timeout,
	}
}

// NewGetLoginsParamsWithContext creates a new GetLoginsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoginsParamsWithContext(ctx context.Context) *GetLoginsParams {
	var ()
	return &GetLoginsParams{

		Context: ctx,
	}
}

// NewGetLoginsParamsWithHTTPClient creates a new GetLoginsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoginsParamsWithHTTPClient(client *http.Client) *GetLoginsParams {
	var ()
	return &GetLoginsParams{
		HTTPClient: client,
	}
}

/*GetLoginsParams contains all the parameters to send to the API endpoint
for the get logins operation typically these are written to a http.Request
*/
type GetLoginsParams struct {

	/*Limit
	  Limited the result. Default is 100. Maximum can be 1000.

	*/
	Limit *int64
	/*Page
	  Set current Page. Default is 1.

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get logins params
func (o *GetLoginsParams) WithTimeout(timeout time.Duration) *GetLoginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logins params
func (o *GetLoginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logins params
func (o *GetLoginsParams) WithContext(ctx context.Context) *GetLoginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logins params
func (o *GetLoginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logins params
func (o *GetLoginsParams) WithHTTPClient(client *http.Client) *GetLoginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logins params
func (o *GetLoginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get logins params
func (o *GetLoginsParams) WithLimit(limit *int64) *GetLoginsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get logins params
func (o *GetLoginsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get logins params
func (o *GetLoginsParams) WithPage(page *int64) *GetLoginsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get logins params
func (o *GetLoginsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
