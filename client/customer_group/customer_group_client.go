// Code generated by go-swagger; DO NOT EDIT.

package customer_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCustomerGroupsID(params *DeleteCustomerGroupsIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCustomerGroupsIDNoContent, error)

	GetCustomerGroups(params *GetCustomerGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerGroupsOK, error)

	GetCustomerGroupsID(params *GetCustomerGroupsIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerGroupsIDOK, error)

	PostCustomerGroups(params *PostCustomerGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomerGroupsCreated, error)

	PutCustomerGroupsID(params *PutCustomerGroupsIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomerGroupsIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteCustomerGroupsID deletes customer group
*/
func (a *Client) DeleteCustomerGroupsID(params *DeleteCustomerGroupsIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCustomerGroupsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomerGroupsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCustomerGroupsID",
		Method:             "DELETE",
		PathPattern:        "/customer-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCustomerGroupsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCustomerGroupsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteCustomerGroupsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCustomerGroups fetches customer group list
*/
func (a *Client) GetCustomerGroups(params *GetCustomerGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomerGroups",
		Method:             "GET",
		PathPattern:        "/customer-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomerGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCustomerGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCustomerGroupsID fetches customer group
*/
func (a *Client) GetCustomerGroupsID(params *GetCustomerGroupsIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerGroupsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerGroupsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCustomerGroupsID",
		Method:             "GET",
		PathPattern:        "/customer-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomerGroupsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerGroupsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCustomerGroupsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCustomerGroups creates customer group
*/
func (a *Client) PostCustomerGroups(params *PostCustomerGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCustomerGroupsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCustomerGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCustomerGroups",
		Method:             "POST",
		PathPattern:        "/customer-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCustomerGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCustomerGroupsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostCustomerGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutCustomerGroupsID updates customer group
*/
func (a *Client) PutCustomerGroupsID(params *PutCustomerGroupsIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutCustomerGroupsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCustomerGroupsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCustomerGroupsID",
		Method:             "PUT",
		PathPattern:        "/customer-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutCustomerGroupsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCustomerGroupsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutCustomerGroupsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
