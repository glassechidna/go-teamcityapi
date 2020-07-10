// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/glassechidna/go-teamcityapi/models"
)

// GetAgentRequirementsReader is a Reader for the GetAgentRequirements structure.
type GetAgentRequirementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgentRequirementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAgentRequirementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAgentRequirementsOK creates a GetAgentRequirementsOK with default headers values
func NewGetAgentRequirementsOK() *GetAgentRequirementsOK {
	return &GetAgentRequirementsOK{}
}

/*GetAgentRequirementsOK handles this case with default header values.

successful operation
*/
type GetAgentRequirementsOK struct {
	Payload *models.AgentRequirements
}

func (o *GetAgentRequirementsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/agent-requirements][%d] getAgentRequirementsOK  %+v", 200, o.Payload)
}

func (o *GetAgentRequirementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentRequirements)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
