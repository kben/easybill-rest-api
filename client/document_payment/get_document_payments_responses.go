// Code generated by go-swagger; DO NOT EDIT.

package document_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// GetDocumentPaymentsReader is a Reader for the GetDocumentPayments structure.
type GetDocumentPaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDocumentPaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDocumentPaymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDocumentPaymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDocumentPaymentsOK creates a GetDocumentPaymentsOK with default headers values
func NewGetDocumentPaymentsOK() *GetDocumentPaymentsOK {
	return &GetDocumentPaymentsOK{}
}

/*GetDocumentPaymentsOK handles this case with default header values.

Successful operation
*/
type GetDocumentPaymentsOK struct {
	Payload *models.DocumentPayments
}

func (o *GetDocumentPaymentsOK) Error() string {
	return fmt.Sprintf("[GET /document-payments][%d] getDocumentPaymentsOK  %+v", 200, o.Payload)
}

func (o *GetDocumentPaymentsOK) GetPayload() *models.DocumentPayments {
	return o.Payload
}

func (o *GetDocumentPaymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentPayments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentPaymentsTooManyRequests creates a GetDocumentPaymentsTooManyRequests with default headers values
func NewGetDocumentPaymentsTooManyRequests() *GetDocumentPaymentsTooManyRequests {
	return &GetDocumentPaymentsTooManyRequests{}
}

/*GetDocumentPaymentsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetDocumentPaymentsTooManyRequests struct {
}

func (o *GetDocumentPaymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /document-payments][%d] getDocumentPaymentsTooManyRequests ", 429)
}

func (o *GetDocumentPaymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
