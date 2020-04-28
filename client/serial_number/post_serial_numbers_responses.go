// Code generated by go-swagger; DO NOT EDIT.

package serial_number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// PostSerialNumbersReader is a Reader for the PostSerialNumbers structure.
type PostSerialNumbersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSerialNumbersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSerialNumbersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSerialNumbersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostSerialNumbersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSerialNumbersOK creates a PostSerialNumbersOK with default headers values
func NewPostSerialNumbersOK() *PostSerialNumbersOK {
	return &PostSerialNumbersOK{}
}

/*PostSerialNumbersOK handles this case with default header values.

Successful operation
*/
type PostSerialNumbersOK struct {
	Payload *models.SerialNumber
}

func (o *PostSerialNumbersOK) Error() string {
	return fmt.Sprintf("[POST /serial-numbers][%d] postSerialNumbersOK  %+v", 200, o.Payload)
}

func (o *PostSerialNumbersOK) GetPayload() *models.SerialNumber {
	return o.Payload
}

func (o *PostSerialNumbersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SerialNumber)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSerialNumbersBadRequest creates a PostSerialNumbersBadRequest with default headers values
func NewPostSerialNumbersBadRequest() *PostSerialNumbersBadRequest {
	return &PostSerialNumbersBadRequest{}
}

/*PostSerialNumbersBadRequest handles this case with default header values.

Invalid PositionID
*/
type PostSerialNumbersBadRequest struct {
}

func (o *PostSerialNumbersBadRequest) Error() string {
	return fmt.Sprintf("[POST /serial-numbers][%d] postSerialNumbersBadRequest ", 400)
}

func (o *PostSerialNumbersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSerialNumbersTooManyRequests creates a PostSerialNumbersTooManyRequests with default headers values
func NewPostSerialNumbersTooManyRequests() *PostSerialNumbersTooManyRequests {
	return &PostSerialNumbersTooManyRequests{}
}

/*PostSerialNumbersTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PostSerialNumbersTooManyRequests struct {
}

func (o *PostSerialNumbersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /serial-numbers][%d] postSerialNumbersTooManyRequests ", 429)
}

func (o *PostSerialNumbersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}