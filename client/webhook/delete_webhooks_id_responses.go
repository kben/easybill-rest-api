// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteWebhooksIDReader is a Reader for the DeleteWebhooksID structure.
type DeleteWebhooksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebhooksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWebhooksIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteWebhooksIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWebhooksIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteWebhooksIDNoContent creates a DeleteWebhooksIDNoContent with default headers values
func NewDeleteWebhooksIDNoContent() *DeleteWebhooksIDNoContent {
	return &DeleteWebhooksIDNoContent{}
}

/*DeleteWebhooksIDNoContent handles this case with default header values.

Successful operation
*/
type DeleteWebhooksIDNoContent struct {
}

func (o *DeleteWebhooksIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{id}][%d] deleteWebhooksIdNoContent ", 204)
}

func (o *DeleteWebhooksIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebhooksIDNotFound creates a DeleteWebhooksIDNotFound with default headers values
func NewDeleteWebhooksIDNotFound() *DeleteWebhooksIDNotFound {
	return &DeleteWebhooksIDNotFound{}
}

/*DeleteWebhooksIDNotFound handles this case with default header values.

Not found
*/
type DeleteWebhooksIDNotFound struct {
}

func (o *DeleteWebhooksIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{id}][%d] deleteWebhooksIdNotFound ", 404)
}

func (o *DeleteWebhooksIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebhooksIDTooManyRequests creates a DeleteWebhooksIDTooManyRequests with default headers values
func NewDeleteWebhooksIDTooManyRequests() *DeleteWebhooksIDTooManyRequests {
	return &DeleteWebhooksIDTooManyRequests{}
}

/*DeleteWebhooksIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type DeleteWebhooksIDTooManyRequests struct {
}

func (o *DeleteWebhooksIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{id}][%d] deleteWebhooksIdTooManyRequests ", 429)
}

func (o *DeleteWebhooksIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}