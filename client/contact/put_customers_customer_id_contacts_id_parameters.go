// Code generated by go-swagger; DO NOT EDIT.

package contact

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

	"github.com/kben/easybill-rest-api/models"
)

// NewPutCustomersCustomerIDContactsIDParams creates a new PutCustomersCustomerIDContactsIDParams object
// with the default values initialized.
func NewPutCustomersCustomerIDContactsIDParams() *PutCustomersCustomerIDContactsIDParams {
	var ()
	return &PutCustomersCustomerIDContactsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerIDContactsIDParamsWithTimeout creates a new PutCustomersCustomerIDContactsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerIDContactsIDParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerIDContactsIDParams {
	var ()
	return &PutCustomersCustomerIDContactsIDParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerIDContactsIDParamsWithContext creates a new PutCustomersCustomerIDContactsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerIDContactsIDParamsWithContext(ctx context.Context) *PutCustomersCustomerIDContactsIDParams {
	var ()
	return &PutCustomersCustomerIDContactsIDParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerIDContactsIDParamsWithHTTPClient creates a new PutCustomersCustomerIDContactsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerIDContactsIDParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerIDContactsIDParams {
	var ()
	return &PutCustomersCustomerIDContactsIDParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerIDContactsIDParams contains all the parameters to send to the API endpoint
for the put customers customer ID contacts ID operation typically these are written to a http.Request
*/
type PutCustomersCustomerIDContactsIDParams struct {

	/*Body*/
	Body *models.Contact
	/*CustomerID
	  ID of customer

	*/
	CustomerID int64
	/*ID
	  ID of contact

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerIDContactsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) WithContext(ctx context.Context) *PutCustomersCustomerIDContactsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerIDContactsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) WithBody(body *models.Contact) *PutCustomersCustomerIDContactsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) SetBody(body *models.Contact) {
	o.Body = body
}

// WithCustomerID adds the customerID to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) WithCustomerID(customerID int64) *PutCustomersCustomerIDContactsIDParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithID adds the id to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) WithID(id int64) *PutCustomersCustomerIDContactsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put customers customer ID contacts ID params
func (o *PutCustomersCustomerIDContactsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerIDContactsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}