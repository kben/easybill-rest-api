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

// PutPositionGroupsIDReader is a Reader for the PutPositionGroupsID structure.
type PutPositionGroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPositionGroupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutPositionGroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutPositionGroupsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutPositionGroupsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutPositionGroupsIDOK creates a PutPositionGroupsIDOK with default headers values
func NewPutPositionGroupsIDOK() *PutPositionGroupsIDOK {
	return &PutPositionGroupsIDOK{}
}

/*PutPositionGroupsIDOK handles this case with default header values.

Successful operation
*/
type PutPositionGroupsIDOK struct {
	Payload *models.PositionGroup
}

func (o *PutPositionGroupsIDOK) Error() string {
	return fmt.Sprintf("[PUT /position-groups/{id}][%d] putPositionGroupsIdOK  %+v", 200, o.Payload)
}

func (o *PutPositionGroupsIDOK) GetPayload() *models.PositionGroup {
	return o.Payload
}

func (o *PutPositionGroupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PositionGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPositionGroupsIDBadRequest creates a PutPositionGroupsIDBadRequest with default headers values
func NewPutPositionGroupsIDBadRequest() *PutPositionGroupsIDBadRequest {
	return &PutPositionGroupsIDBadRequest{}
}

/*PutPositionGroupsIDBadRequest handles this case with default header values.

Invalid position group
*/
type PutPositionGroupsIDBadRequest struct {
}

func (o *PutPositionGroupsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /position-groups/{id}][%d] putPositionGroupsIdBadRequest ", 400)
}

func (o *PutPositionGroupsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutPositionGroupsIDTooManyRequests creates a PutPositionGroupsIDTooManyRequests with default headers values
func NewPutPositionGroupsIDTooManyRequests() *PutPositionGroupsIDTooManyRequests {
	return &PutPositionGroupsIDTooManyRequests{}
}

/*PutPositionGroupsIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type PutPositionGroupsIDTooManyRequests struct {
}

func (o *PutPositionGroupsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /position-groups/{id}][%d] putPositionGroupsIdTooManyRequests ", 429)
}

func (o *PutPositionGroupsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}