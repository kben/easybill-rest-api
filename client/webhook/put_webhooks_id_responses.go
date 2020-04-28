// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// PutWebhooksIDReader is a Reader for the PutWebhooksID structure.
type PutWebhooksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWebhooksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutWebhooksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPutWebhooksIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutWebhooksIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutWebhooksIDOK creates a PutWebhooksIDOK with default headers values
func NewPutWebhooksIDOK() *PutWebhooksIDOK {
	return &PutWebhooksIDOK{}
}

/*PutWebhooksIDOK handles this case with default header values.

Successful operation
*/
type PutWebhooksIDOK struct {
	Payload *models.WebHook
}

func (o *PutWebhooksIDOK) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdOK  %+v", 200, o.Payload)
}

func (o *PutWebhooksIDOK) GetPayload() *models.WebHook {
	return o.Payload
}

func (o *PutWebhooksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebhooksIDNotFound creates a PutWebhooksIDNotFound with default headers values
func NewPutWebhooksIDNotFound() *PutWebhooksIDNotFound {
	return &PutWebhooksIDNotFound{}
}

/*PutWebhooksIDNotFound handles this case with default header values.

Not found
*/
type PutWebhooksIDNotFound struct {
}

func (o *PutWebhooksIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdNotFound ", 404)
}

func (o *PutWebhooksIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWebhooksIDTooManyRequests creates a PutWebhooksIDTooManyRequests with default headers values
func NewPutWebhooksIDTooManyRequests() *PutWebhooksIDTooManyRequests {
	return &PutWebhooksIDTooManyRequests{}
}

/*PutWebhooksIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutWebhooksIDTooManyRequests struct {
}

func (o *PutWebhooksIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{id}][%d] putWebhooksIdTooManyRequests ", 429)
}

func (o *PutWebhooksIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
