// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ServeUserPropertyReader is a Reader for the ServeUserProperty structure.
type ServeUserPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeUserPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeUserPropertyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeUserPropertyOK creates a ServeUserPropertyOK with default headers values
func NewServeUserPropertyOK() *ServeUserPropertyOK {
	return &ServeUserPropertyOK{}
}

/*ServeUserPropertyOK handles this case with default header values.

successful operation
*/
type ServeUserPropertyOK struct {
	Payload string
}

func (o *ServeUserPropertyOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/users/{userLocator}/properties/{name}][%d] serveUserPropertyOK  %+v", 200, o.Payload)
}

func (o *ServeUserPropertyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
