// Code generated by go-swagger; DO NOT EDIT.

package post_box

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new post box API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for post box API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePostBoxesID(params *DeletePostBoxesIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePostBoxesIDNoContent, error)

	GetPostBoxes(params *GetPostBoxesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPostBoxesOK, error)

	GetPostBoxesID(params *GetPostBoxesIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPostBoxesIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeletePostBoxesID deletes post box
*/
func (a *Client) DeletePostBoxesID(params *DeletePostBoxesIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePostBoxesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePostBoxesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePostBoxesID",
		Method:             "DELETE",
		PathPattern:        "/post-boxes/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePostBoxesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePostBoxesIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeletePostBoxesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPostBoxes fetches post box list
*/
func (a *Client) GetPostBoxes(params *GetPostBoxesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPostBoxesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostBoxesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPostBoxes",
		Method:             "GET",
		PathPattern:        "/post-boxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPostBoxesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPostBoxesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPostBoxes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPostBoxesID fetches post box
*/
func (a *Client) GetPostBoxesID(params *GetPostBoxesIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPostBoxesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostBoxesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPostBoxesID",
		Method:             "GET",
		PathPattern:        "/post-boxes/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPostBoxesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPostBoxesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPostBoxesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
