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

// SetTemplatesReader is a Reader for the SetTemplates structure.
type SetTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetTemplatesOK creates a SetTemplatesOK with default headers values
func NewSetTemplatesOK() *SetTemplatesOK {
	return &SetTemplatesOK{}
}

/*SetTemplatesOK handles this case with default header values.

successful operation
*/
type SetTemplatesOK struct {
	Payload *models.BuildTypes
}

func (o *SetTemplatesOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/templates][%d] setTemplatesOK  %+v", 200, o.Payload)
}

func (o *SetTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
