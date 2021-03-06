// Code generated by go-swagger; DO NOT EDIT.

package sepa_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// PutSepaPaymentsIDReader is a Reader for the PutSepaPaymentsID structure.
type PutSepaPaymentsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSepaPaymentsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSepaPaymentsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutSepaPaymentsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutSepaPaymentsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutSepaPaymentsIDOK creates a PutSepaPaymentsIDOK with default headers values
func NewPutSepaPaymentsIDOK() *PutSepaPaymentsIDOK {
	return &PutSepaPaymentsIDOK{}
}

/*PutSepaPaymentsIDOK handles this case with default header values.

Successful operation
*/
type PutSepaPaymentsIDOK struct {
	Payload *models.SEPAPayment
}

func (o *PutSepaPaymentsIDOK) Error() string {
	return fmt.Sprintf("[PUT /sepa-payments/{id}][%d] putSepaPaymentsIdOK  %+v", 200, o.Payload)
}

func (o *PutSepaPaymentsIDOK) GetPayload() *models.SEPAPayment {
	return o.Payload
}

func (o *PutSepaPaymentsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SEPAPayment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSepaPaymentsIDBadRequest creates a PutSepaPaymentsIDBadRequest with default headers values
func NewPutSepaPaymentsIDBadRequest() *PutSepaPaymentsIDBadRequest {
	return &PutSepaPaymentsIDBadRequest{}
}

/*PutSepaPaymentsIDBadRequest handles this case with default header values.

Invalid SEPA payment
*/
type PutSepaPaymentsIDBadRequest struct {
}

func (o *PutSepaPaymentsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sepa-payments/{id}][%d] putSepaPaymentsIdBadRequest ", 400)
}

func (o *PutSepaPaymentsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSepaPaymentsIDTooManyRequests creates a PutSepaPaymentsIDTooManyRequests with default headers values
func NewPutSepaPaymentsIDTooManyRequests() *PutSepaPaymentsIDTooManyRequests {
	return &PutSepaPaymentsIDTooManyRequests{}
}

/*PutSepaPaymentsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutSepaPaymentsIDTooManyRequests struct {
}

func (o *PutSepaPaymentsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /sepa-payments/{id}][%d] putSepaPaymentsIdTooManyRequests ", 429)
}

func (o *PutSepaPaymentsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
