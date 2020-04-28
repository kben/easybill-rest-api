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

// PostDocumentsReader is a Reader for the PostDocuments structure.
type PostDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostDocumentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPostDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDocumentsCreated creates a PostDocumentsCreated with default headers values
func NewPostDocumentsCreated() *PostDocumentsCreated {
	return &PostDocumentsCreated{}
}

/*PostDocumentsCreated handles this case with default header values.

Successful operation
*/
type PostDocumentsCreated struct {
	Payload *models.Document
}

func (o *PostDocumentsCreated) Error() string {
	return fmt.Sprintf("[POST /documents][%d] postDocumentsCreated  %+v", 201, o.Payload)
}

func (o *PostDocumentsCreated) GetPayload() *models.Document {
	return o.Payload
}

func (o *PostDocumentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Document)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentsTooManyRequests creates a PostDocumentsTooManyRequests with default headers values
func NewPostDocumentsTooManyRequests() *PostDocumentsTooManyRequests {
	return &PostDocumentsTooManyRequests{}
}

/*PostDocumentsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PostDocumentsTooManyRequests struct {
}

func (o *PostDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /documents][%d] postDocumentsTooManyRequests ", 429)
}

func (o *PostDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
