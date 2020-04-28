// Code generated by go-swagger; DO NOT EDIT.

package position_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kben/easybill-rest-api/models"
)

// PostPositionGroupsReader is a Reader for the PostPositionGroups structure.
type PostPositionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPositionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostPositionGroupsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewPostPositionGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPositionGroupsCreated creates a PostPositionGroupsCreated with default headers values
func NewPostPositionGroupsCreated() *PostPositionGroupsCreated {
	return &PostPositionGroupsCreated{}
}

/*PostPositionGroupsCreated handles this case with default header values.

Successful operation
*/
type PostPositionGroupsCreated struct {
	Payload *models.PositionGroup
}

func (o *PostPositionGroupsCreated) Error() string {
	return fmt.Sprintf("[POST /position-groups][%d] postPositionGroupsCreated  %+v", 201, o.Payload)
}

func (o *PostPositionGroupsCreated) GetPayload() *models.PositionGroup {
	return o.Payload
}

func (o *PostPositionGroupsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PositionGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPositionGroupsTooManyRequests creates a PostPositionGroupsTooManyRequests with default headers values
func NewPostPositionGroupsTooManyRequests() *PostPositionGroupsTooManyRequests {
	return &PostPositionGroupsTooManyRequests{}
}

/*PostPositionGroupsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PostPositionGroupsTooManyRequests struct {
}

func (o *PostPositionGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /position-groups][%d] postPositionGroupsTooManyRequests ", 429)
}

func (o *PostPositionGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
