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

// ReplaceVcsRootEntriesReader is a Reader for the ReplaceVcsRootEntries structure.
type ReplaceVcsRootEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceVcsRootEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceVcsRootEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceVcsRootEntriesOK creates a ReplaceVcsRootEntriesOK with default headers values
func NewReplaceVcsRootEntriesOK() *ReplaceVcsRootEntriesOK {
	return &ReplaceVcsRootEntriesOK{}
}

/*ReplaceVcsRootEntriesOK handles this case with default header values.

successful operation
*/
type ReplaceVcsRootEntriesOK struct {
	Payload *models.VcsRootEntries
}

func (o *ReplaceVcsRootEntriesOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/vcs-root-entries][%d] replaceVcsRootEntriesOK  %+v", 200, o.Payload)
}

func (o *ReplaceVcsRootEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRootEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}