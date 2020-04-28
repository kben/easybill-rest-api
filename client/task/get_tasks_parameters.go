// Code generated by go-swagger; DO NOT EDIT.

package task

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

// NewGetTasksParams creates a new GetTasksParams object
// with the default values initialized.
func NewGetTasksParams() *GetTasksParams {
	var ()
	return &GetTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksParamsWithTimeout creates a new GetTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTasksParamsWithTimeout(timeout time.Duration) *GetTasksParams {
	var ()
	return &GetTasksParams{

		timeout: timeout,
	}
}

// NewGetTasksParamsWithContext creates a new GetTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTasksParamsWithContext(ctx context.Context) *GetTasksParams {
	var ()
	return &GetTasksParams{

		Context: ctx,
	}
}

// NewGetTasksParamsWithHTTPClient creates a new GetTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTasksParamsWithHTTPClient(client *http.Client) *GetTasksParams {
	var ()
	return &GetTasksParams{
		HTTPClient: client,
	}
}

/*GetTasksParams contains all the parameters to send to the API endpoint
for the get tasks operation typically these are written to a http.Request
*/
type GetTasksParams struct {

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

// WithTimeout adds the timeout to the get tasks params
func (o *GetTasksParams) WithTimeout(timeout time.Duration) *GetTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks params
func (o *GetTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks params
func (o *GetTasksParams) WithContext(ctx context.Context) *GetTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks params
func (o *GetTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks params
func (o *GetTasksParams) WithHTTPClient(client *http.Client) *GetTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks params
func (o *GetTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get tasks params
func (o *GetTasksParams) WithLimit(limit *int64) *GetTasksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get tasks params
func (o *GetTasksParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get tasks params
func (o *GetTasksParams) WithPage(page *int64) *GetTasksParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get tasks params
func (o *GetTasksParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
