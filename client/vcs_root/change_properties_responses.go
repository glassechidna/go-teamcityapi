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

// ChangePropertiesReader is a Reader for the ChangeProperties structure.
type ChangePropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangePropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangePropertiesOK creates a ChangePropertiesOK with default headers values
func NewChangePropertiesOK() *ChangePropertiesOK {
	return &ChangePropertiesOK{}
}

/*ChangePropertiesOK handles this case with default header values.

successful operation
*/
type ChangePropertiesOK struct {
	Payload *models.Properties
}

func (o *ChangePropertiesOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/vcs-roots/{vcsRootLocator}/properties][%d] changePropertiesOK  %+v", 200, o.Payload)
}

func (o *ChangePropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Properties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
