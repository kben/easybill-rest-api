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

// GetCustomerGroupsIDReader is a Reader for the GetCustomerGroupsID structure.
type GetCustomerGroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerGroupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerGroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCustomerGroupsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCustomerGroupsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomerGroupsIDOK creates a GetCustomerGroupsIDOK with default headers values
func NewGetCustomerGroupsIDOK() *GetCustomerGroupsIDOK {
	return &GetCustomerGroupsIDOK{}
}

/*GetCustomerGroupsIDOK handles this case with default header values.

Successful operation
*/
type GetCustomerGroupsIDOK struct {
	Payload *models.CustomerGroup
}

func (o *GetCustomerGroupsIDOK) Error() string {
	return fmt.Sprintf("[GET /customer-groups/{id}][%d] getCustomerGroupsIdOK  %+v", 200, o.Payload)
}

func (o *GetCustomerGroupsIDOK) GetPayload() *models.CustomerGroup {
	return o.Payload
}

func (o *GetCustomerGroupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerGroupsIDNotFound creates a GetCustomerGroupsIDNotFound with default headers values
func NewGetCustomerGroupsIDNotFound() *GetCustomerGroupsIDNotFound {
	return &GetCustomerGroupsIDNotFound{}
}

/*GetCustomerGroupsIDNotFound handles this case with default header values.

Not found
*/
type GetCustomerGroupsIDNotFound struct {
}

func (o *GetCustomerGroupsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /customer-groups/{id}][%d] getCustomerGroupsIdNotFound ", 404)
}

func (o *GetCustomerGroupsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomerGroupsIDTooManyRequests creates a GetCustomerGroupsIDTooManyRequests with default headers values
func NewGetCustomerGroupsIDTooManyRequests() *GetCustomerGroupsIDTooManyRequests {
	return &GetCustomerGroupsIDTooManyRequests{}
}

/*GetCustomerGroupsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetCustomerGroupsIDTooManyRequests struct {
}

func (o *GetCustomerGroupsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /customer-groups/{id}][%d] getCustomerGroupsIdTooManyRequests ", 429)
}

func (o *GetCustomerGroupsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
