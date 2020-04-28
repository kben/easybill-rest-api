// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stock API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stock API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetStocks(params *GetStocksParams, authInfo runtime.ClientAuthInfoWriter) (*GetStocksOK, error)

	GetStocksID(params *GetStocksIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetStocksIDOK, error)

	PostStocks(params *PostStocksParams, authInfo runtime.ClientAuthInfoWriter) (*PostStocksCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetStocks fetches a list of stock entries for positions
*/
func (a *Client) GetStocks(params *GetStocksParams, authInfo runtime.ClientAuthInfoWriter) (*GetStocksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStocksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStocks",
		Method:             "GET",
		PathPattern:        "/stocks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStocksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStocksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStocks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStocksID fetches an stock entry for a position
*/
func (a *Client) GetStocksID(params *GetStocksIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetStocksIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStocksIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStocksID",
		Method:             "GET",
		PathPattern:        "/stocks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStocksIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStocksIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStocksID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostStocks creates a stock entry for a position
*/
func (a *Client) PostStocks(params *PostStocksParams, authInfo runtime.ClientAuthInfoWriter) (*PostStocksCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStocksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStocks",
		Method:             "POST",
		PathPattern:        "/stocks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStocksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostStocksCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostStocks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}