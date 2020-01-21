// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5ade5406d5b81c209ce2e5a7Reader is a Reader for the Nr5ade5406d5b81c209ce2e5a7 structure.
type Nr5ade5406d5b81c209ce2e5a7Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ade5406d5b81c209ce2e5a7Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNr5ade5406d5b81c209ce2e5a7Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ade5406d5b81c209ce2e5a7BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ade5406d5b81c209ce2e5a7Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ade5406d5b81c209ce2e5a7Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5ade5406d5b81c209ce2e5a7TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ade5406d5b81c209ce2e5a7Created creates a Nr5ade5406d5b81c209ce2e5a7Created with default headers values
func NewNr5ade5406d5b81c209ce2e5a7Created() *Nr5ade5406d5b81c209ce2e5a7Created {
	return &Nr5ade5406d5b81c209ce2e5a7Created{}
}

/*Nr5ade5406d5b81c209ce2e5a7Created handles this case with default header values.

The ID of the created role.
*/
type Nr5ade5406d5b81c209ce2e5a7Created struct {
}

func (o *Nr5ade5406d5b81c209ce2e5a7Created) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/customprebuiltentities/{entityId}/roles][%d] 5ade5406d5b81c209ce2e5a7Created ", 201)
}

func (o *Nr5ade5406d5b81c209ce2e5a7Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5406d5b81c209ce2e5a7BadRequest creates a Nr5ade5406d5b81c209ce2e5a7BadRequest with default headers values
func NewNr5ade5406d5b81c209ce2e5a7BadRequest() *Nr5ade5406d5b81c209ce2e5a7BadRequest {
	return &Nr5ade5406d5b81c209ce2e5a7BadRequest{}
}

/*Nr5ade5406d5b81c209ce2e5a7BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5ade5406d5b81c209ce2e5a7BadRequest struct {
}

func (o *Nr5ade5406d5b81c209ce2e5a7BadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/customprebuiltentities/{entityId}/roles][%d] 5ade5406d5b81c209ce2e5a7BadRequest ", 400)
}

func (o *Nr5ade5406d5b81c209ce2e5a7BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5406d5b81c209ce2e5a7Unauthorized creates a Nr5ade5406d5b81c209ce2e5a7Unauthorized with default headers values
func NewNr5ade5406d5b81c209ce2e5a7Unauthorized() *Nr5ade5406d5b81c209ce2e5a7Unauthorized {
	return &Nr5ade5406d5b81c209ce2e5a7Unauthorized{}
}

/*Nr5ade5406d5b81c209ce2e5a7Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ade5406d5b81c209ce2e5a7Unauthorized struct {
}

func (o *Nr5ade5406d5b81c209ce2e5a7Unauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/customprebuiltentities/{entityId}/roles][%d] 5ade5406d5b81c209ce2e5a7Unauthorized ", 401)
}

func (o *Nr5ade5406d5b81c209ce2e5a7Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5406d5b81c209ce2e5a7Forbidden creates a Nr5ade5406d5b81c209ce2e5a7Forbidden with default headers values
func NewNr5ade5406d5b81c209ce2e5a7Forbidden() *Nr5ade5406d5b81c209ce2e5a7Forbidden {
	return &Nr5ade5406d5b81c209ce2e5a7Forbidden{}
}

/*Nr5ade5406d5b81c209ce2e5a7Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ade5406d5b81c209ce2e5a7Forbidden struct {
}

func (o *Nr5ade5406d5b81c209ce2e5a7Forbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/customprebuiltentities/{entityId}/roles][%d] 5ade5406d5b81c209ce2e5a7Forbidden ", 403)
}

func (o *Nr5ade5406d5b81c209ce2e5a7Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade5406d5b81c209ce2e5a7TooManyRequests creates a Nr5ade5406d5b81c209ce2e5a7TooManyRequests with default headers values
func NewNr5ade5406d5b81c209ce2e5a7TooManyRequests() *Nr5ade5406d5b81c209ce2e5a7TooManyRequests {
	return &Nr5ade5406d5b81c209ce2e5a7TooManyRequests{}
}

/*Nr5ade5406d5b81c209ce2e5a7TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5ade5406d5b81c209ce2e5a7TooManyRequests struct {
}

func (o *Nr5ade5406d5b81c209ce2e5a7TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/customprebuiltentities/{entityId}/roles][%d] 5ade5406d5b81c209ce2e5a7TooManyRequests ", 429)
}

func (o *Nr5ade5406d5b81c209ce2e5a7TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
