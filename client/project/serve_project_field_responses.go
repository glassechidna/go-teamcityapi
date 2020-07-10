// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ServeProjectFieldReader is a Reader for the ServeProjectField structure.
type ServeProjectFieldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeProjectFieldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeProjectFieldOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeProjectFieldOK creates a ServeProjectFieldOK with default headers values
func NewServeProjectFieldOK() *ServeProjectFieldOK {
	return &ServeProjectFieldOK{}
}

/*ServeProjectFieldOK handles this case with default header values.

successful operation
*/
type ServeProjectFieldOK struct {
	Payload string
}

func (o *ServeProjectFieldOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/{field}][%d] serveProjectFieldOK  %+v", 200, o.Payload)
}

func (o *ServeProjectFieldOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
