// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// PutDocumentsIDDoneReader is a Reader for the PutDocumentsIDDone structure.
type PutDocumentsIDDoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDocumentsIDDoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDocumentsIDDoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPutDocumentsIDDoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutDocumentsIDDoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDocumentsIDDoneOK creates a PutDocumentsIDDoneOK with default headers values
func NewPutDocumentsIDDoneOK() *PutDocumentsIDDoneOK {
	return &PutDocumentsIDDoneOK{}
}

/*PutDocumentsIDDoneOK handles this case with default header values.

Successful operation
*/
type PutDocumentsIDDoneOK struct {
	Payload *models.Document
}

func (o *PutDocumentsIDDoneOK) Error() string {
	return fmt.Sprintf("[PUT /documents/{id}/done][%d] putDocumentsIdDoneOK  %+v", 200, o.Payload)
}

func (o *PutDocumentsIDDoneOK) GetPayload() *models.Document {
	return o.Payload
}

func (o *PutDocumentsIDDoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Document)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDocumentsIDDoneNotFound creates a PutDocumentsIDDoneNotFound with default headers values
func NewPutDocumentsIDDoneNotFound() *PutDocumentsIDDoneNotFound {
	return &PutDocumentsIDDoneNotFound{}
}

/*PutDocumentsIDDoneNotFound handles this case with default header values.

Not found
*/
type PutDocumentsIDDoneNotFound struct {
}

func (o *PutDocumentsIDDoneNotFound) Error() string {
	return fmt.Sprintf("[PUT /documents/{id}/done][%d] putDocumentsIdDoneNotFound ", 404)
}

func (o *PutDocumentsIDDoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDocumentsIDDoneTooManyRequests creates a PutDocumentsIDDoneTooManyRequests with default headers values
func NewPutDocumentsIDDoneTooManyRequests() *PutDocumentsIDDoneTooManyRequests {
	return &PutDocumentsIDDoneTooManyRequests{}
}

/*PutDocumentsIDDoneTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutDocumentsIDDoneTooManyRequests struct {
}

func (o *PutDocumentsIDDoneTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /documents/{id}/done][%d] putDocumentsIdDoneTooManyRequests ", 429)
}

func (o *PutDocumentsIDDoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
