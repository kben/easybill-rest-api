// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// GetCustomersCustomerIDContactsReader is a Reader for the GetCustomersCustomerIDContacts structure.
type GetCustomersCustomerIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerIDContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetCustomersCustomerIDContactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomersCustomerIDContactsOK creates a GetCustomersCustomerIDContactsOK with default headers values
func NewGetCustomersCustomerIDContactsOK() *GetCustomersCustomerIDContactsOK {
	return &GetCustomersCustomerIDContactsOK{}
}

/*GetCustomersCustomerIDContactsOK handles this case with default header values.

Successful operation
*/
type GetCustomersCustomerIDContactsOK struct {
	Payload *models.Contacts
}

func (o *GetCustomersCustomerIDContactsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerId}/contacts][%d] getCustomersCustomerIdContactsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerIDContactsOK) GetPayload() *models.Contacts {
	return o.Payload
}

func (o *GetCustomersCustomerIDContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contacts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerIDContactsTooManyRequests creates a GetCustomersCustomerIDContactsTooManyRequests with default headers values
func NewGetCustomersCustomerIDContactsTooManyRequests() *GetCustomersCustomerIDContactsTooManyRequests {
	return &GetCustomersCustomerIDContactsTooManyRequests{}
}

/*GetCustomersCustomerIDContactsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetCustomersCustomerIDContactsTooManyRequests struct {
}

func (o *GetCustomersCustomerIDContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /customers/{customerId}/contacts][%d] getCustomersCustomerIdContactsTooManyRequests ", 429)
}

func (o *GetCustomersCustomerIDContactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}