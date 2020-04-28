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

// PostTextTemplatesReader is a Reader for the PostTextTemplates structure.
type PostTextTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTextTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostTextTemplatesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPostTextTemplatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTextTemplatesCreated creates a PostTextTemplatesCreated with default headers values
func NewPostTextTemplatesCreated() *PostTextTemplatesCreated {
	return &PostTextTemplatesCreated{}
}

/*PostTextTemplatesCreated handles this case with default header values.

Successful operation
*/
type PostTextTemplatesCreated struct {
	Payload *models.TextTemplate
}

func (o *PostTextTemplatesCreated) Error() string {
	return fmt.Sprintf("[POST /text-templates][%d] postTextTemplatesCreated  %+v", 201, o.Payload)
}

func (o *PostTextTemplatesCreated) GetPayload() *models.TextTemplate {
	return o.Payload
}

func (o *PostTextTemplatesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TextTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextTemplatesTooManyRequests creates a PostTextTemplatesTooManyRequests with default headers values
func NewPostTextTemplatesTooManyRequests() *PostTextTemplatesTooManyRequests {
	return &PostTextTemplatesTooManyRequests{}
}

/*PostTextTemplatesTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PostTextTemplatesTooManyRequests struct {
}

func (o *PostTextTemplatesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /text-templates][%d] postTextTemplatesTooManyRequests ", 429)
}

func (o *PostTextTemplatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}