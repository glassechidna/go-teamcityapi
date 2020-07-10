// Code generated by go-swagger; DO NOT EDIT.

package build_queue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new build queue API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for build queue API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostAppRestBuildQueueQueuedBuildLocator post app rest build queue queued build locator API
*/
func (a *Client) PostAppRestBuildQueueQueuedBuildLocator(params *PostAppRestBuildQueueQueuedBuildLocatorParams) (*PostAppRestBuildQueueQueuedBuildLocatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAppRestBuildQueueQueuedBuildLocatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAppRestBuildQueueQueuedBuildLocator",
		Method:             "POST",
		PathPattern:        "/app/rest/buildQueue/{queuedBuildLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAppRestBuildQueueQueuedBuildLocatorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAppRestBuildQueueQueuedBuildLocatorOK), nil

}

/*
DeleteBuildsExperimental delete builds experimental API
*/
func (a *Client) DeleteBuildsExperimental(params *DeleteBuildsExperimentalParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBuildsExperimentalParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBuildsExperimental",
		Method:             "DELETE",
		PathPattern:        "/app/rest/buildQueue",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBuildsExperimentalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetBuild get build API
*/
func (a *Client) GetBuild(params *GetBuildParams) (*GetBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuild",
		Method:             "GET",
		PathPattern:        "/app/rest/buildQueue/{queuedBuildLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBuildOK), nil

}

/*
GetBuilds get builds API
*/
func (a *Client) GetBuilds(params *GetBuildsParams) (*GetBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuilds",
		Method:             "GET",
		PathPattern:        "/app/rest/buildQueue",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBuildsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBuildsOK), nil

}

/*
QueueNewBuild queue new build API
*/
func (a *Client) QueueNewBuild(params *QueueNewBuildParams) (*QueueNewBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueueNewBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queueNewBuild",
		Method:             "POST",
		PathPattern:        "/app/rest/buildQueue",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueueNewBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueueNewBuildOK), nil

}

/*
ReplaceBuilds replace builds API
*/
func (a *Client) ReplaceBuilds(params *ReplaceBuildsParams) (*ReplaceBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBuilds",
		Method:             "PUT",
		PathPattern:        "/app/rest/buildQueue",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceBuildsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReplaceBuildsOK), nil

}

/*
ServeCompatibleAgents serve compatible agents API
*/
func (a *Client) ServeCompatibleAgents(params *ServeCompatibleAgentsParams) (*ServeCompatibleAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeCompatibleAgentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveCompatibleAgents",
		Method:             "GET",
		PathPattern:        "/app/rest/buildQueue/{queuedBuildLocator}/compatibleAgents",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServeCompatibleAgentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeCompatibleAgentsOK), nil

}

/*
SetBuildQueueOrder set build queue order API
*/
func (a *Client) SetBuildQueueOrder(params *SetBuildQueueOrderParams) (*SetBuildQueueOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetBuildQueueOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setBuildQueueOrder",
		Method:             "PUT",
		PathPattern:        "/app/rest/buildQueue/order",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetBuildQueueOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetBuildQueueOrderOK), nil

}

/*
SetBuildQueuePosition set build queue position API
*/
func (a *Client) SetBuildQueuePosition(params *SetBuildQueuePositionParams) (*SetBuildQueuePositionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetBuildQueuePositionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setBuildQueuePosition",
		Method:             "PUT",
		PathPattern:        "/app/rest/buildQueue/order/{queuePosition}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetBuildQueuePositionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetBuildQueuePositionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
