// Code generated by go-swagger; DO NOT EDIT.

package document_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDocumentPaymentsIDReader is a Reader for the DeleteDocumentPaymentsID structure.
type DeleteDocumentPaymentsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDocumentPaymentsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDocumentPaymentsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteDocumentPaymentsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteDocumentPaymentsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDocumentPaymentsIDNoContent creates a DeleteDocumentPaymentsIDNoContent with default headers values
func NewDeleteDocumentPaymentsIDNoContent() *DeleteDocumentPaymentsIDNoContent {
	return &DeleteDocumentPaymentsIDNoContent{}
}

/*DeleteDocumentPaymentsIDNoContent handles this case with default header values.

Successful operation
*/
type DeleteDocumentPaymentsIDNoContent struct {
}

func (o *DeleteDocumentPaymentsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /document-payments/{id}][%d] deleteDocumentPaymentsIdNoContent ", 204)
}

func (o *DeleteDocumentPaymentsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDocumentPaymentsIDNotFound creates a DeleteDocumentPaymentsIDNotFound with default headers values
func NewDeleteDocumentPaymentsIDNotFound() *DeleteDocumentPaymentsIDNotFound {
	return &DeleteDocumentPaymentsIDNotFound{}
}

/*DeleteDocumentPaymentsIDNotFound handles this case with default header values.

Not found
*/
type DeleteDocumentPaymentsIDNotFound struct {
}

func (o *DeleteDocumentPaymentsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /document-payments/{id}][%d] deleteDocumentPaymentsIdNotFound ", 404)
}

func (o *DeleteDocumentPaymentsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDocumentPaymentsIDTooManyRequests creates a DeleteDocumentPaymentsIDTooManyRequests with default headers values
func NewDeleteDocumentPaymentsIDTooManyRequests() *DeleteDocumentPaymentsIDTooManyRequests {
	return &DeleteDocumentPaymentsIDTooManyRequests{}
}

/*DeleteDocumentPaymentsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type DeleteDocumentPaymentsIDTooManyRequests struct {
}

func (o *DeleteDocumentPaymentsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /document-payments/{id}][%d] deleteDocumentPaymentsIdTooManyRequests ", 429)
}

func (o *DeleteDocumentPaymentsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
