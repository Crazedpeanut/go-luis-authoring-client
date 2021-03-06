// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetIntentReader is a Reader for the GetIntent structure.
type GetIntentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIntentOK creates a GetIntentOK with default headers values
func NewGetIntentOK() *GetIntentOK {
	return &GetIntentOK{}
}

/*GetIntentOK handles this case with default header values.

An intent model info.
*/
type GetIntentOK struct {
}

func (o *GetIntentOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/intents/{intentId}][%d] getIntentOK ", 200)
}

func (o *GetIntentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIntentBadRequest creates a GetIntentBadRequest with default headers values
func NewGetIntentBadRequest() *GetIntentBadRequest {
	return &GetIntentBadRequest{}
}

/*GetIntentBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type GetIntentBadRequest struct {
}

func (o *GetIntentBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/intents/{intentId}][%d] getIntentBadRequest ", 400)
}

func (o *GetIntentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIntentUnauthorized creates a GetIntentUnauthorized with default headers values
func NewGetIntentUnauthorized() *GetIntentUnauthorized {
	return &GetIntentUnauthorized{}
}

/*GetIntentUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type GetIntentUnauthorized struct {
}

func (o *GetIntentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/intents/{intentId}][%d] getIntentUnauthorized ", 401)
}

func (o *GetIntentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIntentForbidden creates a GetIntentForbidden with default headers values
func NewGetIntentForbidden() *GetIntentForbidden {
	return &GetIntentForbidden{}
}

/*GetIntentForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type GetIntentForbidden struct {
}

func (o *GetIntentForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/intents/{intentId}][%d] getIntentForbidden ", 403)
}

func (o *GetIntentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIntentTooManyRequests creates a GetIntentTooManyRequests with default headers values
func NewGetIntentTooManyRequests() *GetIntentTooManyRequests {
	return &GetIntentTooManyRequests{}
}

/*GetIntentTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type GetIntentTooManyRequests struct {
}

func (o *GetIntentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/intents/{intentId}][%d] getIntentTooManyRequests ", 429)
}

func (o *GetIntentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
