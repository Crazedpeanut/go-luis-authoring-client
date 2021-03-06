// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateCompositeEntityReader is a Reader for the UpdateCompositeEntity structure.
type UpdateCompositeEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCompositeEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCompositeEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCompositeEntityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCompositeEntityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCompositeEntityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCompositeEntityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCompositeEntityOK creates a UpdateCompositeEntityOK with default headers values
func NewUpdateCompositeEntityOK() *UpdateCompositeEntityOK {
	return &UpdateCompositeEntityOK{}
}

/*UpdateCompositeEntityOK handles this case with default header values.

UpdateCompositeEntityOK update composite entity o k
*/
type UpdateCompositeEntityOK struct {
}

func (o *UpdateCompositeEntityOK) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] updateCompositeEntityOK ", 200)
}

func (o *UpdateCompositeEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCompositeEntityBadRequest creates a UpdateCompositeEntityBadRequest with default headers values
func NewUpdateCompositeEntityBadRequest() *UpdateCompositeEntityBadRequest {
	return &UpdateCompositeEntityBadRequest{}
}

/*UpdateCompositeEntityBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.

This error can be returned if you are attempting to create a composite entity child when the max child count is exceeded.

```
  "error": {
    "code": "BadArgument",
    "message": "Number of entity children cannot exceed 10"
  }
```
*/
type UpdateCompositeEntityBadRequest struct {
}

func (o *UpdateCompositeEntityBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] updateCompositeEntityBadRequest ", 400)
}

func (o *UpdateCompositeEntityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCompositeEntityUnauthorized creates a UpdateCompositeEntityUnauthorized with default headers values
func NewUpdateCompositeEntityUnauthorized() *UpdateCompositeEntityUnauthorized {
	return &UpdateCompositeEntityUnauthorized{}
}

/*UpdateCompositeEntityUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type UpdateCompositeEntityUnauthorized struct {
}

func (o *UpdateCompositeEntityUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] updateCompositeEntityUnauthorized ", 401)
}

func (o *UpdateCompositeEntityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCompositeEntityForbidden creates a UpdateCompositeEntityForbidden with default headers values
func NewUpdateCompositeEntityForbidden() *UpdateCompositeEntityForbidden {
	return &UpdateCompositeEntityForbidden{}
}

/*UpdateCompositeEntityForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type UpdateCompositeEntityForbidden struct {
}

func (o *UpdateCompositeEntityForbidden) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] updateCompositeEntityForbidden ", 403)
}

func (o *UpdateCompositeEntityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCompositeEntityTooManyRequests creates a UpdateCompositeEntityTooManyRequests with default headers values
func NewUpdateCompositeEntityTooManyRequests() *UpdateCompositeEntityTooManyRequests {
	return &UpdateCompositeEntityTooManyRequests{}
}

/*UpdateCompositeEntityTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type UpdateCompositeEntityTooManyRequests struct {
}

func (o *UpdateCompositeEntityTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}][%d] updateCompositeEntityTooManyRequests ", 429)
}

func (o *UpdateCompositeEntityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
