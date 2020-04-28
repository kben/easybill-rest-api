// Code generated by go-swagger; DO NOT EDIT.

package post_box

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// GetPostBoxesIDReader is a Reader for the GetPostBoxesID structure.
type GetPostBoxesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPostBoxesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPostBoxesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPostBoxesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPostBoxesIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPostBoxesIDOK creates a GetPostBoxesIDOK with default headers values
func NewGetPostBoxesIDOK() *GetPostBoxesIDOK {
	return &GetPostBoxesIDOK{}
}

/*GetPostBoxesIDOK handles this case with default header values.

Successful operation
*/
type GetPostBoxesIDOK struct {
	Payload *models.PostBox
}

func (o *GetPostBoxesIDOK) Error() string {
	return fmt.Sprintf("[GET /post-boxes/{id}][%d] getPostBoxesIdOK  %+v", 200, o.Payload)
}

func (o *GetPostBoxesIDOK) GetPayload() *models.PostBox {
	return o.Payload
}

func (o *GetPostBoxesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostBox)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPostBoxesIDNotFound creates a GetPostBoxesIDNotFound with default headers values
func NewGetPostBoxesIDNotFound() *GetPostBoxesIDNotFound {
	return &GetPostBoxesIDNotFound{}
}

/*GetPostBoxesIDNotFound handles this case with default header values.

Not found
*/
type GetPostBoxesIDNotFound struct {
}

func (o *GetPostBoxesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /post-boxes/{id}][%d] getPostBoxesIdNotFound ", 404)
}

func (o *GetPostBoxesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPostBoxesIDTooManyRequests creates a GetPostBoxesIDTooManyRequests with default headers values
func NewGetPostBoxesIDTooManyRequests() *GetPostBoxesIDTooManyRequests {
	return &GetPostBoxesIDTooManyRequests{}
}

/*GetPostBoxesIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetPostBoxesIDTooManyRequests struct {
}

func (o *GetPostBoxesIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /post-boxes/{id}][%d] getPostBoxesIdTooManyRequests ", 429)
}

func (o *GetPostBoxesIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
