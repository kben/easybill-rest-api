// Code generated by go-swagger; DO NOT EDIT.

package time_tracking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// PutTimeTrackingsIDReader is a Reader for the PutTimeTrackingsID structure.
type PutTimeTrackingsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTimeTrackingsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTimeTrackingsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTimeTrackingsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTimeTrackingsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutTimeTrackingsIDOK creates a PutTimeTrackingsIDOK with default headers values
func NewPutTimeTrackingsIDOK() *PutTimeTrackingsIDOK {
	return &PutTimeTrackingsIDOK{}
}

/*PutTimeTrackingsIDOK handles this case with default header values.

Successful operation
*/
type PutTimeTrackingsIDOK struct {
	Payload *models.TimeTracking
}

func (o *PutTimeTrackingsIDOK) Error() string {
	return fmt.Sprintf("[PUT /time-trackings/{id}][%d] putTimeTrackingsIdOK  %+v", 200, o.Payload)
}

func (o *PutTimeTrackingsIDOK) GetPayload() *models.TimeTracking {
	return o.Payload
}

func (o *PutTimeTrackingsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeTracking)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTimeTrackingsIDBadRequest creates a PutTimeTrackingsIDBadRequest with default headers values
func NewPutTimeTrackingsIDBadRequest() *PutTimeTrackingsIDBadRequest {
	return &PutTimeTrackingsIDBadRequest{}
}

/*PutTimeTrackingsIDBadRequest handles this case with default header values.

Invalid time tracking
*/
type PutTimeTrackingsIDBadRequest struct {
}

func (o *PutTimeTrackingsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /time-trackings/{id}][%d] putTimeTrackingsIdBadRequest ", 400)
}

func (o *PutTimeTrackingsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutTimeTrackingsIDTooManyRequests creates a PutTimeTrackingsIDTooManyRequests with default headers values
func NewPutTimeTrackingsIDTooManyRequests() *PutTimeTrackingsIDTooManyRequests {
	return &PutTimeTrackingsIDTooManyRequests{}
}

/*PutTimeTrackingsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutTimeTrackingsIDTooManyRequests struct {
}

func (o *PutTimeTrackingsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /time-trackings/{id}][%d] putTimeTrackingsIdTooManyRequests ", 429)
}

func (o *PutTimeTrackingsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
