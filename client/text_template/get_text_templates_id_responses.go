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

// GetTextTemplatesIDReader is a Reader for the GetTextTemplatesID structure.
type GetTextTemplatesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTextTemplatesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTextTemplatesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTextTemplatesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTextTemplatesIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTextTemplatesIDOK creates a GetTextTemplatesIDOK with default headers values
func NewGetTextTemplatesIDOK() *GetTextTemplatesIDOK {
	return &GetTextTemplatesIDOK{}
}

/*GetTextTemplatesIDOK handles this case with default header values.

Successful operation
*/
type GetTextTemplatesIDOK struct {
	Payload *models.TextTemplate
}

func (o *GetTextTemplatesIDOK) Error() string {
	return fmt.Sprintf("[GET /text-templates/{id}][%d] getTextTemplatesIdOK  %+v", 200, o.Payload)
}

func (o *GetTextTemplatesIDOK) GetPayload() *models.TextTemplate {
	return o.Payload
}

func (o *GetTextTemplatesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TextTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextTemplatesIDNotFound creates a GetTextTemplatesIDNotFound with default headers values
func NewGetTextTemplatesIDNotFound() *GetTextTemplatesIDNotFound {
	return &GetTextTemplatesIDNotFound{}
}

/*GetTextTemplatesIDNotFound handles this case with default header values.

Not found
*/
type GetTextTemplatesIDNotFound struct {
}

func (o *GetTextTemplatesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /text-templates/{id}][%d] getTextTemplatesIdNotFound ", 404)
}

func (o *GetTextTemplatesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTextTemplatesIDTooManyRequests creates a GetTextTemplatesIDTooManyRequests with default headers values
func NewGetTextTemplatesIDTooManyRequests() *GetTextTemplatesIDTooManyRequests {
	return &GetTextTemplatesIDTooManyRequests{}
}

/*GetTextTemplatesIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetTextTemplatesIDTooManyRequests struct {
}

func (o *GetTextTemplatesIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /text-templates/{id}][%d] getTextTemplatesIdTooManyRequests ", 429)
}

func (o *GetTextTemplatesIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
