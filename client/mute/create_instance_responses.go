// Code generated by go-swagger; DO NOT EDIT.

package mute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/glassechidna/go-teamcityapi/models"
)

// CreateInstanceReader is a Reader for the CreateInstance structure.
type CreateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateInstanceOK creates a CreateInstanceOK with default headers values
func NewCreateInstanceOK() *CreateInstanceOK {
	return &CreateInstanceOK{}
}

/*CreateInstanceOK handles this case with default header values.

successful operation
*/
type CreateInstanceOK struct {
	Payload *models.Mute
}

func (o *CreateInstanceOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/mutes][%d] createInstanceOK  %+v", 200, o.Payload)
}

func (o *CreateInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Mute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
