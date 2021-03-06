// Code generated by go-swagger; DO NOT EDIT.

package discount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// GetDiscountsPositionGroupIDReader is a Reader for the GetDiscountsPositionGroupID structure.
type GetDiscountsPositionGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscountsPositionGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscountsPositionGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDiscountsPositionGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDiscountsPositionGroupIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDiscountsPositionGroupIDOK creates a GetDiscountsPositionGroupIDOK with default headers values
func NewGetDiscountsPositionGroupIDOK() *GetDiscountsPositionGroupIDOK {
	return &GetDiscountsPositionGroupIDOK{}
}

/*GetDiscountsPositionGroupIDOK handles this case with default header values.

Successful operation
*/
type GetDiscountsPositionGroupIDOK struct {
	Payload *models.DiscountPositionGroup
}

func (o *GetDiscountsPositionGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /discounts/position-group/{id}][%d] getDiscountsPositionGroupIdOK  %+v", 200, o.Payload)
}

func (o *GetDiscountsPositionGroupIDOK) GetPayload() *models.DiscountPositionGroup {
	return o.Payload
}

func (o *GetDiscountsPositionGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DiscountPositionGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscountsPositionGroupIDNotFound creates a GetDiscountsPositionGroupIDNotFound with default headers values
func NewGetDiscountsPositionGroupIDNotFound() *GetDiscountsPositionGroupIDNotFound {
	return &GetDiscountsPositionGroupIDNotFound{}
}

/*GetDiscountsPositionGroupIDNotFound handles this case with default header values.

Not found
*/
type GetDiscountsPositionGroupIDNotFound struct {
}

func (o *GetDiscountsPositionGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /discounts/position-group/{id}][%d] getDiscountsPositionGroupIdNotFound ", 404)
}

func (o *GetDiscountsPositionGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiscountsPositionGroupIDTooManyRequests creates a GetDiscountsPositionGroupIDTooManyRequests with default headers values
func NewGetDiscountsPositionGroupIDTooManyRequests() *GetDiscountsPositionGroupIDTooManyRequests {
	return &GetDiscountsPositionGroupIDTooManyRequests{}
}

/*GetDiscountsPositionGroupIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetDiscountsPositionGroupIDTooManyRequests struct {
}

func (o *GetDiscountsPositionGroupIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /discounts/position-group/{id}][%d] getDiscountsPositionGroupIdTooManyRequests ", 429)
}

func (o *GetDiscountsPositionGroupIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
