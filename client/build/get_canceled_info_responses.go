// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/glassechidna/go-teamcityapi/models"
)

// GetCanceledInfoReader is a Reader for the GetCanceledInfo structure.
type GetCanceledInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCanceledInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCanceledInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCanceledInfoOK creates a GetCanceledInfoOK with default headers values
func NewGetCanceledInfoOK() *GetCanceledInfoOK {
	return &GetCanceledInfoOK{}
}

/*GetCanceledInfoOK handles this case with default header values.

successful operation
*/
type GetCanceledInfoOK struct {
	Payload *models.Comment
}

func (o *GetCanceledInfoOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/canceledInfo][%d] getCanceledInfoOK  %+v", 200, o.Payload)
}

func (o *GetCanceledInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Comment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
