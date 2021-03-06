// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/crazedpeanut/go-luis-authoring-client/models"
)

// Nr5ade1315d5b81c209ce2e59fReader is a Reader for the Nr5ade1315d5b81c209ce2e59f structure.
type Nr5ade1315d5b81c209ce2e59fReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ade1315d5b81c209ce2e59fReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNr5ade1315d5b81c209ce2e59fCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ade1315d5b81c209ce2e59fBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ade1315d5b81c209ce2e59fUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ade1315d5b81c209ce2e59fForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5ade1315d5b81c209ce2e59fTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ade1315d5b81c209ce2e59fCreated creates a Nr5ade1315d5b81c209ce2e59fCreated with default headers values
func NewNr5ade1315d5b81c209ce2e59fCreated() *Nr5ade1315d5b81c209ce2e59fCreated {
	return &Nr5ade1315d5b81c209ce2e59fCreated{}
}

/*Nr5ade1315d5b81c209ce2e59fCreated handles this case with default header values.

The ID of the created role.
*/
type Nr5ade1315d5b81c209ce2e59fCreated struct {
}

func (o *Nr5ade1315d5b81c209ce2e59fCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade1315d5b81c209ce2e59fCreated ", 201)
}

func (o *Nr5ade1315d5b81c209ce2e59fCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade1315d5b81c209ce2e59fBadRequest creates a Nr5ade1315d5b81c209ce2e59fBadRequest with default headers values
func NewNr5ade1315d5b81c209ce2e59fBadRequest() *Nr5ade1315d5b81c209ce2e59fBadRequest {
	return &Nr5ade1315d5b81c209ce2e59fBadRequest{}
}

/*Nr5ade1315d5b81c209ce2e59fBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5ade1315d5b81c209ce2e59fBadRequest struct {
	Payload *models.ErrorResponseObject
}

func (o *Nr5ade1315d5b81c209ce2e59fBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade1315d5b81c209ce2e59fBadRequest  %+v", 400, o.Payload)
}

func (o *Nr5ade1315d5b81c209ce2e59fBadRequest) GetPayload() *models.ErrorResponseObject {
	return o.Payload
}

func (o *Nr5ade1315d5b81c209ce2e59fBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNr5ade1315d5b81c209ce2e59fUnauthorized creates a Nr5ade1315d5b81c209ce2e59fUnauthorized with default headers values
func NewNr5ade1315d5b81c209ce2e59fUnauthorized() *Nr5ade1315d5b81c209ce2e59fUnauthorized {
	return &Nr5ade1315d5b81c209ce2e59fUnauthorized{}
}

/*Nr5ade1315d5b81c209ce2e59fUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ade1315d5b81c209ce2e59fUnauthorized struct {
}

func (o *Nr5ade1315d5b81c209ce2e59fUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade1315d5b81c209ce2e59fUnauthorized ", 401)
}

func (o *Nr5ade1315d5b81c209ce2e59fUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade1315d5b81c209ce2e59fForbidden creates a Nr5ade1315d5b81c209ce2e59fForbidden with default headers values
func NewNr5ade1315d5b81c209ce2e59fForbidden() *Nr5ade1315d5b81c209ce2e59fForbidden {
	return &Nr5ade1315d5b81c209ce2e59fForbidden{}
}

/*Nr5ade1315d5b81c209ce2e59fForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ade1315d5b81c209ce2e59fForbidden struct {
}

func (o *Nr5ade1315d5b81c209ce2e59fForbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade1315d5b81c209ce2e59fForbidden ", 403)
}

func (o *Nr5ade1315d5b81c209ce2e59fForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade1315d5b81c209ce2e59fTooManyRequests creates a Nr5ade1315d5b81c209ce2e59fTooManyRequests with default headers values
func NewNr5ade1315d5b81c209ce2e59fTooManyRequests() *Nr5ade1315d5b81c209ce2e59fTooManyRequests {
	return &Nr5ade1315d5b81c209ce2e59fTooManyRequests{}
}

/*Nr5ade1315d5b81c209ce2e59fTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5ade1315d5b81c209ce2e59fTooManyRequests struct {
}

func (o *Nr5ade1315d5b81c209ce2e59fTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles][%d] 5ade1315d5b81c209ce2e59fTooManyRequests ", 429)
}

func (o *Nr5ade1315d5b81c209ce2e59fTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
