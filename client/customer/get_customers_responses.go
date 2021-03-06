// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// GetCustomersReader is a Reader for the GetCustomers structure.
type GetCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetCustomersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomersOK creates a GetCustomersOK with default headers values
func NewGetCustomersOK() *GetCustomersOK {
	return &GetCustomersOK{}
}

/*GetCustomersOK handles this case with default header values.

Successful operation
*/
type GetCustomersOK struct {
	Payload *models.Customers
}

func (o *GetCustomersOK) Error() string {
	return fmt.Sprintf("[GET /customers][%d] getCustomersOK  %+v", 200, o.Payload)
}

func (o *GetCustomersOK) GetPayload() *models.Customers {
	return o.Payload
}

func (o *GetCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersTooManyRequests creates a GetCustomersTooManyRequests with default headers values
func NewGetCustomersTooManyRequests() *GetCustomersTooManyRequests {
	return &GetCustomersTooManyRequests{}
}

/*GetCustomersTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetCustomersTooManyRequests struct {
}

func (o *GetCustomersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /customers][%d] getCustomersTooManyRequests ", 429)
}

func (o *GetCustomersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
