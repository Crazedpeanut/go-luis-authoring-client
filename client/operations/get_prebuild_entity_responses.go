// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetPrebuildEntityReader is a Reader for the GetPrebuildEntity structure.
type GetPrebuildEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrebuildEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrebuildEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPrebuildEntityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPrebuildEntityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPrebuildEntityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPrebuildEntityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrebuildEntityOK creates a GetPrebuildEntityOK with default headers values
func NewGetPrebuildEntityOK() *GetPrebuildEntityOK {
	return &GetPrebuildEntityOK{}
}

/*GetPrebuildEntityOK handles this case with default header values.

A prebuilt entity models info.
*/
type GetPrebuildEntityOK struct {
}

func (o *GetPrebuildEntityOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] getPrebuildEntityOK ", 200)
}

func (o *GetPrebuildEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrebuildEntityBadRequest creates a GetPrebuildEntityBadRequest with default headers values
func NewGetPrebuildEntityBadRequest() *GetPrebuildEntityBadRequest {
	return &GetPrebuildEntityBadRequest{}
}

/*GetPrebuildEntityBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetPrebuildEntityBadRequest struct {
}

func (o *GetPrebuildEntityBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] getPrebuildEntityBadRequest ", 400)
}

func (o *GetPrebuildEntityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrebuildEntityUnauthorized creates a GetPrebuildEntityUnauthorized with default headers values
func NewGetPrebuildEntityUnauthorized() *GetPrebuildEntityUnauthorized {
	return &GetPrebuildEntityUnauthorized{}
}

/*GetPrebuildEntityUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type GetPrebuildEntityUnauthorized struct {
}

func (o *GetPrebuildEntityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] getPrebuildEntityUnauthorized ", 401)
}

func (o *GetPrebuildEntityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrebuildEntityForbidden creates a GetPrebuildEntityForbidden with default headers values
func NewGetPrebuildEntityForbidden() *GetPrebuildEntityForbidden {
	return &GetPrebuildEntityForbidden{}
}

/*GetPrebuildEntityForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetPrebuildEntityForbidden struct {
}

func (o *GetPrebuildEntityForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] getPrebuildEntityForbidden ", 403)
}

func (o *GetPrebuildEntityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrebuildEntityTooManyRequests creates a GetPrebuildEntityTooManyRequests with default headers values
func NewGetPrebuildEntityTooManyRequests() *GetPrebuildEntityTooManyRequests {
	return &GetPrebuildEntityTooManyRequests{}
}

/*GetPrebuildEntityTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetPrebuildEntityTooManyRequests struct {
}

func (o *GetPrebuildEntityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] getPrebuildEntityTooManyRequests ", 429)
}

func (o *GetPrebuildEntityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
