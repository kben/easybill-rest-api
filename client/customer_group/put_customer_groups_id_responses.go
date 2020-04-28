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

// PutCustomerGroupsIDReader is a Reader for the PutCustomerGroupsID structure.
type PutCustomerGroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomerGroupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomerGroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCustomerGroupsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutCustomerGroupsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomerGroupsIDOK creates a PutCustomerGroupsIDOK with default headers values
func NewPutCustomerGroupsIDOK() *PutCustomerGroupsIDOK {
	return &PutCustomerGroupsIDOK{}
}

/*PutCustomerGroupsIDOK handles this case with default header values.

Successful operation
*/
type PutCustomerGroupsIDOK struct {
	Payload *models.CustomerGroup
}

func (o *PutCustomerGroupsIDOK) Error() string {
	return fmt.Sprintf("[PUT /customer-groups/{id}][%d] putCustomerGroupsIdOK  %+v", 200, o.Payload)
}

func (o *PutCustomerGroupsIDOK) GetPayload() *models.CustomerGroup {
	return o.Payload
}

func (o *PutCustomerGroupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomerGroupsIDBadRequest creates a PutCustomerGroupsIDBadRequest with default headers values
func NewPutCustomerGroupsIDBadRequest() *PutCustomerGroupsIDBadRequest {
	return &PutCustomerGroupsIDBadRequest{}
}

/*PutCustomerGroupsIDBadRequest handles this case with default header values.

Invalid customer group
*/
type PutCustomerGroupsIDBadRequest struct {
}

func (o *PutCustomerGroupsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /customer-groups/{id}][%d] putCustomerGroupsIdBadRequest ", 400)
}

func (o *PutCustomerGroupsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomerGroupsIDTooManyRequests creates a PutCustomerGroupsIDTooManyRequests with default headers values
func NewPutCustomerGroupsIDTooManyRequests() *PutCustomerGroupsIDTooManyRequests {
	return &PutCustomerGroupsIDTooManyRequests{}
}

/*PutCustomerGroupsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutCustomerGroupsIDTooManyRequests struct {
}

func (o *PutCustomerGroupsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /customer-groups/{id}][%d] putCustomerGroupsIdTooManyRequests ", 429)
}

func (o *PutCustomerGroupsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
