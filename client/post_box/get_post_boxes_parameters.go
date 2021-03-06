// Code generated by go-swagger; DO NOT EDIT.

package post_box

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

// NewGetPostBoxesParams creates a new GetPostBoxesParams object
// with the default values initialized.
func NewGetPostBoxesParams() *GetPostBoxesParams {
	var ()
	return &GetPostBoxesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPostBoxesParamsWithTimeout creates a new GetPostBoxesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPostBoxesParamsWithTimeout(timeout time.Duration) *GetPostBoxesParams {
	var ()
	return &GetPostBoxesParams{

		timeout: timeout,
	}
}

// NewGetPostBoxesParamsWithContext creates a new GetPostBoxesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPostBoxesParamsWithContext(ctx context.Context) *GetPostBoxesParams {
	var ()
	return &GetPostBoxesParams{

		Context: ctx,
	}
}

// NewGetPostBoxesParamsWithHTTPClient creates a new GetPostBoxesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPostBoxesParamsWithHTTPClient(client *http.Client) *GetPostBoxesParams {
	var ()
	return &GetPostBoxesParams{
		HTTPClient: client,
	}
}

/*GetPostBoxesParams contains all the parameters to send to the API endpoint
for the get post boxes operation typically these are written to a http.Request
*/
type GetPostBoxesParams struct {

	/*DocumentID
	  Filter post boxes by document_id. You can add multiple document ids separate by comma like id,id,id.

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
	/*Status
	  Filter post boxes by status.

	*/
	Status *string
	/*Type
	  Filter post boxes by type. Multiple typs seperate with , like type=EMAIL,FAX.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get post boxes params
func (o *GetPostBoxesParams) WithTimeout(timeout time.Duration) *GetPostBoxesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get post boxes params
func (o *GetPostBoxesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get post boxes params
func (o *GetPostBoxesParams) WithContext(ctx context.Context) *GetPostBoxesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get post boxes params
func (o *GetPostBoxesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get post boxes params
func (o *GetPostBoxesParams) WithHTTPClient(client *http.Client) *GetPostBoxesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get post boxes params
func (o *GetPostBoxesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get post boxes params
func (o *GetPostBoxesParams) WithDocumentID(documentID *string) *GetPostBoxesParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get post boxes params
func (o *GetPostBoxesParams) SetDocumentID(documentID *string) {
	o.DocumentID = documentID
}

// WithLimit adds the limit to the get post boxes params
func (o *GetPostBoxesParams) WithLimit(limit *int64) *GetPostBoxesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get post boxes params
func (o *GetPostBoxesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get post boxes params
func (o *GetPostBoxesParams) WithPage(page *int64) *GetPostBoxesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get post boxes params
func (o *GetPostBoxesParams) SetPage(page *int64) {
	o.Page = page
}

// WithStatus adds the status to the get post boxes params
func (o *GetPostBoxesParams) WithStatus(status *string) *GetPostBoxesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get post boxes params
func (o *GetPostBoxesParams) SetStatus(status *string) {
	o.Status = status
}

// WithType adds the typeVar to the get post boxes params
func (o *GetPostBoxesParams) WithType(typeVar *string) *GetPostBoxesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get post boxes params
func (o *GetPostBoxesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPostBoxesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
