// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// Nr5890b47c39e2bb052c5b9c1eReader is a Reader for the Nr5890b47c39e2bb052c5b9c1e structure.
type Nr5890b47c39e2bb052c5b9c1eReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Nr5890b47c39e2bb052c5b9c1eReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNr5890b47c39e2bb052c5b9c1eOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNr5890b47c39e2bb052c5b9c1eBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNr5890b47c39e2bb052c5b9c1eUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNr5890b47c39e2bb052c5b9c1eForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNr5890b47c39e2bb052c5b9c1eTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNr5890b47c39e2bb052c5b9c1eOK creates a Nr5890b47c39e2bb052c5b9c1eOK with default headers values
func NewNr5890b47c39e2bb052c5b9c1eOK() *Nr5890b47c39e2bb052c5b9c1eOK {
	return &Nr5890b47c39e2bb052c5b9c1eOK{}
}

/*Nr5890b47c39e2bb052c5b9c1eOK handles this case with default header values.

Nr5890b47c39e2bb052c5b9c1eOK 5890b47c39e2bb052c5b9c1e o k
*/
type Nr5890b47c39e2bb052c5b9c1eOK struct {
}

func (o *Nr5890b47c39e2bb052c5b9c1eOK) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] 5890b47c39e2bb052c5b9c1eOK ", 200)
}

func (o *Nr5890b47c39e2bb052c5b9c1eOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c1eBadRequest creates a Nr5890b47c39e2bb052c5b9c1eBadRequest with default headers values
func NewNr5890b47c39e2bb052c5b9c1eBadRequest() *Nr5890b47c39e2bb052c5b9c1eBadRequest {
	return &Nr5890b47c39e2bb052c5b9c1eBadRequest{}
}

/*Nr5890b47c39e2bb052c5b9c1eBadRequest handles this case with default header values.

This error can be returned if the request's parameters are incorrect meaning the required parameters are missing, malformed, or too large.

This error can be returned if the request's body is incorrect meaning the JSON is missing, malformed, or too large.
*/
type Nr5890b47c39e2bb052c5b9c1eBadRequest struct {
}

func (o *Nr5890b47c39e2bb052c5b9c1eBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] 5890b47c39e2bb052c5b9c1eBadRequest ", 400)
}

func (o *Nr5890b47c39e2bb052c5b9c1eBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c1eUnauthorized creates a Nr5890b47c39e2bb052c5b9c1eUnauthorized with default headers values
func NewNr5890b47c39e2bb052c5b9c1eUnauthorized() *Nr5890b47c39e2bb052c5b9c1eUnauthorized {
	return &Nr5890b47c39e2bb052c5b9c1eUnauthorized{}
}

/*Nr5890b47c39e2bb052c5b9c1eUnauthorized handles this case with default header values.

You do not have access.

Reasons can include:

* used endpoint subscription key, instead of authoring key
* invalid, malformed, or empty authoring key
* authoring key doesn't match region
* you are not the owner or collaborator
* invalid order of API calls

*/
type Nr5890b47c39e2bb052c5b9c1eUnauthorized struct {
}

func (o *Nr5890b47c39e2bb052c5b9c1eUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] 5890b47c39e2bb052c5b9c1eUnauthorized ", 401)
}

func (o *Nr5890b47c39e2bb052c5b9c1eUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c1eForbidden creates a Nr5890b47c39e2bb052c5b9c1eForbidden with default headers values
func NewNr5890b47c39e2bb052c5b9c1eForbidden() *Nr5890b47c39e2bb052c5b9c1eForbidden {
	return &Nr5890b47c39e2bb052c5b9c1eForbidden{}
}

/*Nr5890b47c39e2bb052c5b9c1eForbidden handles this case with default header values.

Total monthly key quota limit exceeded
*/
type Nr5890b47c39e2bb052c5b9c1eForbidden struct {
}

func (o *Nr5890b47c39e2bb052c5b9c1eForbidden) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] 5890b47c39e2bb052c5b9c1eForbidden ", 403)
}

func (o *Nr5890b47c39e2bb052c5b9c1eForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNr5890b47c39e2bb052c5b9c1eTooManyRequests creates a Nr5890b47c39e2bb052c5b9c1eTooManyRequests with default headers values
func NewNr5890b47c39e2bb052c5b9c1eTooManyRequests() *Nr5890b47c39e2bb052c5b9c1eTooManyRequests {
	return &Nr5890b47c39e2bb052c5b9c1eTooManyRequests{}
}

/*Nr5890b47c39e2bb052c5b9c1eTooManyRequests handles this case with default header values.

Rate limit is exceeded.
*/
type Nr5890b47c39e2bb052c5b9c1eTooManyRequests struct {
}

func (o *Nr5890b47c39e2bb052c5b9c1eTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /apps/{appId}/versions/{versionId}/entities/{entityId}][%d] 5890b47c39e2bb052c5b9c1eTooManyRequests ", 429)
}

func (o *Nr5890b47c39e2bb052c5b9c1eTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
