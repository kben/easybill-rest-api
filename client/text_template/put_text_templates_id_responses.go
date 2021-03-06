// Code generated by go-swagger; DO NOT EDIT.

package text_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// PutTextTemplatesIDReader is a Reader for the PutTextTemplatesID structure.
type PutTextTemplatesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTextTemplatesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTextTemplatesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTextTemplatesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTextTemplatesIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutTextTemplatesIDOK creates a PutTextTemplatesIDOK with default headers values
func NewPutTextTemplatesIDOK() *PutTextTemplatesIDOK {
	return &PutTextTemplatesIDOK{}
}

/*PutTextTemplatesIDOK handles this case with default header values.

Successful operation
*/
type PutTextTemplatesIDOK struct {
	Payload *models.TextTemplate
}

func (o *PutTextTemplatesIDOK) Error() string {
	return fmt.Sprintf("[PUT /text-templates/{id}][%d] putTextTemplatesIdOK  %+v", 200, o.Payload)
}

func (o *PutTextTemplatesIDOK) GetPayload() *models.TextTemplate {
	return o.Payload
}

func (o *PutTextTemplatesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TextTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTextTemplatesIDBadRequest creates a PutTextTemplatesIDBadRequest with default headers values
func NewPutTextTemplatesIDBadRequest() *PutTextTemplatesIDBadRequest {
	return &PutTextTemplatesIDBadRequest{}
}

/*PutTextTemplatesIDBadRequest handles this case with default header values.

Invalid text template
*/
type PutTextTemplatesIDBadRequest struct {
}

func (o *PutTextTemplatesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /text-templates/{id}][%d] putTextTemplatesIdBadRequest ", 400)
}

func (o *PutTextTemplatesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutTextTemplatesIDTooManyRequests creates a PutTextTemplatesIDTooManyRequests with default headers values
func NewPutTextTemplatesIDTooManyRequests() *PutTextTemplatesIDTooManyRequests {
	return &PutTextTemplatesIDTooManyRequests{}
}

/*PutTextTemplatesIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutTextTemplatesIDTooManyRequests struct {
}

func (o *PutTextTemplatesIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /text-templates/{id}][%d] putTextTemplatesIdTooManyRequests ", 429)
}

func (o *PutTextTemplatesIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
