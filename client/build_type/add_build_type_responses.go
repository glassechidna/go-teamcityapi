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

// AddBuildTypeReader is a Reader for the AddBuildType structure.
type AddBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddBuildTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddBuildTypeOK creates a AddBuildTypeOK with default headers values
func NewAddBuildTypeOK() *AddBuildTypeOK {
	return &AddBuildTypeOK{}
}

/*AddBuildTypeOK handles this case with default header values.

successful operation
*/
type AddBuildTypeOK struct {
	Payload *models.BuildType
}

func (o *AddBuildTypeOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/buildTypes][%d] addBuildTypeOK  %+v", 200, o.Payload)
}

func (o *AddBuildTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
