// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddIntentReader is a Reader for the AddIntent structure.
type AddIntentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddIntentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddIntentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddIntentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddIntentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddIntentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddIntentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddIntentCreated creates a AddIntentCreated with default headers values
func NewAddIntentCreated() *AddIntentCreated {
	return &AddIntentCreated{}
}

/*AddIntentCreated handles this case with default header values.

The ID of the created model.
*/
type AddIntentCreated struct {
}

func (o *AddIntentCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/intents][%d] addIntentCreated ", 201)
}

func (o *AddIntentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddIntentBadRequest creates a AddIntentBadRequest with default headers values
func NewAddIntentBadRequest() *AddIntentBadRequest {
	return &AddIntentBadRequest{}
}

/*AddIntentBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.

This error can be returned if you are attempting to create an intent when the max intent count is exceeded.
```
{
      "error": {
        "code": "BadArgument",
        "message": "An application cannot have more than 80 intents."
      }
    }
```

*/
type AddIntentBadRequest struct {
}

func (o *AddIntentBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/intents][%d] addIntentBadRequest ", 400)
}

func (o *AddIntentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddIntentUnauthorized creates a AddIntentUnauthorized with default headers values
func NewAddIntentUnauthorized() *AddIntentUnauthorized {
	return &AddIntentUnauthorized{}
}

/*AddIntentUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type AddIntentUnauthorized struct {
}

func (o *AddIntentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/intents][%d] addIntentUnauthorized ", 401)
}

func (o *AddIntentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddIntentForbidden creates a AddIntentForbidden with default headers values
func NewAddIntentForbidden() *AddIntentForbidden {
	return &AddIntentForbidden{}
}

/*AddIntentForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type AddIntentForbidden struct {
}

func (o *AddIntentForbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/intents][%d] addIntentForbidden ", 403)
}

func (o *AddIntentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddIntentTooManyRequests creates a AddIntentTooManyRequests with default headers values
func NewAddIntentTooManyRequests() *AddIntentTooManyRequests {
	return &AddIntentTooManyRequests{}
}

/*AddIntentTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type AddIntentTooManyRequests struct {
}

func (o *AddIntentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appId}/versions/{versionId}/intents][%d] addIntentTooManyRequests ", 429)
}

func (o *AddIntentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
