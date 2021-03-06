// Code generated by go-swagger; DO NOT EDIT.

package text_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteTextTemplatesIDReader is a Reader for the DeleteTextTemplatesID structure.
type DeleteTextTemplatesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTextTemplatesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTextTemplatesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTextTemplatesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTextTemplatesIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTextTemplatesIDNoContent creates a DeleteTextTemplatesIDNoContent with default headers values
func NewDeleteTextTemplatesIDNoContent() *DeleteTextTemplatesIDNoContent {
	return &DeleteTextTemplatesIDNoContent{}
}

/*DeleteTextTemplatesIDNoContent handles this case with default header values.

Successful operation
*/
type DeleteTextTemplatesIDNoContent struct {
}

func (o *DeleteTextTemplatesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /text-templates/{id}][%d] deleteTextTemplatesIdNoContent ", 204)
}

func (o *DeleteTextTemplatesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTextTemplatesIDNotFound creates a DeleteTextTemplatesIDNotFound with default headers values
func NewDeleteTextTemplatesIDNotFound() *DeleteTextTemplatesIDNotFound {
	return &DeleteTextTemplatesIDNotFound{}
}

/*DeleteTextTemplatesIDNotFound handles this case with default header values.

Not found
*/
type DeleteTextTemplatesIDNotFound struct {
}

func (o *DeleteTextTemplatesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /text-templates/{id}][%d] deleteTextTemplatesIdNotFound ", 404)
}

func (o *DeleteTextTemplatesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTextTemplatesIDTooManyRequests creates a DeleteTextTemplatesIDTooManyRequests with default headers values
func NewDeleteTextTemplatesIDTooManyRequests() *DeleteTextTemplatesIDTooManyRequests {
	return &DeleteTextTemplatesIDTooManyRequests{}
}

/*DeleteTextTemplatesIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type DeleteTextTemplatesIDTooManyRequests struct {
}

func (o *DeleteTextTemplatesIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /text-templates/{id}][%d] deleteTextTemplatesIdTooManyRequests ", 429)
}

func (o *DeleteTextTemplatesIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
