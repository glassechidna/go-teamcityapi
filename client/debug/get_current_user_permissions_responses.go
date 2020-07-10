// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetCurrentUserPermissionsReader is a Reader for the GetCurrentUserPermissions structure.
type GetCurrentUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCurrentUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrentUserPermissionsOK creates a GetCurrentUserPermissionsOK with default headers values
func NewGetCurrentUserPermissionsOK() *GetCurrentUserPermissionsOK {
	return &GetCurrentUserPermissionsOK{}
}

/*GetCurrentUserPermissionsOK handles this case with default header values.

successful operation
*/
type GetCurrentUserPermissionsOK struct {
	Payload string
}

func (o *GetCurrentUserPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/currentUserPermissions][%d] getCurrentUserPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
