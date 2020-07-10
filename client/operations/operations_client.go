// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServeAPIVersion serve Api version API
*/
func (a *Client) ServeAPIVersion(params *ServeAPIVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ServeAPIVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeAPIVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveApiVersion",
		Method:             "GET",
		PathPattern:        "/app/rest/apiVersion",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServeAPIVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeAPIVersionOK), nil

}

/*
ServeBuildFieldShort serve build field short API
*/
func (a *Client) ServeBuildFieldShort(params *ServeBuildFieldShortParams, authInfo runtime.ClientAuthInfoWriter) (*ServeBuildFieldShortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeBuildFieldShortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveBuildFieldShort",
		Method:             "GET",
		PathPattern:        "/app/rest/{projectLocator}/{btLocator}/{buildLocator}/{field}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServeBuildFieldShortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeBuildFieldShortOK), nil

}

/*
ServePluginInfo serve plugin info API
*/
func (a *Client) ServePluginInfo(params *ServePluginInfoParams, authInfo runtime.ClientAuthInfoWriter) (*ServePluginInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServePluginInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servePluginInfo",
		Method:             "GET",
		PathPattern:        "/app/rest/info",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServePluginInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServePluginInfoOK), nil

}

/*
ServeVersion serve version API
*/
func (a *Client) ServeVersion(params *ServeVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ServeVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveVersion",
		Method:             "GET",
		PathPattern:        "/app/rest/version",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServeVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeVersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
