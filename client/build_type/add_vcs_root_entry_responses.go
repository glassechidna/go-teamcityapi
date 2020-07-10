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

// AddVcsRootEntryReader is a Reader for the AddVcsRootEntry structure.
type AddVcsRootEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddVcsRootEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddVcsRootEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddVcsRootEntryOK creates a AddVcsRootEntryOK with default headers values
func NewAddVcsRootEntryOK() *AddVcsRootEntryOK {
	return &AddVcsRootEntryOK{}
}

/*AddVcsRootEntryOK handles this case with default header values.

successful operation
*/
type AddVcsRootEntryOK struct {
	Payload *models.VcsRootEntry
}

func (o *AddVcsRootEntryOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/buildTypes/{btLocator}/vcs-root-entries][%d] addVcsRootEntryOK  %+v", 200, o.Payload)
}

func (o *AddVcsRootEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRootEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
