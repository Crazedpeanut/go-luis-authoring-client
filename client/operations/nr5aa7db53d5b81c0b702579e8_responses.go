// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5aa7db53d5b81c0b702579e8Reader is a Reader for the Nr5aa7db53d5b81c0b702579e8 structure.
type Nr5aa7db53d5b81c0b702579e8Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5aa7db53d5b81c0b702579e8Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5aa7db53d5b81c0b702579e8OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5aa7db53d5b81c0b702579e8BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5aa7db53d5b81c0b702579e8Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5aa7db53d5b81c0b702579e8Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5aa7db53d5b81c0b702579e8TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5aa7db53d5b81c0b702579e8OK creates a Nr5aa7db53d5b81c0b702579e8OK with default headers values
func NewNr5aa7db53d5b81c0b702579e8OK() *Nr5aa7db53d5b81c0b702579e8OK {
	return &Nr5aa7db53d5b81c0b702579e8OK{}
}

/*Nr5aa7db53d5b81c0b702579e8OK handles this case with default header values.

Nr5aa7db53d5b81c0b702579e8OK 5aa7db53d5b81c0b702579e8 o k
*/
type Nr5aa7db53d5b81c0b702579e8OK struct {
}

func (o *Nr5aa7db53d5b81c0b702579e8OK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/regexentities/{regexEntityId}][%d] 5aa7db53d5b81c0b702579e8OK ", 200)
}

func (o *Nr5aa7db53d5b81c0b702579e8OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5aa7db53d5b81c0b702579e8BadRequest creates a Nr5aa7db53d5b81c0b702579e8BadRequest with default headers values
func NewNr5aa7db53d5b81c0b702579e8BadRequest() *Nr5aa7db53d5b81c0b702579e8BadRequest {
	return &Nr5aa7db53d5b81c0b702579e8BadRequest{}
}

/*Nr5aa7db53d5b81c0b702579e8BadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing or malformed.
*/
type Nr5aa7db53d5b81c0b702579e8BadRequest struct {
}

func (o *Nr5aa7db53d5b81c0b702579e8BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/regexentities/{regexEntityId}][%d] 5aa7db53d5b81c0b702579e8BadRequest ", 400)
}

func (o *Nr5aa7db53d5b81c0b702579e8BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5aa7db53d5b81c0b702579e8Unauthorized creates a Nr5aa7db53d5b81c0b702579e8Unauthorized with default headers values
func NewNr5aa7db53d5b81c0b702579e8Unauthorized() *Nr5aa7db53d5b81c0b702579e8Unauthorized {
	return &Nr5aa7db53d5b81c0b702579e8Unauthorized{}
}

/*Nr5aa7db53d5b81c0b702579e8Unauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type Nr5aa7db53d5b81c0b702579e8Unauthorized struct {
}

func (o *Nr5aa7db53d5b81c0b702579e8Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/regexentities/{regexEntityId}][%d] 5aa7db53d5b81c0b702579e8Unauthorized ", 401)
}

func (o *Nr5aa7db53d5b81c0b702579e8Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5aa7db53d5b81c0b702579e8Forbidden creates a Nr5aa7db53d5b81c0b702579e8Forbidden with default headers values
func NewNr5aa7db53d5b81c0b702579e8Forbidden() *Nr5aa7db53d5b81c0b702579e8Forbidden {
	return &Nr5aa7db53d5b81c0b702579e8Forbidden{}
}

/*Nr5aa7db53d5b81c0b702579e8Forbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5aa7db53d5b81c0b702579e8Forbidden struct {
}

func (o *Nr5aa7db53d5b81c0b702579e8Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/regexentities/{regexEntityId}][%d] 5aa7db53d5b81c0b702579e8Forbidden ", 403)
}

func (o *Nr5aa7db53d5b81c0b702579e8Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5aa7db53d5b81c0b702579e8TooManyRequests creates a Nr5aa7db53d5b81c0b702579e8TooManyRequests with default headers values
func NewNr5aa7db53d5b81c0b702579e8TooManyRequests() *Nr5aa7db53d5b81c0b702579e8TooManyRequests {
	return &Nr5aa7db53d5b81c0b702579e8TooManyRequests{}
}

/*Nr5aa7db53d5b81c0b702579e8TooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5aa7db53d5b81c0b702579e8TooManyRequests struct {
}

func (o *Nr5aa7db53d5b81c0b702579e8TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/regexentities/{regexEntityId}][%d] 5aa7db53d5b81c0b702579e8TooManyRequests ", 429)
}

func (o *Nr5aa7db53d5b81c0b702579e8TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
