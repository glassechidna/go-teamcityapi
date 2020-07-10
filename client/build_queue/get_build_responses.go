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

// GetBuildReader is a Reader for the GetBuild structure.
type GetBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBuildOK creates a GetBuildOK with default headers values
func NewGetBuildOK() *GetBuildOK {
	return &GetBuildOK{}
}

/*GetBuildOK handles this case with default header values.

successful operation
*/
type GetBuildOK struct {
	Payload *models.Build
}

func (o *GetBuildOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildQueue/{queuedBuildLocator}][%d] getBuildOK  %+v", 200, o.Payload)
}

func (o *GetBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Build)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
