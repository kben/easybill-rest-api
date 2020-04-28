// Code generated by go-swagger; DO NOT EDIT.

package time_tracking

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

// NewGetTimeTrackingsParams creates a new GetTimeTrackingsParams object
// with the default values initialized.
func NewGetTimeTrackingsParams() *GetTimeTrackingsParams {
	var ()
	return &GetTimeTrackingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTimeTrackingsParamsWithTimeout creates a new GetTimeTrackingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTimeTrackingsParamsWithTimeout(timeout time.Duration) *GetTimeTrackingsParams {
	var ()
	return &GetTimeTrackingsParams{

		timeout: timeout,
	}
}

// NewGetTimeTrackingsParamsWithContext creates a new GetTimeTrackingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTimeTrackingsParamsWithContext(ctx context.Context) *GetTimeTrackingsParams {
	var ()
	return &GetTimeTrackingsParams{

		Context: ctx,
	}
}

// NewGetTimeTrackingsParamsWithHTTPClient creates a new GetTimeTrackingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTimeTrackingsParamsWithHTTPClient(client *http.Client) *GetTimeTrackingsParams {
	var ()
	return &GetTimeTrackingsParams{
		HTTPClient: client,
	}
}

/*GetTimeTrackingsParams contains all the parameters to send to the API endpoint
for the get time trackings operation typically these are written to a http.Request
*/
type GetTimeTrackingsParams struct {

	/*Limit
	  Limited the result. Default is 100. Maximum can be 1000.

	*/
	Limit *int64
	/*LoginID
	  Filter time-tracking by login_id. You can add multiple ids separate by comma like id,id,id.

	*/
	LoginID *string
	/*Page
	  Set current Page. Default is 1.

	*/
	Page *int64
	/*ProjectID
	  Filter time-tracking by project_id. You can add multiple ids separate by comma like id,id,id.

	*/
	ProjectID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get time trackings params
func (o *GetTimeTrackingsParams) WithTimeout(timeout time.Duration) *GetTimeTrackingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get time trackings params
func (o *GetTimeTrackingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get time trackings params
func (o *GetTimeTrackingsParams) WithContext(ctx context.Context) *GetTimeTrackingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get time trackings params
func (o *GetTimeTrackingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get time trackings params
func (o *GetTimeTrackingsParams) WithHTTPClient(client *http.Client) *GetTimeTrackingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get time trackings params
func (o *GetTimeTrackingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get time trackings params
func (o *GetTimeTrackingsParams) WithLimit(limit *int64) *GetTimeTrackingsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get time trackings params
func (o *GetTimeTrackingsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLoginID adds the loginID to the get time trackings params
func (o *GetTimeTrackingsParams) WithLoginID(loginID *string) *GetTimeTrackingsParams {
	o.SetLoginID(loginID)
	return o
}

// SetLoginID adds the loginId to the get time trackings params
func (o *GetTimeTrackingsParams) SetLoginID(loginID *string) {
	o.LoginID = loginID
}

// WithPage adds the page to the get time trackings params
func (o *GetTimeTrackingsParams) WithPage(page *int64) *GetTimeTrackingsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get time trackings params
func (o *GetTimeTrackingsParams) SetPage(page *int64) {
	o.Page = page
}

// WithProjectID adds the projectID to the get time trackings params
func (o *GetTimeTrackingsParams) WithProjectID(projectID *string) *GetTimeTrackingsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get time trackings params
func (o *GetTimeTrackingsParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTimeTrackingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LoginID != nil {

		// query param login_id
		var qrLoginID string
		if o.LoginID != nil {
			qrLoginID = *o.LoginID
		}
		qLoginID := qrLoginID
		if qLoginID != "" {
			if err := r.SetQueryParam("login_id", qLoginID); err != nil {
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

	if o.ProjectID != nil {

		// query param project_id
		var qrProjectID string
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {
			if err := r.SetQueryParam("project_id", qProjectID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
