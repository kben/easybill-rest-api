// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectsIDReader is a Reader for the DeleteProjectsID structure.
type DeleteProjectsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteProjectsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteProjectsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProjectsIDNoContent creates a DeleteProjectsIDNoContent with default headers values
func NewDeleteProjectsIDNoContent() *DeleteProjectsIDNoContent {
	return &DeleteProjectsIDNoContent{}
}

/*DeleteProjectsIDNoContent handles this case with default header values.

Successful operation
*/
type DeleteProjectsIDNoContent struct {
}

func (o *DeleteProjectsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /projects/{id}][%d] deleteProjectsIdNoContent ", 204)
}

func (o *DeleteProjectsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsIDNotFound creates a DeleteProjectsIDNotFound with default headers values
func NewDeleteProjectsIDNotFound() *DeleteProjectsIDNotFound {
	return &DeleteProjectsIDNotFound{}
}

/*DeleteProjectsIDNotFound handles this case with default header values.

Not found
*/
type DeleteProjectsIDNotFound struct {
}

func (o *DeleteProjectsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{id}][%d] deleteProjectsIdNotFound ", 404)
}

func (o *DeleteProjectsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsIDTooManyRequests creates a DeleteProjectsIDTooManyRequests with default headers values
func NewDeleteProjectsIDTooManyRequests() *DeleteProjectsIDTooManyRequests {
	return &DeleteProjectsIDTooManyRequests{}
}

/*DeleteProjectsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type DeleteProjectsIDTooManyRequests struct {
}

func (o *DeleteProjectsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /projects/{id}][%d] deleteProjectsIdTooManyRequests ", 429)
}

func (o *DeleteProjectsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}