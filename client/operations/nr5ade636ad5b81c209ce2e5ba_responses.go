// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5ade636ad5b81c209ce2e5baReader is a Reader for the Nr5ade636ad5b81c209ce2e5ba structure.
type Nr5ade636ad5b81c209ce2e5baReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5ade636ad5b81c209ce2e5baReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5ade636ad5b81c209ce2e5baOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5ade636ad5b81c209ce2e5baBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5ade636ad5b81c209ce2e5baUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5ade636ad5b81c209ce2e5baForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5ade636ad5b81c209ce2e5baTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5ade636ad5b81c209ce2e5baOK creates a Nr5ade636ad5b81c209ce2e5baOK with default headers values
func NewNr5ade636ad5b81c209ce2e5baOK() *Nr5ade636ad5b81c209ce2e5baOK {
	return &Nr5ade636ad5b81c209ce2e5baOK{}
}

/*Nr5ade636ad5b81c209ce2e5baOK handles this case with default header values.

Nr5ade636ad5b81c209ce2e5baOK 5ade636ad5b81c209ce2e5ba o k
*/
type Nr5ade636ad5b81c209ce2e5baOK struct {
}

func (o *Nr5ade636ad5b81c209ce2e5baOK) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles/{roleId}][%d] 5ade636ad5b81c209ce2e5baOK ", 200)
}

func (o *Nr5ade636ad5b81c209ce2e5baOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade636ad5b81c209ce2e5baBadRequest creates a Nr5ade636ad5b81c209ce2e5baBadRequest with default headers values
func NewNr5ade636ad5b81c209ce2e5baBadRequest() *Nr5ade636ad5b81c209ce2e5baBadRequest {
	return &Nr5ade636ad5b81c209ce2e5baBadRequest{}
}

/*Nr5ade636ad5b81c209ce2e5baBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.
*/
type Nr5ade636ad5b81c209ce2e5baBadRequest struct {
}

func (o *Nr5ade636ad5b81c209ce2e5baBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles/{roleId}][%d] 5ade636ad5b81c209ce2e5baBadRequest ", 400)
}

func (o *Nr5ade636ad5b81c209ce2e5baBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade636ad5b81c209ce2e5baUnauthorized creates a Nr5ade636ad5b81c209ce2e5baUnauthorized with default headers values
func NewNr5ade636ad5b81c209ce2e5baUnauthorized() *Nr5ade636ad5b81c209ce2e5baUnauthorized {
	return &Nr5ade636ad5b81c209ce2e5baUnauthorized{}
}

/*Nr5ade636ad5b81c209ce2e5baUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls
*/
type Nr5ade636ad5b81c209ce2e5baUnauthorized struct {
}

func (o *Nr5ade636ad5b81c209ce2e5baUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles/{roleId}][%d] 5ade636ad5b81c209ce2e5baUnauthorized ", 401)
}

func (o *Nr5ade636ad5b81c209ce2e5baUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade636ad5b81c209ce2e5baForbidden creates a Nr5ade636ad5b81c209ce2e5baForbidden with default headers values
func NewNr5ade636ad5b81c209ce2e5baForbidden() *Nr5ade636ad5b81c209ce2e5baForbidden {
	return &Nr5ade636ad5b81c209ce2e5baForbidden{}
}

/*Nr5ade636ad5b81c209ce2e5baForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5ade636ad5b81c209ce2e5baForbidden struct {
}

func (o *Nr5ade636ad5b81c209ce2e5baForbidden) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles/{roleId}][%d] 5ade636ad5b81c209ce2e5baForbidden ", 403)
}

func (o *Nr5ade636ad5b81c209ce2e5baForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5ade636ad5b81c209ce2e5baTooManyRequests creates a Nr5ade636ad5b81c209ce2e5baTooManyRequests with default headers values
func NewNr5ade636ad5b81c209ce2e5baTooManyRequests() *Nr5ade636ad5b81c209ce2e5baTooManyRequests {
	return &Nr5ade636ad5b81c209ce2e5baTooManyRequests{}
}

/*Nr5ade636ad5b81c209ce2e5baTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5ade636ad5b81c209ce2e5baTooManyRequests struct {
}

func (o *Nr5ade636ad5b81c209ce2e5baTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/compositeentities/{cEntityId}/roles/{roleId}][%d] 5ade636ad5b81c209ce2e5baTooManyRequests ", 429)
}

func (o *Nr5ade636ad5b81c209ce2e5baTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
