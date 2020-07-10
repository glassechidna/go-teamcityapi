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

// PutAppRestBuildsBuildLocatorAttributesNameReader is a Reader for the PutAppRestBuildsBuildLocatorAttributesName structure.
type PutAppRestBuildsBuildLocatorAttributesNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAppRestBuildsBuildLocatorAttributesNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAppRestBuildsBuildLocatorAttributesNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAppRestBuildsBuildLocatorAttributesNameOK creates a PutAppRestBuildsBuildLocatorAttributesNameOK with default headers values
func NewPutAppRestBuildsBuildLocatorAttributesNameOK() *PutAppRestBuildsBuildLocatorAttributesNameOK {
	return &PutAppRestBuildsBuildLocatorAttributesNameOK{}
}

/*PutAppRestBuildsBuildLocatorAttributesNameOK handles this case with default header values.

successful operation
*/
type PutAppRestBuildsBuildLocatorAttributesNameOK struct {
	Payload *models.Property
}

func (o *PutAppRestBuildsBuildLocatorAttributesNameOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/builds/{buildLocator}/attributes/{name}][%d] putAppRestBuildsBuildLocatorAttributesNameOK  %+v", 200, o.Payload)
}

func (o *PutAppRestBuildsBuildLocatorAttributesNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Property)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
