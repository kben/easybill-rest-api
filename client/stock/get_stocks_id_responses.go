// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// GetStocksIDReader is a Reader for the GetStocksID structure.
type GetStocksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStocksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStocksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStocksIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetStocksIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStocksIDOK creates a GetStocksIDOK with default headers values
func NewGetStocksIDOK() *GetStocksIDOK {
	return &GetStocksIDOK{}
}

/*GetStocksIDOK handles this case with default header values.

Successful operation
*/
type GetStocksIDOK struct {
	Payload *models.Stock
}

func (o *GetStocksIDOK) Error() string {
	return fmt.Sprintf("[GET /stocks/{id}][%d] getStocksIdOK  %+v", 200, o.Payload)
}

func (o *GetStocksIDOK) GetPayload() *models.Stock {
	return o.Payload
}

func (o *GetStocksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Stock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStocksIDNotFound creates a GetStocksIDNotFound with default headers values
func NewGetStocksIDNotFound() *GetStocksIDNotFound {
	return &GetStocksIDNotFound{}
}

/*GetStocksIDNotFound handles this case with default header values.

Not found
*/
type GetStocksIDNotFound struct {
}

func (o *GetStocksIDNotFound) Error() string {
	return fmt.Sprintf("[GET /stocks/{id}][%d] getStocksIdNotFound ", 404)
}

func (o *GetStocksIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStocksIDTooManyRequests creates a GetStocksIDTooManyRequests with default headers values
func NewGetStocksIDTooManyRequests() *GetStocksIDTooManyRequests {
	return &GetStocksIDTooManyRequests{}
}

/*GetStocksIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetStocksIDTooManyRequests struct {
}

func (o *GetStocksIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /stocks/{id}][%d] getStocksIdTooManyRequests ", 429)
}

func (o *GetStocksIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}