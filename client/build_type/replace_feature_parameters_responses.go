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

// ReplaceFeatureParametersReader is a Reader for the ReplaceFeatureParameters structure.
type ReplaceFeatureParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceFeatureParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceFeatureParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceFeatureParametersOK creates a ReplaceFeatureParametersOK with default headers values
func NewReplaceFeatureParametersOK() *ReplaceFeatureParametersOK {
	return &ReplaceFeatureParametersOK{}
}

/*ReplaceFeatureParametersOK handles this case with default header values.

successful operation
*/
type ReplaceFeatureParametersOK struct {
	Payload *models.Properties
}

func (o *ReplaceFeatureParametersOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/features/{featureId}/parameters][%d] replaceFeatureParametersOK  %+v", 200, o.Payload)
}

func (o *ReplaceFeatureParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Properties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
