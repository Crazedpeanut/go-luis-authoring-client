// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5ade64aed5b81c209ce2e5beReader is a Reader for the Nr5ade64aed5b81c209ce2e5be structure.
type Nr5ade64aed5b81c209ce2e5beReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ade64aed5b81c209ce2e5beReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5ade64aed5b81c209ce2e5beOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ade64aed5b81c209ce2e5beBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ade64aed5b81c209ce2e5beUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ade64aed5b81c209ce2e5beForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5ade64aed5b81c209ce2e5beTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ade64aed5b81c209ce2e5beOK creates a Nr5ade64aed5b81c209ce2e5beOK with default headers values
func NewNr5ade64aed5b81c209ce2e5beOK() *Nr5ade64aed5b81c209ce2e5beOK {
	return &Nr5ade64aed5b81c209ce2e5beOK{}
}

/*Nr5ade64aed5b81c209ce2e5beOK handles this case with default header values.

Nr5ade64aed5b81c209ce2e5beOK 5ade64aed5b81c209ce2e5be o k
*/
type Nr5ade64aed5b81c209ce2e5beOK struct {
}

func (o *Nr5ade64aed5b81c209ce2e5beOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles/{roleId}][%d] 5ade64aed5b81c209ce2e5beOK ", 200)
}

func (o *Nr5ade64aed5b81c209ce2e5beOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade64aed5b81c209ce2e5beBadRequest creates a Nr5ade64aed5b81c209ce2e5beBadRequest with default headers values
func NewNr5ade64aed5b81c209ce2e5beBadRequest() *Nr5ade64aed5b81c209ce2e5beBadRequest {
	return &Nr5ade64aed5b81c209ce2e5beBadRequest{}
}

/*Nr5ade64aed5b81c209ce2e5beBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5ade64aed5b81c209ce2e5beBadRequest struct {
}

func (o *Nr5ade64aed5b81c209ce2e5beBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles/{roleId}][%d] 5ade64aed5b81c209ce2e5beBadRequest ", 400)
}

func (o *Nr5ade64aed5b81c209ce2e5beBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade64aed5b81c209ce2e5beUnauthorized creates a Nr5ade64aed5b81c209ce2e5beUnauthorized with default headers values
func NewNr5ade64aed5b81c209ce2e5beUnauthorized() *Nr5ade64aed5b81c209ce2e5beUnauthorized {
	return &Nr5ade64aed5b81c209ce2e5beUnauthorized{}
}

/*Nr5ade64aed5b81c209ce2e5beUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ade64aed5b81c209ce2e5beUnauthorized struct {
}

func (o *Nr5ade64aed5b81c209ce2e5beUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles/{roleId}][%d] 5ade64aed5b81c209ce2e5beUnauthorized ", 401)
}

func (o *Nr5ade64aed5b81c209ce2e5beUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade64aed5b81c209ce2e5beForbidden creates a Nr5ade64aed5b81c209ce2e5beForbidden with default headers values
func NewNr5ade64aed5b81c209ce2e5beForbidden() *Nr5ade64aed5b81c209ce2e5beForbidden {
	return &Nr5ade64aed5b81c209ce2e5beForbidden{}
}

/*Nr5ade64aed5b81c209ce2e5beForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ade64aed5b81c209ce2e5beForbidden struct {
}

func (o *Nr5ade64aed5b81c209ce2e5beForbidden) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles/{roleId}][%d] 5ade64aed5b81c209ce2e5beForbidden ", 403)
}

func (o *Nr5ade64aed5b81c209ce2e5beForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade64aed5b81c209ce2e5beTooManyRequests creates a Nr5ade64aed5b81c209ce2e5beTooManyRequests with default headers values
func NewNr5ade64aed5b81c209ce2e5beTooManyRequests() *Nr5ade64aed5b81c209ce2e5beTooManyRequests {
	return &Nr5ade64aed5b81c209ce2e5beTooManyRequests{}
}

/*Nr5ade64aed5b81c209ce2e5beTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5ade64aed5b81c209ce2e5beTooManyRequests struct {
}

func (o *Nr5ade64aed5b81c209ce2e5beTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appId}/versions/{versionId}/patternanyentities/{entityId}/roles/{roleId}][%d] 5ade64aed5b81c209ce2e5beTooManyRequests ", 429)
}

func (o *Nr5ade64aed5b81c209ce2e5beTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
