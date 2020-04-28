// Code generated by go-swagger; DO NOT EDIT.

package attachment

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

// NewGetAttachmentsParams creates a new GetAttachmentsParams object
// with the default values initialized.
func NewGetAttachmentsParams() *GetAttachmentsParams {
	var ()
	return &GetAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAttachmentsParamsWithTimeout creates a new GetAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAttachmentsParamsWithTimeout(timeout time.Duration) *GetAttachmentsParams {
	var ()
	return &GetAttachmentsParams{

		timeout: timeout,
	}
}

// NewGetAttachmentsParamsWithContext creates a new GetAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAttachmentsParamsWithContext(ctx context.Context) *GetAttachmentsParams {
	var ()
	return &GetAttachmentsParams{

		Context: ctx,
	}
}

// NewGetAttachmentsParamsWithHTTPClient creates a new GetAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAttachmentsParamsWithHTTPClient(client *http.Client) *GetAttachmentsParams {
	var ()
	return &GetAttachmentsParams{
		HTTPClient: client,
	}
}

/*GetAttachmentsParams contains all the parameters to send to the API endpoint
for the get attachments operation typically these are written to a http.Request
*/
type GetAttachmentsParams struct {

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

// WithTimeout adds the timeout to the get attachments params
func (o *GetAttachmentsParams) WithTimeout(timeout time.Duration) *GetAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get attachments params
func (o *GetAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get attachments params
func (o *GetAttachmentsParams) WithContext(ctx context.Context) *GetAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get attachments params
func (o *GetAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get attachments params
func (o *GetAttachmentsParams) WithHTTPClient(client *http.Client) *GetAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get attachments params
func (o *GetAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get attachments params
func (o *GetAttachmentsParams) WithLimit(limit *int64) *GetAttachmentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get attachments params
func (o *GetAttachmentsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get attachments params
func (o *GetAttachmentsParams) WithPage(page *int64) *GetAttachmentsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get attachments params
func (o *GetAttachmentsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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