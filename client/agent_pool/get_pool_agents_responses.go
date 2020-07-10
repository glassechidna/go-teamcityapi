// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/glassechidna/go-teamcityapi/models"
)

// GetPoolAgentsReader is a Reader for the GetPoolAgents structure.
type GetPoolAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPoolAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPoolAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPoolAgentsOK creates a GetPoolAgentsOK with default headers values
func NewGetPoolAgentsOK() *GetPoolAgentsOK {
	return &GetPoolAgentsOK{}
}

/*GetPoolAgentsOK handles this case with default header values.

successful operation
*/
type GetPoolAgentsOK struct {
	Payload *models.Agents
}

func (o *GetPoolAgentsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/agentPools/{agentPoolLocator}/agents][%d] getPoolAgentsOK  %+v", 200, o.Payload)
}

func (o *GetPoolAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Agents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}