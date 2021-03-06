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

// PutCustomersIDReader is a Reader for the PutCustomersID structure.
type PutCustomersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCustomersIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutCustomersIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersIDOK creates a PutCustomersIDOK with default headers values
func NewPutCustomersIDOK() *PutCustomersIDOK {
	return &PutCustomersIDOK{}
}

/*PutCustomersIDOK handles this case with default header values.

Successful operation
*/
type PutCustomersIDOK struct {
	Payload *models.Customer
}

func (o *PutCustomersIDOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{id}][%d] putCustomersIdOK  %+v", 200, o.Payload)
}

func (o *PutCustomersIDOK) GetPayload() *models.Customer {
	return o.Payload
}

func (o *PutCustomersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersIDBadRequest creates a PutCustomersIDBadRequest with default headers values
func NewPutCustomersIDBadRequest() *PutCustomersIDBadRequest {
	return &PutCustomersIDBadRequest{}
}

/*PutCustomersIDBadRequest handles this case with default header values.

Invalid Customer
*/
type PutCustomersIDBadRequest struct {
}

func (o *PutCustomersIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /customers/{id}][%d] putCustomersIdBadRequest ", 400)
}

func (o *PutCustomersIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersIDTooManyRequests creates a PutCustomersIDTooManyRequests with default headers values
func NewPutCustomersIDTooManyRequests() *PutCustomersIDTooManyRequests {
	return &PutCustomersIDTooManyRequests{}
}

/*PutCustomersIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutCustomersIDTooManyRequests struct {
}

func (o *PutCustomersIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /customers/{id}][%d] putCustomersIdTooManyRequests ", 429)
}

func (o *PutCustomersIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
