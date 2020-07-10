// Code generated by go-swagger; DO NOT EDIT.

package build_queue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/glassechidna/go-teamcityapi/models"
)

// ReplaceBuildsReader is a Reader for the ReplaceBuilds structure.
type ReplaceBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceBuildsOK creates a ReplaceBuildsOK with default headers values
func NewReplaceBuildsOK() *ReplaceBuildsOK {
	return &ReplaceBuildsOK{}
}

/*ReplaceBuildsOK handles this case with default header values.

successful operation
*/
type ReplaceBuildsOK struct {
	Payload *models.Builds
}

func (o *ReplaceBuildsOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildQueue][%d] replaceBuildsOK  %+v", 200, o.Payload)
}

func (o *ReplaceBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Builds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
