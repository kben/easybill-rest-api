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

// PostStocksReader is a Reader for the PostStocks structure.
type PostStocksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStocksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostStocksCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStocksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostStocksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostStocksCreated creates a PostStocksCreated with default headers values
func NewPostStocksCreated() *PostStocksCreated {
	return &PostStocksCreated{}
}

/*PostStocksCreated handles this case with default header values.

Successful operation
*/
type PostStocksCreated struct {
	Payload *models.Stock
}

func (o *PostStocksCreated) Error() string {
	return fmt.Sprintf("[POST /stocks][%d] postStocksCreated  %+v", 201, o.Payload)
}

func (o *PostStocksCreated) GetPayload() *models.Stock {
	return o.Payload
}

func (o *PostStocksCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Stock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStocksBadRequest creates a PostStocksBadRequest with default headers values
func NewPostStocksBadRequest() *PostStocksBadRequest {
	return &PostStocksBadRequest{}
}

/*PostStocksBadRequest handles this case with default header values.

Invalid position_id or stock_count
*/
type PostStocksBadRequest struct {
}

func (o *PostStocksBadRequest) Error() string {
	return fmt.Sprintf("[POST /stocks][%d] postStocksBadRequest ", 400)
}

func (o *PostStocksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStocksTooManyRequests creates a PostStocksTooManyRequests with default headers values
func NewPostStocksTooManyRequests() *PostStocksTooManyRequests {
	return &PostStocksTooManyRequests{}
}

/*PostStocksTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PostStocksTooManyRequests struct {
}

func (o *PostStocksTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /stocks][%d] postStocksTooManyRequests ", 429)
}

func (o *PostStocksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}