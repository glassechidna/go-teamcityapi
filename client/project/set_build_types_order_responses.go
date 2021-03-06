// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/glassechidna/go-teamcityapi/models"
)

// SetBuildTypesOrderReader is a Reader for the SetBuildTypesOrder structure.
type SetBuildTypesOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetBuildTypesOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetBuildTypesOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetBuildTypesOrderOK creates a SetBuildTypesOrderOK with default headers values
func NewSetBuildTypesOrderOK() *SetBuildTypesOrderOK {
	return &SetBuildTypesOrderOK{}
}

/*SetBuildTypesOrderOK handles this case with default header values.

successful operation
*/
type SetBuildTypesOrderOK struct {
	Payload *models.BuildTypes
}

func (o *SetBuildTypesOrderOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/projects/{projectLocator}/order/buildTypes][%d] setBuildTypesOrderOK  %+v", 200, o.Payload)
}

func (o *SetBuildTypesOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
