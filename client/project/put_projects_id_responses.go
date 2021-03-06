// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// PutProjectsIDReader is a Reader for the PutProjectsID structure.
type PutProjectsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutProjectsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutProjectsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutProjectsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutProjectsIDOK creates a PutProjectsIDOK with default headers values
func NewPutProjectsIDOK() *PutProjectsIDOK {
	return &PutProjectsIDOK{}
}

/*PutProjectsIDOK handles this case with default header values.

Successful operation
*/
type PutProjectsIDOK struct {
	Payload *models.Project
}

func (o *PutProjectsIDOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] putProjectsIdOK  %+v", 200, o.Payload)
}

func (o *PutProjectsIDOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *PutProjectsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutProjectsIDBadRequest creates a PutProjectsIDBadRequest with default headers values
func NewPutProjectsIDBadRequest() *PutProjectsIDBadRequest {
	return &PutProjectsIDBadRequest{}
}

/*PutProjectsIDBadRequest handles this case with default header values.

Invalid project
*/
type PutProjectsIDBadRequest struct {
}

func (o *PutProjectsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] putProjectsIdBadRequest ", 400)
}

func (o *PutProjectsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsIDTooManyRequests creates a PutProjectsIDTooManyRequests with default headers values
func NewPutProjectsIDTooManyRequests() *PutProjectsIDTooManyRequests {
	return &PutProjectsIDTooManyRequests{}
}

/*PutProjectsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutProjectsIDTooManyRequests struct {
}

func (o *PutProjectsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] putProjectsIdTooManyRequests ", 429)
}

func (o *PutProjectsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
