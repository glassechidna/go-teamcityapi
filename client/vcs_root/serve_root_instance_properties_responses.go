// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/glassechidna/go-teamcityapi/models"
)

// ServeRootInstancePropertiesReader is a Reader for the ServeRootInstanceProperties structure.
type ServeRootInstancePropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeRootInstancePropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeRootInstancePropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeRootInstancePropertiesOK creates a ServeRootInstancePropertiesOK with default headers values
func NewServeRootInstancePropertiesOK() *ServeRootInstancePropertiesOK {
	return &ServeRootInstancePropertiesOK{}
}

/*ServeRootInstancePropertiesOK handles this case with default header values.

successful operation
*/
type ServeRootInstancePropertiesOK struct {
	Payload *models.Properties
}

func (o *ServeRootInstancePropertiesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-roots/{vcsRootLocator}/instances/{vcsRootInstanceLocator}/properties][%d] serveRootInstancePropertiesOK  %+v", 200, o.Payload)
}

func (o *ServeRootInstancePropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Properties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
