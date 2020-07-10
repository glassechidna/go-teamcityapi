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

// ReplaceFeatureReader is a Reader for the ReplaceFeature structure.
type ReplaceFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceFeatureOK creates a ReplaceFeatureOK with default headers values
func NewReplaceFeatureOK() *ReplaceFeatureOK {
	return &ReplaceFeatureOK{}
}

/*ReplaceFeatureOK handles this case with default header values.

successful operation
*/
type ReplaceFeatureOK struct {
	Payload *models.Feature
}

func (o *ReplaceFeatureOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/features/{featureId}][%d] replaceFeatureOK  %+v", 200, o.Payload)
}

func (o *ReplaceFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Feature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}