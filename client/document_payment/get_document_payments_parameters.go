// Code generated by go-swagger; DO NOT EDIT.

package document_payment

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

// NewGetDocumentPaymentsParams creates a new GetDocumentPaymentsParams object
// with the default values initialized.
func NewGetDocumentPaymentsParams() *GetDocumentPaymentsParams {
	var ()
	return &GetDocumentPaymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDocumentPaymentsParamsWithTimeout creates a new GetDocumentPaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDocumentPaymentsParamsWithTimeout(timeout time.Duration) *GetDocumentPaymentsParams {
	var ()
	return &GetDocumentPaymentsParams{

		timeout: timeout,
	}
}

// NewGetDocumentPaymentsParamsWithContext creates a new GetDocumentPaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDocumentPaymentsParamsWithContext(ctx context.Context) *GetDocumentPaymentsParams {
	var ()
	return &GetDocumentPaymentsParams{

		Context: ctx,
	}
}

// NewGetDocumentPaymentsParamsWithHTTPClient creates a new GetDocumentPaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDocumentPaymentsParamsWithHTTPClient(client *http.Client) *GetDocumentPaymentsParams {
	var ()
	return &GetDocumentPaymentsParams{
		HTTPClient: client,
	}
}

/*GetDocumentPaymentsParams contains all the parameters to send to the API endpoint
for the get document payments operation typically these are written to a http.Request
*/
type GetDocumentPaymentsParams struct {

	/*DocumentID
	  Filter payments by document_id. You can add multiple ids separate by comma like id,id,id.

	*/
	DocumentID *string
	/*Limit
	  Limited the result. Default is 100. Maximum can be 1000.

	*/
	Limit *int64
	/*Page
	  Set current Page. Default is 1.

	*/
	Page *int64
	/*PaymentAt
	  Filter payments by payment_at. You can filter one date with payment_at=2014-12-10 or between like 2015-01-01,2015-12-31.

	*/
	PaymentAt *string
	/*Reference
	  Filter payments by reference. You can add multiple references separate by comma like id,id,id.

	*/
	Reference *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get document payments params
func (o *GetDocumentPaymentsParams) WithTimeout(timeout time.Duration) *GetDocumentPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get document payments params
func (o *GetDocumentPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get document payments params
func (o *GetDocumentPaymentsParams) WithContext(ctx context.Context) *GetDocumentPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get document payments params
func (o *GetDocumentPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get document payments params
func (o *GetDocumentPaymentsParams) WithHTTPClient(client *http.Client) *GetDocumentPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get document payments params
func (o *GetDocumentPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get document payments params
func (o *GetDocumentPaymentsParams) WithDocumentID(documentID *string) *GetDocumentPaymentsParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get document payments params
func (o *GetDocumentPaymentsParams) SetDocumentID(documentID *string) {
	o.DocumentID = documentID
}

// WithLimit adds the limit to the get document payments params
func (o *GetDocumentPaymentsParams) WithLimit(limit *int64) *GetDocumentPaymentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get document payments params
func (o *GetDocumentPaymentsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get document payments params
func (o *GetDocumentPaymentsParams) WithPage(page *int64) *GetDocumentPaymentsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get document payments params
func (o *GetDocumentPaymentsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPaymentAt adds the paymentAt to the get document payments params
func (o *GetDocumentPaymentsParams) WithPaymentAt(paymentAt *string) *GetDocumentPaymentsParams {
	o.SetPaymentAt(paymentAt)
	return o
}

// SetPaymentAt adds the paymentAt to the get document payments params
func (o *GetDocumentPaymentsParams) SetPaymentAt(paymentAt *string) {
	o.PaymentAt = paymentAt
}

// WithReference adds the reference to the get document payments params
func (o *GetDocumentPaymentsParams) WithReference(reference *string) *GetDocumentPaymentsParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the get document payments params
func (o *GetDocumentPaymentsParams) SetReference(reference *string) {
	o.Reference = reference
}

// WriteToRequest writes these params to a swagger request
func (o *GetDocumentPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DocumentID != nil {

		// query param document_id
		var qrDocumentID string
		if o.DocumentID != nil {
			qrDocumentID = *o.DocumentID
		}
		qDocumentID := qrDocumentID
		if qDocumentID != "" {
			if err := r.SetQueryParam("document_id", qDocumentID); err != nil {
				return err
			}
		}

	}

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

	if o.PaymentAt != nil {

		// query param payment_at
		var qrPaymentAt string
		if o.PaymentAt != nil {
			qrPaymentAt = *o.PaymentAt
		}
		qPaymentAt := qrPaymentAt
		if qPaymentAt != "" {
			if err := r.SetQueryParam("payment_at", qPaymentAt); err != nil {
				return err
			}
		}

	}

	if o.Reference != nil {

		// query param reference
		var qrReference string
		if o.Reference != nil {
			qrReference = *o.Reference
		}
		qReference := qrReference
		if qReference != "" {
			if err := r.SetQueryParam("reference", qReference); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
