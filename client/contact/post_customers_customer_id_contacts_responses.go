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

// PostCustomersCustomerIDContactsReader is a Reader for the PostCustomersCustomerIDContacts structure.
type PostCustomersCustomerIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerIDContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCustomersCustomerIDContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostCustomersCustomerIDContactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCustomersCustomerIDContactsOK creates a PostCustomersCustomerIDContactsOK with default headers values
func NewPostCustomersCustomerIDContactsOK() *PostCustomersCustomerIDContactsOK {
	return &PostCustomersCustomerIDContactsOK{}
}

/*PostCustomersCustomerIDContactsOK handles this case with default header values.

Successful operation
*/
type PostCustomersCustomerIDContactsOK struct {
	Payload *models.Contact
}

func (o *PostCustomersCustomerIDContactsOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerId}/contacts][%d] postCustomersCustomerIdContactsOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerIDContactsOK) GetPayload() *models.Contact {
	return o.Payload
}

func (o *PostCustomersCustomerIDContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerIDContactsBadRequest creates a PostCustomersCustomerIDContactsBadRequest with default headers values
func NewPostCustomersCustomerIDContactsBadRequest() *PostCustomersCustomerIDContactsBadRequest {
	return &PostCustomersCustomerIDContactsBadRequest{}
}

/*PostCustomersCustomerIDContactsBadRequest handles this case with default header values.

Invalid contact
*/
type PostCustomersCustomerIDContactsBadRequest struct {
}

func (o *PostCustomersCustomerIDContactsBadRequest) Error() string {
	return fmt.Sprintf("[POST /customers/{customerId}/contacts][%d] postCustomersCustomerIdContactsBadRequest ", 400)
}

func (o *PostCustomersCustomerIDContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersCustomerIDContactsTooManyRequests creates a PostCustomersCustomerIDContactsTooManyRequests with default headers values
func NewPostCustomersCustomerIDContactsTooManyRequests() *PostCustomersCustomerIDContactsTooManyRequests {
	return &PostCustomersCustomerIDContactsTooManyRequests{}
}

/*PostCustomersCustomerIDContactsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PostCustomersCustomerIDContactsTooManyRequests struct {
}

func (o *PostCustomersCustomerIDContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /customers/{customerId}/contacts][%d] postCustomersCustomerIdContactsTooManyRequests ", 429)
}

func (o *PostCustomersCustomerIDContactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}