// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5890b47c39e2bb052c5b9c2aReader is a Reader for the Nr5890b47c39e2bb052c5b9c2a structure.
type Nr5890b47c39e2bb052c5b9c2aReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5890b47c39e2bb052c5b9c2aReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5890b47c39e2bb052c5b9c2aOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5890b47c39e2bb052c5b9c2aBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5890b47c39e2bb052c5b9c2aUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5890b47c39e2bb052c5b9c2aForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5890b47c39e2bb052c5b9c2aTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5890b47c39e2bb052c5b9c2aOK creates a Nr5890b47c39e2bb052c5b9c2aOK with default headers values
func NewNr5890b47c39e2bb052c5b9c2aOK() *Nr5890b47c39e2bb052c5b9c2aOK {
	return &Nr5890b47c39e2bb052c5b9c2aOK{}
}

/*Nr5890b47c39e2bb052c5b9c2aOK handles this case with default header values.

A prebuilt entity models info.
*/
type Nr5890b47c39e2bb052c5b9c2aOK struct {
}

func (o *Nr5890b47c39e2bb052c5b9c2aOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] 5890b47c39e2bb052c5b9c2aOK ", 200)
}

func (o *Nr5890b47c39e2bb052c5b9c2aOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c2aBadRequest creates a Nr5890b47c39e2bb052c5b9c2aBadRequest with default headers values
func NewNr5890b47c39e2bb052c5b9c2aBadRequest() *Nr5890b47c39e2bb052c5b9c2aBadRequest {
	return &Nr5890b47c39e2bb052c5b9c2aBadRequest{}
}

/*Nr5890b47c39e2bb052c5b9c2aBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5890b47c39e2bb052c5b9c2aBadRequest struct {
}

func (o *Nr5890b47c39e2bb052c5b9c2aBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] 5890b47c39e2bb052c5b9c2aBadRequest ", 400)
}

func (o *Nr5890b47c39e2bb052c5b9c2aBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c2aUnauthorized creates a Nr5890b47c39e2bb052c5b9c2aUnauthorized with default headers values
func NewNr5890b47c39e2bb052c5b9c2aUnauthorized() *Nr5890b47c39e2bb052c5b9c2aUnauthorized {
	return &Nr5890b47c39e2bb052c5b9c2aUnauthorized{}
}

/*Nr5890b47c39e2bb052c5b9c2aUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type Nr5890b47c39e2bb052c5b9c2aUnauthorized struct {
}

func (o *Nr5890b47c39e2bb052c5b9c2aUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] 5890b47c39e2bb052c5b9c2aUnauthorized ", 401)
}

func (o *Nr5890b47c39e2bb052c5b9c2aUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c2aForbidden creates a Nr5890b47c39e2bb052c5b9c2aForbidden with default headers values
func NewNr5890b47c39e2bb052c5b9c2aForbidden() *Nr5890b47c39e2bb052c5b9c2aForbidden {
	return &Nr5890b47c39e2bb052c5b9c2aForbidden{}
}

/*Nr5890b47c39e2bb052c5b9c2aForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5890b47c39e2bb052c5b9c2aForbidden struct {
}

func (o *Nr5890b47c39e2bb052c5b9c2aForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] 5890b47c39e2bb052c5b9c2aForbidden ", 403)
}

func (o *Nr5890b47c39e2bb052c5b9c2aForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c2aTooManyRequests creates a Nr5890b47c39e2bb052c5b9c2aTooManyRequests with default headers values
func NewNr5890b47c39e2bb052c5b9c2aTooManyRequests() *Nr5890b47c39e2bb052c5b9c2aTooManyRequests {
	return &Nr5890b47c39e2bb052c5b9c2aTooManyRequests{}
}

/*Nr5890b47c39e2bb052c5b9c2aTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5890b47c39e2bb052c5b9c2aTooManyRequests struct {
}

func (o *Nr5890b47c39e2bb052c5b9c2aTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appId}/versions/{versionId}/prebuilts/{prebuiltId}][%d] 5890b47c39e2bb052c5b9c2aTooManyRequests ", 429)
}

func (o *Nr5890b47c39e2bb052c5b9c2aTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
