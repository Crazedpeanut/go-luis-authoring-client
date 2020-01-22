// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/crazedpeanut/luis/models"
)

// GetVersionsReader is a Reader for the GetVersions structure.
type GetVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVersionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVersionsOK creates a GetVersionsOK with default headers values
func NewGetVersionsOK() *GetVersionsOK {
	return &GetVersionsOK{}
}

/*GetVersionsOK handles this case with default header values.

A JSON object containing the tasks info.
*/
type GetVersionsOK struct {
	Payload []*models.GetVersionObject
}

func (o *GetVersionsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions][%d] getVersionsOK  %+v", 200, o.Payload)
}

func (o *GetVersionsOK) GetPayload() []*models.GetVersionObject {
	return o.Payload
}

func (o *GetVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionsBadRequest creates a GetVersionsBadRequest with default headers values
func NewGetVersionsBadRequest() *GetVersionsBadRequest {
	return &GetVersionsBadRequest{}
}

/*GetVersionsBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetVersionsBadRequest struct {
}

func (o *GetVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions][%d] getVersionsBadRequest ", 400)
}

func (o *GetVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsUnauthorized creates a GetVersionsUnauthorized with default headers values
func NewGetVersionsUnauthorized() *GetVersionsUnauthorized {
	return &GetVersionsUnauthorized{}
}

/*GetVersionsUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type GetVersionsUnauthorized struct {
}

func (o *GetVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions][%d] getVersionsUnauthorized ", 401)
}

func (o *GetVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsForbidden creates a GetVersionsForbidden with default headers values
func NewGetVersionsForbidden() *GetVersionsForbidden {
	return &GetVersionsForbidden{}
}

/*GetVersionsForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetVersionsForbidden struct {
}

func (o *GetVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions][%d] getVersionsForbidden ", 403)
}

func (o *GetVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsTooManyRequests creates a GetVersionsTooManyRequests with default headers values
func NewGetVersionsTooManyRequests() *GetVersionsTooManyRequests {
	return &GetVersionsTooManyRequests{}
}

/*GetVersionsTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetVersionsTooManyRequests struct {
}

func (o *GetVersionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions][%d] getVersionsTooManyRequests ", 429)
}

func (o *GetVersionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}