// Code generated by go-swagger; DO NOT EDIT.

package customer_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// GetCustomerGroupsReader is a Reader for the GetCustomerGroups structure.
type GetCustomerGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetCustomerGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomerGroupsOK creates a GetCustomerGroupsOK with default headers values
func NewGetCustomerGroupsOK() *GetCustomerGroupsOK {
	return &GetCustomerGroupsOK{}
}

/*GetCustomerGroupsOK handles this case with default header values.

Successful operation
*/
type GetCustomerGroupsOK struct {
	Payload *models.CustomerGroups
}

func (o *GetCustomerGroupsOK) Error() string {
	return fmt.Sprintf("[GET /customer-groups][%d] getCustomerGroupsOK  %+v", 200, o.Payload)
}

func (o *GetCustomerGroupsOK) GetPayload() *models.CustomerGroups {
	return o.Payload
}

func (o *GetCustomerGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerGroups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerGroupsTooManyRequests creates a GetCustomerGroupsTooManyRequests with default headers values
func NewGetCustomerGroupsTooManyRequests() *GetCustomerGroupsTooManyRequests {
	return &GetCustomerGroupsTooManyRequests{}
}

/*GetCustomerGroupsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetCustomerGroupsTooManyRequests struct {
}

func (o *GetCustomerGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /customer-groups][%d] getCustomerGroupsTooManyRequests ", 429)
}

func (o *GetCustomerGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}