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

// ReplaceSnapshotDepsReader is a Reader for the ReplaceSnapshotDeps structure.
type ReplaceSnapshotDepsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceSnapshotDepsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceSnapshotDepsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceSnapshotDepsOK creates a ReplaceSnapshotDepsOK with default headers values
func NewReplaceSnapshotDepsOK() *ReplaceSnapshotDepsOK {
	return &ReplaceSnapshotDepsOK{}
}

/*ReplaceSnapshotDepsOK handles this case with default header values.

successful operation
*/
type ReplaceSnapshotDepsOK struct {
	Payload *models.SnapshotDependencies
}

func (o *ReplaceSnapshotDepsOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/snapshot-dependencies][%d] replaceSnapshotDepsOK  %+v", 200, o.Payload)
}

func (o *ReplaceSnapshotDepsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotDependencies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
