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

// ServeRootInstancesReader is a Reader for the ServeRootInstances structure.
type ServeRootInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeRootInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeRootInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeRootInstancesOK creates a ServeRootInstancesOK with default headers values
func NewServeRootInstancesOK() *ServeRootInstancesOK {
	return &ServeRootInstancesOK{}
}

/*ServeRootInstancesOK handles this case with default header values.

successful operation
*/
type ServeRootInstancesOK struct {
	Payload *models.VcsRootInstances
}

func (o *ServeRootInstancesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-roots/{vcsRootLocator}/instances][%d] serveRootInstancesOK  %+v", 200, o.Payload)
}

func (o *ServeRootInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRootInstances)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
