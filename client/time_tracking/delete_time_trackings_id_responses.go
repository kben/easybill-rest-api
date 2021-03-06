// Code generated by go-swagger; DO NOT EDIT.

package time_tracking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteTimeTrackingsIDReader is a Reader for the DeleteTimeTrackingsID structure.
type DeleteTimeTrackingsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTimeTrackingsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTimeTrackingsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTimeTrackingsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTimeTrackingsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTimeTrackingsIDNoContent creates a DeleteTimeTrackingsIDNoContent with default headers values
func NewDeleteTimeTrackingsIDNoContent() *DeleteTimeTrackingsIDNoContent {
	return &DeleteTimeTrackingsIDNoContent{}
}

/*DeleteTimeTrackingsIDNoContent handles this case with default header values.

Successful operation
*/
type DeleteTimeTrackingsIDNoContent struct {
}

func (o *DeleteTimeTrackingsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /time-trackings/{id}][%d] deleteTimeTrackingsIdNoContent ", 204)
}

func (o *DeleteTimeTrackingsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTimeTrackingsIDNotFound creates a DeleteTimeTrackingsIDNotFound with default headers values
func NewDeleteTimeTrackingsIDNotFound() *DeleteTimeTrackingsIDNotFound {
	return &DeleteTimeTrackingsIDNotFound{}
}

/*DeleteTimeTrackingsIDNotFound handles this case with default header values.

Not found
*/
type DeleteTimeTrackingsIDNotFound struct {
}

func (o *DeleteTimeTrackingsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /time-trackings/{id}][%d] deleteTimeTrackingsIdNotFound ", 404)
}

func (o *DeleteTimeTrackingsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTimeTrackingsIDTooManyRequests creates a DeleteTimeTrackingsIDTooManyRequests with default headers values
func NewDeleteTimeTrackingsIDTooManyRequests() *DeleteTimeTrackingsIDTooManyRequests {
	return &DeleteTimeTrackingsIDTooManyRequests{}
}

/*DeleteTimeTrackingsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type DeleteTimeTrackingsIDTooManyRequests struct {
}

func (o *DeleteTimeTrackingsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /time-trackings/{id}][%d] deleteTimeTrackingsIdTooManyRequests ", 429)
}

func (o *DeleteTimeTrackingsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
