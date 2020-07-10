// Code generated by go-swagger; DO NOT EDIT.

package build_queue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/glassechidna/go-teamcityapi/models"
)

// ServeCompatibleAgentsReader is a Reader for the ServeCompatibleAgents structure.
type ServeCompatibleAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeCompatibleAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeCompatibleAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeCompatibleAgentsOK creates a ServeCompatibleAgentsOK with default headers values
func NewServeCompatibleAgentsOK() *ServeCompatibleAgentsOK {
	return &ServeCompatibleAgentsOK{}
}

/*ServeCompatibleAgentsOK handles this case with default header values.

successful operation
*/
type ServeCompatibleAgentsOK struct {
	Payload *models.Agents
}

func (o *ServeCompatibleAgentsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildQueue/{queuedBuildLocator}/compatibleAgents][%d] serveCompatibleAgentsOK  %+v", 200, o.Payload)
}

func (o *ServeCompatibleAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Agents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
